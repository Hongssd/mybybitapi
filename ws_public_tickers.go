package mybybitapi

import "fmt"

func getTickerSubscribeArg(symbol string) string {
	return fmt.Sprintf("tickers.%s", symbol)
}

// 订阅单个行情 如: "BTCUSDT"
func (ws *PublicWsStreamClient) SubscribeTicker(symbol string) (*Subscription[WsTicker], error) {
	return ws.SubscribeTickerMultiple([]string{symbol})
}

// 批量订阅行情 如: ["BTCUSDT","ETHUSDT"]
func (ws *PublicWsStreamClient) SubscribeTickerMultiple(symbols []string) (*Subscription[WsTicker], error) {
	args := []string{}
	for _, s := range symbols {
		arg := getTickerSubscribeArg(s)
		args = append(args, arg)
	}
	id, err := generateReqId()
	if err != nil {
		return nil, err
	}
	sub := &Subscription[WsTicker]{
		SubId:      id,
		Op:         SUBSCRIBE,
		Args:       args,
		resultChan: make(chan WsTicker, 50),
		errChan:    make(chan error),
		closeChan:  make(chan struct{}),
		Ws:         &ws.WsStreamClient,
	}
	for _, arg := range args {
		ws.tickerSubMap.Store(arg, sub)
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
		log.Infof("SubscribeTicker Success: args:%v", doSub.Args)
	}()

	for _, arg := range args {
		ws.commonSubMap.Store(arg, doSub)
	}
	return sub, nil
}

// 取消订阅单个行情 如: "BTCUSDT"
func (ws *PublicWsStreamClient) UnSubscribeTicker(symbol string) error {
	return ws.UnSubscribeTickerMultiple([]string{symbol})
}

// 批量取消订阅行情 如: ["BTCUSDT","ETHUSDT"]
func (ws *PublicWsStreamClient) UnSubscribeTickerMultiple(symbols []string) error {
	args := []string{}
	for _, s := range symbols {
		arg := getTickerSubscribeArg(s)
		args = append(args, arg)
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
	log.Infof("UnSubscribeTicker Success: args:%v", doSub.Args)

	for _, arg := range args {
		ws.deferUnSubscribeTicker(arg)
	}
	return nil
}

func (ws *PublicWsStreamClient) deferUnSubscribeTicker(arg string) {
	if sub, ok := ws.commonSubMap.Load(arg); ok {
		newArgs := []string{}
		for _, a := range sub.Args {
			if a != arg {
				newArgs = append(newArgs, a)
			}
		}
		sub.Args = newArgs
	}
	if sub, ok := ws.tickerSubMap.Load(arg); ok {
		newArgs := []string{}
		for _, a := range sub.Args {
			if a != arg {
				newArgs = append(newArgs, a)
			}
		}
		sub.Args = newArgs
	}
	ws.tickerSubMap.Delete(arg)
	ws.commonSubMap.Delete(arg)
}
