package mybybitapi

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/gorilla/websocket"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	BYBIT_API_WS_PUBLIC_SPOT    = "/v5/public/spot"    //现货公共频道
	BYBIT_API_WS_PUBLIC_LINEAR  = "/v5/public/linear"  //合约公共频道
	BYBIT_API_WS_PUBLIC_INVERSE = "/v5/public/inverse" //反向合约公共频道
	BYBIT_API_WS_PUBLIC_OPTION  = "/v5/public/option"  //期权公共频道
	BYBIT_API_WS_PRIVATE        = "/v5/private"        //私有频道
)

const (
	AUTH         = "auth"         //鉴权
	SUBSCRIBE    = "subscribe"    //订阅
	UNSUBSCRIBE  = "unsubscribe"  //取消订阅
	ORDER_CREATE = "order.create" //下单
	ORDER_AMEND  = "order.amend"  //撤单
	ORDER_CANCEL = "order.cancel" //批量撤单
)

var (
	WebsocketTimeout        = time.Second * 10
	WebsocketKeepalive      = true
	SUBSCRIBE_INTERVAL_TIME = 500 * time.Millisecond //订阅间隔
)

// 数据流订阅基础客户端
type WsStreamClient struct {
	client       *RestClient
	lastAuth     *WsAuthReq
	apiType      APIType
	conn         *websocket.Conn
	connId       string
	commonSubMap MySyncMap[string, *Subscription[WsActionResult]] //订阅的返回结果

	waitSubResultMap MySyncMap[string, *Subscription[WsActionResult]]

	klineSubMap MySyncMap[string, *Subscription[WsKline]]
	depthSubMap MySyncMap[string, *Subscription[WsDepth]]
	tradeSubMap MySyncMap[string, *Subscription[WsTrade]]

	orderSubMap     MySyncMap[string, *Subscription[WsOrder]]
	walletSubMap    MySyncMap[string, *Subscription[WsWallet]]
	positionSubMap  MySyncMap[string, *Subscription[WsPosition]]
	executionSubMap MySyncMap[string, *Subscription[WsExecution]]

	resultChan chan []byte
	errChan    chan error
	isClose    bool

	reSubscribeMu      *sync.Mutex
	AutoReConnectTimes int //自动重连次数
}

// 登陆请求相关
type WsAuthReq struct {
	ReqId string         `json:"req_id"`
	Op    string         `json:"op"`   //String 是操作
	Args  [3]interface{} `json:"args"` //Array 是请求订阅的频道列表
}

type WsAuthArg struct {
	ApiKey    string `json:"apiKey"`
	Expire    int64  `json:"expire"`
	Signature string `json:"signature"`
}

// 订阅请求相关
type WsSubscribeReq struct {
	ReqId string   `json:"req_id"`
	Op    string   `json:"op"`   //String 是操作
	Args  []string `json:"args"` //Array 是请求订阅的频道列表
}

// 登陆及订阅返回结果
type WsActionResult struct {
	Success bool   `json:"success"`
	RetMsg  string `json:"ret_msg"`
	ConnId  string `json:"conn_id"`
	ReqId   string `json:"req_id"`
	Op      string `json:"op"`
}

// 数据流订阅标准结构体
type Subscription[T any] struct {
	SubId        int64           //订阅ID
	Ws           *WsStreamClient //订阅的连接
	Op           string          //订阅方法
	Args         []string        //订阅参数
	resultChan   chan T          //接收订阅结果的通道
	errChan      chan error      //接收订阅错误的通道
	closeChan    chan struct{}   //接收订阅关闭的通道
	subResultMap map[string]bool //订阅结果
}

// 获取订阅结果
func (sub *Subscription[T]) ResultChan() chan T {
	return sub.resultChan
}

// 获取错误订阅
func (sub *Subscription[T]) ErrChan() chan error {
	return sub.errChan
}

// 获取关闭订阅信号
func (sub *Subscription[T]) CloseChan() chan struct{} {
	return sub.closeChan
}

func (ws *WsStreamClient) GetConn() *websocket.Conn {
	return ws.conn
}

