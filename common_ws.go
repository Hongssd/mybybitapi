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
	BYBIT_API_WS_TRADE          = "/v5/trade"          //ws下单交易频道
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

	waitAuthResult   *Subscription[WsAuthResult]
	waitAuthResultMu *sync.Mutex
	waitSubResultMap MySyncMap[string, *Subscription[WsActionResult]]

	klineSubMap MySyncMap[string, *Subscription[WsKline]]
	depthSubMap MySyncMap[string, *Subscription[WsDepth]]
	tradeSubMap MySyncMap[string, *Subscription[WsTrade]]

	orderSubMap     MySyncMap[string, *Subscription[WsOrder]]
	walletSubMap    MySyncMap[string, *Subscription[WsWallet]]
	positionSubMap  MySyncMap[string, *Subscription[WsPosition]]
	executionSubMap MySyncMap[string, *Subscription[WsExecution]]

	waitOrderCreateResult *WsOrderParamSend[OrderCreateReq, OrderCreateRes]
	waitOrderCancelResult *WsOrderParamSend[OrderCancelReq, OrderCancelRes]
	waitOrderAmendResult  *WsOrderParamSend[OrderAmendReq, OrderAmendRes]
	waitOrderCreateMu     *sync.Mutex
	waitOrderCancelMu     *sync.Mutex
	waitOrderAmendMu      *sync.Mutex

	resultChan chan []byte
	errChan    chan error
	isClose    bool

	reSubscribeMu      *sync.Mutex
	AutoReConnectTimes int //自动重连次数
}

// 登陆请求相关
type WsAuthReq struct {
	ReqId string         `json:"req_id,omitempty"`
	Op    string         `json:"op"`   //String 是操作
	Args  [3]interface{} `json:"args"` //Array 是请求订阅的频道列表
}

// 授权请求参数
type WsAuthArg struct {
	ApiKey    string `json:"apiKey"`
	Expire    int64  `json:"expire"`
	Signature string `json:"signature"`
}

type OrderReqType interface {
	OrderCreateReq | OrderCancelReq | OrderAmendReq
}

type OrderResType interface {
	OrderCreateRes | OrderCancelRes | OrderAmendRes
}

type WsOrderReqHeader struct {
	XBapiTimestamp  string `json:"X-BAPI-TIMESTAMP"`
	XBapiRecvWindow string `json:"X-BAPI-RECV-WINDOW"`
	Referer         string `json:"Referer"`
}

type WsOrderResHeader struct {
	Traceid                  string `json:"Traceid"`                      //Trace ID, 用於追蹤請求鏈路 (內部使用)
	Timenow                  string `json:"Timenow"`                      //當前時間戳
	XBapiLimit               string `json:"X-Bapi-Limit"`                 //該類型請求的帳戶總頻率
	XBapiLimitStatus         string `json:"X-Bapi-Limit-Status"`          //該類型請求的帳戶剩餘可用頻率
	XBapiLimitResetTimestamp string `json:"X-Bapi-Limit-Reset-Timestamp"` //如果您已超過該接口當前窗口頻率限製，該字段表示下個可用時間窗口的時間戳（毫秒）即什麽時候可以恢復訪問；如果您未超過該接口當前窗口頻率限製，該字段表示返回的是當前服務器時間（毫秒).
}

