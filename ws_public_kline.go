package mybybitapi

import "fmt"

func getKlineSubscribeArg(symbol, interval string) string {
	return fmt.Sprintf("kline.%s.%s", interval, symbol)
}

// 订阅单个K线 如: "BTCUSDT","1"
func (ws *PublicWsStreamClient) SubscribeKline(symbol, interval string) (*Subscription[WsKline], error) {
	return ws.SubscribeKlineMultiple([]string{symbol}, []string{interval})
}

// 批量订阅K线 如: ["BTCUSDT","ETHUSDT"],["1","5"]
func (ws *PublicWsStreamClient) SubscribeKlineMultiple(symbols []string, intervals []string) (*Subscription[WsKline], error) {
	args := []string{}
	for _, s := range symbols {
		for _, i := range intervals {
			arg := getKlineSubscribeArg(s, i)
			args = append(args, arg)
		}
	}
	id, err := generateReqId()
	if err != nil {
		return nil, err
	}
	sub := &Subscription[WsKline]{
		SubId:      id,
		Op:         SUBSCRIBE,
		Args:       args,
		resultChan: make(chan WsKline, 50),
		errChan:    make(chan error),
		closeChan:  make(chan struct{}),
		Ws:         &ws.WsStreamClient,
	}
	for _, arg := range args {
		ws.klineSubMap.Store(arg, sub)
	}
	doSub, err := ws.subscribe(id, SUBSCRIBE, args)
	if err != nil {
		return nil, err
	}
	go func() {
		err = ws.catchSubscribeResult(doSub)
		if err != nil {
			doSub.ErrChan() <- err
		}
		log.Infof("SubscribeKline Success: args:%v", doSub.Args)

	}()

	for _, arg := range args {
		ws.commonSubMap.Store(arg, doSub)
	}
	return sub, nil
}

// 取消订阅单个K线 如: "BTCUSDT","1m"
func (ws *PublicWsStreamClient) UnSubscribeKline(symbol, interval string) error {
	return ws.UnSubscribeKlineMultiple([]string{symbol}, []string{interval})
}

// 批量取消订阅K线 如: ["BTCUSDT","ETHUSDT"],["1","5"]
func (ws *PublicWsStreamClient) UnSubscribeKlineMultiple(symbols, intervals []string) error {
	args := []string{}
	for _, s := range symbols {
		for _, i := range intervals {
			arg := getKlineSubscribeArg(s, i)
			args = append(args, arg)
		}
	}
	id, err := generateReqId()
	if err != nil {
		return err
	}
	doSub, err := ws.subscribe(id, UNSUBSCRIBE, args)
	if err != nil {
		return err
	}
	err = ws.catchSubscribeResult(doSub)
	if err != nil {
		return err
	}
	log.Infof("UnSubscribeKline Success: args:%v", doSub.Args)

	for _, arg := range args {
		ws.deferUnSubscribeKline(arg)
	}
	return nil
}

func (ws *PublicWsStreamClient) deferUnSubscribeKline(arg string) {
	if sub, ok := ws.commonSubMap.Load(arg); ok {
		newArgs := []string{}
		for _, a := range sub.Args {
			if a != arg {
				newArgs = append(newArgs, a)
			}
		}
		sub.Args = newArgs
	}
	if sub, ok := ws.klineSubMap.Load(arg); ok {
		newArgs := []string{}
		for _, a := range sub.Args {
			if a != arg {
				newArgs = append(newArgs, a)
			}
		}
		sub.Args = newArgs
	}
	ws.klineSubMap.Delete(arg)
	ws.commonSubMap.Delete(arg)
}