func (ws *WsStreamClient) auth(op string, arg WsAuthArg) (*Subscription[WsActionResult], error) {
	if ws == nil || ws.conn == nil || ws.isClose {
		return nil, fmt.Errorf("websocket is close")
	}
	node, err := snowflake.NewNode(3)
	if err != nil {
		return nil, err
	}
	id := node.Generate().Int64()

	authReq := WsAuthReq{
		ReqId: strconv.FormatInt(id, 10),
		Op:    op,
		Args:  [3]interface{}{arg.ApiKey, arg.Expire, arg.Signature},
	}
	data, err := json.Marshal(authReq)
	if err != nil {
		return nil, err
	}

	ws.lastAuth = &authReq
	resultSub := &Subscription[WsActionResult]{
		SubId:        id,
		Op:           op,
		Args:         []string{},
		resultChan:   make(chan WsActionResult, 50),
		errChan:      make(chan error),
		closeChan:    make(chan struct{}),
		Ws:           ws,
		subResultMap: map[string]bool{},
	}
	ws.waitSubResultMap.Store(strconv.FormatInt(id, 10), resultSub)
	log.Debugf("send msg: %s", string(data))
	err = ws.conn.WriteMessage(websocket.TextMessage, data)
	if err != nil {
		return nil, err
	}
	return resultSub, nil
}

func (ws *WsStreamClient) subscribe(op string, args []string) (*Subscription[WsActionResult], error) {
	if ws == nil || ws.conn == nil || ws.isClose {
		return nil, fmt.Errorf("websocket is close")
	}

	node, err := snowflake.NewNode(3)
	if err != nil {
		return nil, err
	}
	id := node.Generate().Int64()

	subscribeReq := WsSubscribeReq{
		ReqId: strconv.FormatInt(id, 10),
		Op:    op,
		Args:  args,
	}
	data, err := json.Marshal(subscribeReq)
	if err != nil {
		return nil, err
	}

	resultSub := &Subscription[WsActionResult]{
		SubId:        id,
		Op:           op,
		Args:         args,
		resultChan:   make(chan WsActionResult, 50),
		errChan:      make(chan error),
		closeChan:    make(chan struct{}),
		Ws:           ws,
		subResultMap: map[string]bool{},
	}
	ws.waitSubResultMap.Store(strconv.FormatInt(id, 10), resultSub)

	log.Debugf("send msg: %s", string(data))
	err = ws.conn.WriteMessage(websocket.TextMessage, data)
	if err != nil {
		return nil, err
	}

	return resultSub, nil
}

func (ws *WsStreamClient) Close() error {
	ws.isClose = true
	ws.connId = ""

	err := ws.conn.Close()
	if err != nil {
		return err
	}
	//手动关闭成功，给所有订阅发送关闭信号
	ws.sendWsCloseToAllSub()

	//初始化连接状态
	ws.conn = nil
	close(ws.resultChan)
	close(ws.errChan)
	ws.resultChan = nil
	ws.errChan = nil
	ws.lastAuth = nil
	ws.commonSubMap = NewMySyncMap[string, *Subscription[WsActionResult]]()
	ws.klineSubMap = NewMySyncMap[string, *Subscription[WsKline]]()
	ws.depthSubMap = NewMySyncMap[string, *Subscription[WsDepth]]()
	ws.tradeSubMap = NewMySyncMap[string, *Subscription[WsTrade]]()
	ws.orderSubMap = NewMySyncMap[string, *Subscription[WsOrder]]()
	ws.walletSubMap = NewMySyncMap[string, *Subscription[WsWallet]]()
	ws.positionSubMap = NewMySyncMap[string, *Subscription[WsPosition]]()
	ws.executionSubMap = NewMySyncMap[string, *Subscription[WsExecution]]()

	if ws.waitSubResultMap.Length() != 0 {
		//给当前等待订阅结果的请求返回错误
		ws.waitSubResultMap.Range(func(key string, value *Subscription[WsActionResult]) bool {
			value.errChan <- fmt.Errorf("websocket is closed")
			ws.waitSubResultMap.Delete(key)
			return true
		})
	}

	return nil
}