// ws下单/撤单/改单请求参数
type WsOrderArg[T OrderReqType] struct {
	ReqId  string           `json:"req_id"` //請求reqId, 可作為請求的唯一標識, 若有傳, 則響應會返回該字段 當傳, 需保證唯一, 否則將會拿到錯誤 "20006"
	Header WsOrderReqHeader `json:"header"` //請求頭
	Op     string           `json:"op"`     //Op類型 order.create: 創建訂單 order.amend: 修改訂單 order.cancel: 撤銷訂單
	Args   [1]T             `json:"args"`   //參數數組, 目前僅支持一個元素 order.create: 請參閱創建訂單請求參數 order.amend: 請參閱修改訂單參數 order.cancel: 請參閱撤銷訂單參數
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

// 单独授权接口返回结果
type WsAuthResult struct {
	RetCode int    `json:"retCode"`
	RetMsg  string `json:"retMsg"`
	Op      string `json:"op"`
	ConnId  string `json:"connId"`
}

// 下单/撤单/改单返回结果
type WsOrderResult[T OrderResType] struct {
	ReqId   string           `json:"req_id"`   //若請求有傳, 則響應存在該字段 若請求不傳, 則響應沒有該字段
	RetCode int              `json:"ret_code"` //0: 成功 10404: 1. op類型未找到; 2. category不支持/未找到 10429: 觸發系統級別的頻率保護 20006: reqId重複 10016: 1.內部錯誤; 2. 服務重啟 10019: ws下單服務正在重啟, 拒絕新的請求, 正在處理中的請求不受影響. 您可以重新/新建連接, 會分配到正常的服務上
	RetMsg  string           `json:"ret_msg"`  //OK "" 報錯信息
	Op      string           `json:"op"`       //Op類型
	Data    T                `json:"data"`     //業務數據, 和rest api響應的result字段業務數據一致 order.create: 請參閱創建訂單響應參數 order.amend: 請參閱修改訂單響應參數 order.cancel: 請參閱取消訂單響應參數
	Header  WsOrderResHeader `json:"header"`   //響應頭信息 TraceId string Trace ID, 用於追蹤請求鏈路 (內部使用) Timenow string 當前時間戳 X-Bapi-Limit string 該類型請求的帳戶總頻率 X-Bapi-Limit-Status string 該類型請求的帳戶剩餘可用頻率 X-Bapi-Limit-Reset-Timestamp string 如果您已超過該接口當前窗口頻率限製，該字段表示下個可用時間窗口的時間戳（毫秒）即什麽時候可以恢復訪問；如果您未超過該接口當前窗口頻率限製，該字段表示返回的是當前服務器時間（毫秒).
	ConnId  string           `json:"conn_id"`  //連接的唯一id
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

// ws下单/撤单/改单单次请求
type WsOrderParamSend[T OrderReqType, R OrderResType] struct {
	ReqId      string          //请求ID
	Ws         *WsStreamClient //订阅的连接
	Op         string          //订阅方法
	Args       WsOrderArg[T]
	resultChan chan WsOrderResult[R]
	errChan    chan error
	closeChan  chan struct{}
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

// 获取订阅结果
func (sub *WsOrderParamSend[T, R]) ResultChan() chan WsOrderResult[R] {
	return sub.resultChan
}

// 获取错误订阅
func (sub *WsOrderParamSend[T, R]) ErrChan() chan error {
	return sub.errChan
}

// 获取关闭订阅信号
func (sub *WsOrderParamSend[T, R]) CloseChan() chan struct{} {
	return sub.closeChan
}

func (ws *WsStreamClient) authPrivate(op string, arg WsAuthArg) (*Subscription[WsActionResult], error) {
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

func (ws *WsStreamClient) authTrade(op string, arg WsAuthArg) (*Subscription[WsAuthResult], error) {
	if ws == nil || ws.conn == nil || ws.isClose {
		return nil, fmt.Errorf("websocket is close")
	}

	authReq := WsAuthReq{
		Op:   op,
		Args: [3]interface{}{arg.ApiKey, arg.Expire, arg.Signature},
	}
	data, err := json.Marshal(authReq)
	if err != nil {
		return nil, err
	}

	ws.lastAuth = &authReq
	resultSub := &Subscription[WsAuthResult]{
		Op:           op,
		Args:         []string{},
		resultChan:   make(chan WsAuthResult, 50),
		errChan:      make(chan error),
		closeChan:    make(chan struct{}),
		Ws:           ws,
		subResultMap: map[string]bool{},
	}
	ws.waitAuthResult = resultSub
	log.Debugf("send msg: %s", string(data))
	err = ws.conn.WriteMessage(websocket.TextMessage, data)
	if err != nil {
		return nil, err
	}
	return resultSub, nil
}

func (ws *WsStreamClient) subscribe(id int64, op string, args []string) (*Subscription[WsActionResult], error) {
	if ws == nil || ws.conn == nil || ws.isClose {
		return nil, fmt.Errorf("websocket is close")
	}

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

func doOrder[T OrderReqType, R OrderResType](ws *WsStreamClient, op string, arg T) (*WsOrderParamSend[T, R], error) {
	if ws == nil || ws.conn == nil || ws.isClose {
		return nil, fmt.Errorf("websocket is close")
	}

	node, err := snowflake.NewNode(3)
	if err != nil {
		return nil, err
	}
	id := node.Generate().Int64()

	doOrderArg := WsOrderArg[T]{
		ReqId: strconv.FormatInt(id, 10),
		Header: WsOrderReqHeader{
			XBapiTimestamp:  strconv.FormatInt(time.Now().UnixNano()/1e6, 10),
			XBapiRecvWindow: ws.client.c.RecvWindow,
			Referer:         ws.client.c.Referer,
		},
		Op:   op,
		Args: [1]T{arg},
	}
	data, err := json.Marshal(doOrderArg)
	if err != nil {
		return nil, err
	}

	resultParamSend := &WsOrderParamSend[T, R]{
		ReqId:      strconv.FormatInt(id, 10),
		Op:         op,
		Args:       doOrderArg,
		resultChan: make(chan WsOrderResult[R], 50),
		errChan:    make(chan error),
		closeChan:  make(chan struct{}),
		Ws:         ws,
	}
	log.Debugf("send msg: %s", string(data))

	err = ws.conn.WriteMessage(websocket.TextMessage, data)
	if err != nil {
		return nil, err
	}
	return resultParamSend, nil
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

	if ws.waitAuthResult != nil {
		ws.waitAuthResult.errChan <- fmt.Errorf("websocket is closed")
		ws.waitAuthResult = nil
	}

	if ws.waitSubResultMap.Length() != 0 {
		//给当前等待订阅结果的请求返回错误
		ws.waitSubResultMap.Range(func(key string, value *Subscription[WsActionResult]) bool {
			value.errChan <- fmt.Errorf("websocket is closed")
			ws.waitSubResultMap.Delete(key)
			return true
		})
	}

	if ws.waitOrderCreateResult != nil {
		//给当前等待下单结果的请求返回错误
		ws.waitOrderCreateResult.errChan <- fmt.Errorf("websocket is closed")
		ws.waitOrderCreateResult = nil
	}

	if ws.waitOrderCancelResult != nil {
		//给当前等待撤单结果的请求返回错误
		ws.waitOrderCancelResult.errChan <- fmt.Errorf("websocket is closed")
		ws.waitOrderCancelResult = nil
	}

	if ws.waitOrderAmendResult != nil {
		//给当前等待改单结果的请求返回错误
		ws.waitOrderAmendResult.errChan <- fmt.Errorf("websocket is closed")
		ws.waitOrderAmendResult = nil

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
type TradeWsStreamClient struct {
	WsStreamClient
}

func newWsStreamClient(apiType APIType) WsStreamClient {
	return WsStreamClient{
		apiType:      apiType,
		commonSubMap: NewMySyncMap[string, *Subscription[WsActionResult]](),

		waitAuthResult:   nil,
		waitAuthResultMu: &sync.Mutex{},
		waitSubResultMap: NewMySyncMap[string, *Subscription[WsActionResult]](),

		klineSubMap:     NewMySyncMap[string, *Subscription[WsKline]](),
		depthSubMap:     NewMySyncMap[string, *Subscription[WsDepth]](),
		tradeSubMap:     NewMySyncMap[string, *Subscription[WsTrade]](),
		orderSubMap:     NewMySyncMap[string, *Subscription[WsOrder]](),
		walletSubMap:    NewMySyncMap[string, *Subscription[WsWallet]](),
		positionSubMap:  NewMySyncMap[string, *Subscription[WsPosition]](),
		executionSubMap: NewMySyncMap[string, *Subscription[WsExecution]](),

		waitOrderCreateResult: nil,
		waitOrderCancelResult: nil,
		waitOrderAmendResult:  nil,
		waitOrderCreateMu:     &sync.Mutex{},
		waitOrderCancelMu:     &sync.Mutex{},
		waitOrderAmendMu:      &sync.Mutex{},
		reSubscribeMu:         &sync.Mutex{},
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

func NewTradeWsStreamClient() *TradeWsStreamClient {
	return &TradeWsStreamClient{
		WsStreamClient: newWsStreamClient(WS_TRADE),
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

func (ws *WsStreamClient) sendAuthResultToChan(result WsAuthResult) {
	if ws.connId == "" && result.ConnId != "" {
		ws.connId = result.ConnId
	}
	if ws.waitAuthResult != nil {
		if result.RetCode != 0 {
			d, _ := json.Marshal(result)
			ws.waitAuthResult.errChan <- fmt.Errorf("errHandler: %s", string(d))
		} else {
			ws.waitAuthResult.resultChan <- result
		}
	}
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
		id, err := generateReqId()
		if err != nil {
			log.Error(err)
			wErr = err
			return false
		}
		reSub, err := ws.subscribe(id, sub.Op, sub.Args)
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
				//log.Debug("receive result: ", string(data))
				//处理订阅或查询订阅列表请求返回结果
				if strings.Contains(string(data), "success") {
					result := WsActionResult{}
					err := json.Unmarshal(data, &result)
					if err != nil {
						log.Error(err)
						continue
					}
					ws.sendSubscribeResultToChan(result)
					continue
				}
				if strings.Contains(string(data), "op\":\"auth") {
					result := WsAuthResult{}
					err := json.Unmarshal(data, &result)
					if err != nil {
						log.Error(err)
						continue
					}
					ws.sendAuthResultToChan(result)
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

				//ws下单结果处理
				if strings.Contains(string(data), "op\":\"order.create") {
					c, err := handleWsDoOrderResult[OrderCreateRes](data)

					if ws.waitOrderCreateResult != nil {
						if err != nil {
							ws.waitOrderCreateResult.errChan <- err
							continue
						}
						ws.waitOrderCreateResult.resultChan <- *c
					}
					continue
				}

				//ws撤单结果处理
				if strings.Contains(string(data), "op\":\"order.cancel") {
					c, err := handleWsDoOrderResult[OrderCancelRes](data)

					if ws.waitOrderCancelResult != nil {
						if err != nil {
							ws.waitOrderCancelResult.errChan <- err
							continue
						}
						ws.waitOrderCancelResult.resultChan <- *c
					}
					continue
				}

				//ws改单结果处理
				if strings.Contains(string(data), "op\":\"order.amend") {
					c, err := handleWsDoOrderResult[OrderAmendRes](data)

					if ws.waitOrderAmendResult != nil {
						if err != nil {
							ws.waitOrderAmendResult.errChan <- err
							continue
						}
						ws.waitOrderAmendResult.resultChan <- *c
					}
					continue
				}
			}
		}
	}()
}

// 取消订阅
func (sub *Subscription[T]) Unsubscribe() error {
	id, err := generateReqId()
	if err != nil {
		return err
	}
	unSub, err := sub.Ws.subscribe(id, UNSUBSCRIBE, sub.Args)
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

// 捕获鉴权结果
func (ws *WsStreamClient) CatchPrivateAuthResult(sub *Subscription[WsActionResult]) error {
	defer func() {
		ws.waitAuthResult = nil
	}()

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
	case <-sub.CloseChan():
		return fmt.Errorf("Auth CloseChan")
	}
}

// 捕获鉴权结果
func (ws *WsStreamClient) CatchTradeAuthResult(sub *Subscription[WsAuthResult]) error {
	defer func() {
		ws.waitAuthResult = nil
	}()

	select {
	case err := <-sub.ErrChan():
		// log.Error(err)
		return fmt.Errorf("auth error: %v", err)
	case authResult := <-sub.ResultChan():
		if authResult.RetCode != 0 {
			err := fmt.Errorf("%s:%s:%s", authResult.RetCode, authResult.RetMsg, authResult.Op)
			log.Error(err)
			return err
		}
		log.Debug("catchAuthResults: ", authResult)
		return nil
	case <-sub.CloseChan():
		return fmt.Errorf("Auth CloseChan")
	}
}

// 捕获订阅结果
func (ws *WsStreamClient) catchSubscribeResult(sub *Subscription[WsActionResult]) error {
	defer func() {
		ws.waitSubResultMap.Delete(strconv.FormatInt(sub.SubId, 10))
	}()

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
	case <-sub.CloseChan():
		return fmt.Errorf("SubAction CloseChan")
	}

	log.Debug(sub.Op, " success: ", sub.Args)
	return nil
}

// 捕获下单/撤单/查单结果
func CatchDoOrderResult[T OrderReqType, R OrderResType](ws *WsStreamClient, send *WsOrderParamSend[T, R]) (*WsOrderResult[R], error) {
	defer func() {
		switch send.Op {
		case ORDER_CREATE:
			ws.waitOrderCreateResult = nil
		case ORDER_CANCEL:
			ws.waitOrderCancelResult = nil
		case ORDER_AMEND:
			ws.waitOrderAmendResult = nil
		}
	}()

	select {
	case err := <-send.ErrChan():
		return nil, fmt.Errorf("%s error: %v", send.Op, err)
	case orderCreateResult := <-send.ResultChan():
		if orderCreateResult.RetCode != 0 {
			err := fmt.Errorf("%s:%s:%s", orderCreateResult.ReqId, orderCreateResult.RetMsg, orderCreateResult.Op)
			log.Error(err)
			return nil, err
		}
		log.Debugf("catch %s Results:%+v ", send.Op, orderCreateResult)
		return &orderCreateResult, nil
	case <-send.CloseChan():
		return nil, fmt.Errorf("DoOrder CloseChan")
	}
}

func (ws *WsStreamClient) Auth(client *RestClient) error {
	if ws.apiType == WS_PRIVATE {
		return ws.authPrivateDo(client)
	} else if ws.apiType == WS_TRADE {
		return ws.authTradeDo(client)
	} else {
		return fmt.Errorf("apiType is error")
	}
}

// bybit websocket登陆功能
func (ws *WsStreamClient) authPrivateDo(client *RestClient) error {
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

	authSub, err := ws.authPrivate(AUTH, authArg)
	if err != nil {
		return err
	}

	err = ws.CatchPrivateAuthResult(authSub)
	if err != nil {
		return err
	}
	log.Infof("Auth Success args:%v ", ws.lastAuth.Args)
	return nil
}

// bybit websocket登陆功能
func (ws *WsStreamClient) authTradeDo(client *RestClient) error {
	ws.client = client
	if ws.waitAuthResultMu.TryLock() {
		defer ws.waitAuthResultMu.Unlock()
	} else {
		return fmt.Errorf("websocket is authing")
	}
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

	authSub, err := ws.authTrade(AUTH, authArg)
	if err != nil {
		return err
	}

	err = ws.CatchTradeAuthResult(authSub)
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
	case WS_TRADE:
		return BYBIT_API_WS_TRADE
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
			if time.Since(lastResponse) > timeout*10 {
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

func generateReqId() (int64, error) {
	node, err := snowflake.NewNode(3)
	if err != nil {
		return 0, err
	}
	id := node.Generate().Int64()
	return id, nil
}