func (ws *WsStreamClient) OpenConn() error {
	if ws.resultChan == nil {
		ws.resultChan = make(chan []byte)
	}
	if ws.errChan == nil {
		ws.errChan = make(chan error)
	}
	apiUrl := handlerWsStreamRequestApi(ws.apiType)
	if ws.conn == nil {
		conn, err := wsStreamServe(apiUrl, ws.resultChan, ws.errChan)
		ws.conn = conn
		ws.isClose = false
		ws.connId = ""
		log.Info("OpenConn success to ", apiUrl)
		ws.handleResult(ws.resultChan, ws.errChan)
		return err
	} else {
		conn, err := wsStreamServe(apiUrl, ws.resultChan, ws.errChan)
		ws.conn = conn
		ws.connId = ""
		log.Info("Auto ReOpenConn success to ", apiUrl)
		return err
	}

}

type PublicWsStreamClient struct {
	WsStreamClient
}
type PrivateWsStreamClient struct {
	WsStreamClient
}

func newWsStreamClient(apiType APIType) WsStreamClient {
	return WsStreamClient{
		apiType:         apiType,
		commonSubMap:    NewMySyncMap[string, *Subscription[WsActionResult]](),
		klineSubMap:     NewMySyncMap[string, *Subscription[WsKline]](),
		depthSubMap:     NewMySyncMap[string, *Subscription[WsDepth]](),
		tradeSubMap:     NewMySyncMap[string, *Subscription[WsTrade]](),
		orderSubMap:     NewMySyncMap[string, *Subscription[WsOrder]](),
		walletSubMap:    NewMySyncMap[string, *Subscription[WsWallet]](),
		positionSubMap:  NewMySyncMap[string, *Subscription[WsPosition]](),
		executionSubMap: NewMySyncMap[string, *Subscription[WsExecution]](),

		waitSubResultMap: NewMySyncMap[string, *Subscription[WsActionResult]](),
		reSubscribeMu:    &sync.Mutex{},
	}
}

func NewPublicSpotWsStreamClient() *PublicWsStreamClient {
	return &PublicWsStreamClient{
		WsStreamClient: newWsStreamClient(WS_PUBLIC_SPOT),
	}
}
func NewPublicLinearWsStreamClient() *PublicWsStreamClient {
	return &PublicWsStreamClient{
		WsStreamClient: newWsStreamClient(WS_PUBLIC_LINEAR),
	}
}
func NewPublicInverseWsStreamClient() *PublicWsStreamClient {
	return &PublicWsStreamClient{
		WsStreamClient: newWsStreamClient(WS_PUBLIC_INVERSE),
	}
}
func NewPublicOptionWsStreamClient() *PublicWsStreamClient {
	return &PublicWsStreamClient{
		WsStreamClient: newWsStreamClient(WS_PUBLIC_OPTION),
	}
}
func NewPrivateWsStreamClient() *PrivateWsStreamClient {
	return &PrivateWsStreamClient{
		WsStreamClient: newWsStreamClient(WS_PRIVATE),
	}
}

func (ws *WsStreamClient) CurrentSubList() []string {
	list := []string{}
	ws.commonSubMap.Range(func(key string, _ *Subscription[WsActionResult]) bool {
		list = append(list, key)
		return true
	})
	return list
}

func (ws *WsStreamClient) sendSubscribeResultToChan(result WsActionResult) {
	if ws.connId == "" && result.ConnId != "" {
		ws.connId = result.ConnId
	}
	if sub, ok := ws.waitSubResultMap.Load(result.ReqId); ok {
		if !result.Success {
			d, _ := json.Marshal(result)
			sub.errChan <- fmt.Errorf("errHandler: %s", string(d))
		} else {
			sub.resultChan <- result
		}
	}
}

func (ws *WsStreamClient) sendUnSubscribeSuccessToCloseChan(args []string) {
	for _, arg := range args {
		data, _ := json.Marshal(&arg)
		key := string(data)
		if sub, ok := ws.klineSubMap.Load(key); ok {
			ws.klineSubMap.Delete(key)
			if sub.closeChan != nil {
				sub.closeChan <- struct{}{}
				sub.closeChan = nil
			}
		}
		if sub, ok := ws.depthSubMap.Load(key); ok {
			ws.depthSubMap.Delete(key)
			if sub.closeChan != nil {
				sub.closeChan <- struct{}{}
				sub.closeChan = nil
			}
		}
		if sub, ok := ws.tradeSubMap.Load(key); ok {
			ws.tradeSubMap.Delete(key)
			if sub.closeChan != nil {
				sub.closeChan <- struct{}{}
				sub.closeChan = nil
			}
		}
		if sub, ok := ws.orderSubMap.Load(key); ok {
			ws.orderSubMap.Delete(key)
			if sub.closeChan != nil {
				sub.closeChan <- struct{}{}
				sub.closeChan = nil
			}
		}
		if sub, ok := ws.walletSubMap.Load(key); ok {
			ws.walletSubMap.Delete(key)
			if sub.closeChan != nil {
				sub.closeChan <- struct{}{}
				sub.closeChan = nil
			}
		}
		if sub, ok := ws.positionSubMap.Load(key); ok {
			ws.positionSubMap.Delete(key)
			if sub.closeChan != nil {
				sub.closeChan <- struct{}{}
				sub.closeChan = nil
			}

		}
		if sub, ok := ws.executionSubMap.Load(key); ok {
			ws.executionSubMap.Delete(key)
			if sub.closeChan != nil {
				sub.closeChan <- struct{}{}
				sub.closeChan = nil
			}

		}
	}
}

func (ws *WsStreamClient) sendWsCloseToAllSub() {
	args := []string{}
	ws.commonSubMap.Range(func(key string, _ *Subscription[WsActionResult]) bool {
		arg := key
		args = append(args, arg)
		return true
	})
	ws.sendUnSubscribeSuccessToCloseChan(args)
}

func (ws *WsStreamClient) reSubscribeForReconnect() error {
	ws.reSubscribeMu.Lock()
	defer ws.reSubscribeMu.Unlock()
	isDoReSubscribe := map[int64]bool{}
	var wErr error
	ws.commonSubMap.Range(func(_ string, sub *Subscription[WsActionResult]) bool {
		if _, ok := isDoReSubscribe[sub.SubId]; ok {
			return true
		}

		reSub, err := ws.subscribe(sub.Op, sub.Args)
		if err != nil {
			log.Error(err)
			wErr = err
			return false
		}
		err = ws.catchSubscribeResult(reSub)
		if err != nil {
			log.Error(err)
			wErr = err
			return false
		}
		log.Infof("reSubscribe Success: args:%v", reSub.Args)

		sub.SubId = reSub.SubId
		isDoReSubscribe[sub.SubId] = true
		time.Sleep(1000 * time.Millisecond)
		return true
	})
	return wErr
}

func (ws *WsStreamClient) handleResult(resultChan chan []byte, errChan chan error) {
	go func() {
		for {
			select {
			case err, ok := <-errChan:
				if !ok {
					log.Error("errChan is closed")
					return
				}
				log.Error(err)
				//错误处理 重连等
				//ws标记为非关闭 且返回错误包含EOF、close、reset时自动重连
				if !ws.isClose && (strings.Contains(err.Error(), "EOF") ||
					strings.Contains(err.Error(), "close") ||
					strings.Contains(err.Error(), "reset")) {
					//重连
					err := ws.OpenConn()
					for err != nil {
						time.Sleep(2000 * time.Millisecond)
						err = ws.OpenConn()
					}
					ws.AutoReConnectTimes += 1
					go func() {
						//重新登陆
						if ws.lastAuth != nil && ws.client != nil {
							err = ws.Auth(ws.client)
							for err != nil {
								time.Sleep(2000 * time.Millisecond)
								err = ws.Auth(ws.client)
							}
						}

						//重新订阅
						err = ws.reSubscribeForReconnect()
						if err != nil {
							log.Error(err)
						}
					}()
				} else {
					continue
				}
			case data, ok := <-resultChan:
				if !ok {
					log.Error("resultChan is closed")
					return
				}
				// log.Debug("receive result: ", string(data))
				//处理订阅或查询订阅列表请求返回结果
				if strings.Contains(string(data), "success") || strings.Contains(string(data), "connId") {
					result := WsActionResult{}
					err := json.Unmarshal(data, &result)
					if err != nil {
						log.Error(err)
						continue
					}
					ws.sendSubscribeResultToChan(result)
					continue
				}

				//处理正常数据的返回结果
				//K线处理
				if strings.Contains(string(data), "topic\":\"kline") {
					c, err := handleWsKline(data)

					key := c.Topic
					if sub, ok := ws.klineSubMap.Load(key); ok {
						if err != nil {
							sub.errChan <- err
							continue
						}
						sub.resultChan <- *c
					}
					continue
				}

				//深度处理
				if strings.Contains(string(data), "topic\":\"orderbook") {
					c, err := handleWsDepth(data)

					key := c.Topic
					if sub, ok := ws.depthSubMap.Load(key); ok {
						if err != nil {
							sub.errChan <- err
							continue
						}
						sub.resultChan <- *c
					}
					continue
				}

				//平台成交处理
				if strings.Contains(string(data), "topic\":\"publicTrade") {
					c, err := handleWsTrade(data)

					key := c.Topic
					if sub, ok := ws.tradeSubMap.Load(key); ok {
						if err != nil {
							sub.errChan <- err
							continue
						}
						sub.resultChan <- *c
					}
					continue
				}

				//订单推送处理
				if strings.Contains(string(data), "topic\":\"order") {
					c, err := handleWsOrder(data)

					key := c.Topic
					if sub, ok := ws.orderSubMap.Load(key); ok {
						if err != nil {
							sub.errChan <- err
							continue
						}
						sub.resultChan <- *c
					}
					continue
				}

				//钱包推送处理
				if strings.Contains(string(data), "topic\":\"wallet") {
					c, err := handleWsWallet(data)

					key := c.Topic
					if sub, ok := ws.walletSubMap.Load(key); ok {
						if err != nil {
							sub.errChan <- err
							continue
						}
						sub.resultChan <- *c
					}
					continue
				}

				//仓位推送处理
				if strings.Contains(string(data), "topic\":\"position") {
					c, err := handleWsPosition(data)

					key := c.Topic
					if sub, ok := ws.positionSubMap.Load(key); ok {
						if err != nil {
							sub.errChan <- err
							continue
						}
						sub.resultChan <- *c
					}
					continue
				}

				//个人成交推送处理
				if strings.Contains(string(data), "topic\":\"execution") {
					c, err := handleWsExecution(data)

					key := c.Topic
					if sub, ok := ws.executionSubMap.Load(key); ok {
						if err != nil {
							sub.errChan <- err
							continue
						}
						sub.resultChan <- *c
					}
					continue
				}
			}
		}
	}()
}

func (ws *WsStreamClient) DeferSub(sub *Subscription[WsActionResult]) {
	ws.waitSubResultMap.Delete(strconv.FormatInt(sub.SubId, 10))
}

// 取消订阅
func (sub *Subscription[T]) Unsubscribe() error {

	unSub, err := sub.Ws.subscribe(UNSUBSCRIBE, sub.Args)
	if err != nil {
		return err
	}
	err = sub.Ws.catchSubscribeResult(unSub)
	if err != nil {
		return err
	}
	log.Debugf("Unsubscribe Success args:%v ", unSub.Args)

	//取消订阅成功，给所有订阅消息的通道发送关闭信号
	sub.Ws.sendUnSubscribeSuccessToCloseChan(unSub.Args)
	//删除当前订阅列表中已存在的记录
	for _, arg := range unSub.Args {
		sub.Ws.commonSubMap.Delete(arg)
	}
	return nil
}

// 捕获订阅结果
func (ws *WsStreamClient) CatchAuthResult(sub *Subscription[WsActionResult]) error {
	defer sub.Ws.DeferSub(sub)

	select {
	case err := <-sub.ErrChan():
		// log.Error(err)
		return fmt.Errorf("auth error: %v", err)
	case authResult := <-sub.ResultChan():
		if !authResult.Success {
			err := fmt.Errorf("%s:%s:%s", authResult.ReqId, authResult.RetMsg, authResult.Op)
			log.Error(err)
			return err
		}
		log.Debug("catchAuthResults: ", authResult)
		return nil
	}
}

// 捕获订阅结果
func (ws *WsStreamClient) catchSubscribeResult(sub *Subscription[WsActionResult]) error {
	defer sub.Ws.DeferSub(sub)

	select {
	case err := <-sub.ErrChan():
		log.Error(err)
		return fmt.Errorf("SubAction Error: %v", err)
	case subResult := <-sub.ResultChan():
		if !subResult.Success {
			err := fmt.Errorf("%s:%s:%s", subResult.ReqId, subResult.RetMsg, subResult.Op)
			log.Error(err)
			return err
		}

		for _, arg := range sub.Args {
			sub.subResultMap[arg] = true
		}

	}

	log.Debug(sub.Op, " success: ", sub.Args)
	return nil
}

// bybit websocket登陆功能
func (ws *WsStreamClient) Auth(client *RestClient) error {
	ws.client = client

	expire := time.Now().UnixNano()/1e6 + 10000
	val := fmt.Sprintf("GET/realtime%d", expire)
	mac := hmac.New(sha256.New, []byte(client.c.SecretKey))
	mac.Write([]byte(val))
	signature := hex.EncodeToString(mac.Sum(nil))
	authArg := WsAuthArg{
		ApiKey:    client.c.APIKey,
		Expire:    expire,
		Signature: signature,
	}
	authSub, err := ws.auth(AUTH, authArg)
	if err != nil {
		return err
	}

	err = ws.CatchAuthResult(authSub)
	if err != nil {
		return err
	}
	log.Infof("Auth Success args:%v ", ws.lastAuth.Args)
	return nil
}

// 标准订阅方法
func wsStreamServe(api string, resultChan chan []byte, errChan chan error) (*websocket.Conn, error) {
	c, _, err := websocket.DefaultDialer.Dial(api, nil)
	if err != nil {
		return nil, err
	}
	c.SetReadLimit(655350)
	go func() {
		if WebsocketKeepalive {
			keepAlive(c, WebsocketTimeout)
		}
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				errChan <- err
				return
			}
			resultChan <- message
		}
	}()
	return c, err
}

// 获取数据流请求URL
func handlerWsStreamRequestApi(apiType APIType) string {
	host := BYBIT_API_WEBSOCKET
	switch SERVER_TYPE {
	case BASE:
		host = BYBIT_API_WEBSOCKET
	case AWS:
		host = BYBIT_API_WEBSOCKET_AWS
	case TEST:
		host = BYBIT_API_WEBSOCKET_TEST
	default:
	}

	u := url.URL{
		Scheme:   "wss",
		Host:     host,
		Path:     getWsApi(apiType),
		RawQuery: "",
	}
	return u.String()
}

// 获取数据流请求Path
func getWsApi(apiType APIType) string {
	switch apiType {
	case WS_PUBLIC_SPOT:
		return BYBIT_API_WS_PUBLIC_SPOT
	case WS_PUBLIC_LINEAR:
		return BYBIT_API_WS_PUBLIC_LINEAR
	case WS_PUBLIC_INVERSE:
		return BYBIT_API_WS_PUBLIC_INVERSE
	case WS_PUBLIC_OPTION:
		return BYBIT_API_WS_PUBLIC_OPTION
	case WS_PRIVATE:
		return BYBIT_API_WS_PRIVATE
	default:
		log.Error("apiType Error is ", apiType)
		return ""
	}
}

// 发送ping/pong消息以检查连接稳定性
func keepAlive(c *websocket.Conn, timeout time.Duration) {
	ticker := time.NewTicker(timeout)

	lastResponse := time.Now()
	c.SetPongHandler(func(msg string) error {
		lastResponse = time.Now()
		return nil
	})

	go func() {
		defer ticker.Stop()
		for {
			deadline := time.Now().Add(10 * time.Second)
			err := c.WriteControl(websocket.PingMessage, []byte{}, deadline)
			if err != nil {
				return
			}
			<-ticker.C
			if time.Since(lastResponse) > timeout {
				err := c.Close()
				if err != nil {
					log.Error(err)
					return
				}
				return
			}
		}
	}()
}
