package mybybitapi

import "fmt"

func getTradeSubscribeArg(symbol string) string {
	return fmt.Sprintf("publicTrade.%s", symbol)
}

// 订阅单个成交 如: "BTCUSDT"
func (ws *PublicWsStreamClient) SubscribeTrade(symbol string) (*Subscription[WsTrade], error) {
	return ws.SubscribeTradeMultiple([]string{symbol})
}

// 批量订阅成交 如: ["BTCUSDT","ETHUSDT"]
func (ws *PublicWsStreamClient) SubscribeTradeMultiple(symbols []string) (*Subscription[WsTrade], error) {
	args := []string{}
	for _, s := range symbols {
		arg := getTradeSubscribeArg(s)
		args = append(args, arg)
	}
	id, err := generateReqId()
	if err != nil {
		return nil, err
	}
	sub := &Subscription[WsTrade]{
		SubId:      id,
		Op:         SUBSCRIBE,
		Args:       args,
		resultChan: make(chan WsTrade, 50),
		errChan:    make(chan error),
		closeChan:  make(chan struct{}),
		Ws:         &ws.WsStreamClient,
	}
	for _, arg := range args {
		ws.tradeSubMap.Store(arg, sub)
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
		log.Infof("SubscribeTrade Success: args:%v", doSub.Args)
	}()

	for _, arg := range args {
		ws.commonSubMap.Store(arg, doSub)
	}
	return sub, nil
}

// 取消订阅单个成交 如: "BTCUSDT"
func (ws *PublicWsStreamClient) UnSubscribeTrade(symbol string) error {
	return ws.UnSubscribeTradeMultiple([]string{symbol})
}

// 批量取消订阅成交 如: ["BTCUSDT","ETHUSDT"]
func (ws *PublicWsStreamClient) UnSubscribeTradeMultiple(symbols []string) error {
	args := []string{}
	for _, s := range symbols {
		arg := getTradeSubscribeArg(s)
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
	log.Infof("UnSubscribeTrade Success: args:%v", doSub.Args)

	for _, arg := range args {
		ws.deferUnSubscribeTrade(arg)
	}
	return nil
}

func (ws *PublicWsStreamClient) deferUnSubscribeTrade(arg string) {
	if sub, ok := ws.commonSubMap.Load(arg); ok {
		newArgs := []string{}
		for _, a := range sub.Args {
			if a != arg {
				newArgs = append(newArgs, a)
			}
		}
		sub.Args = newArgs
	}
	if sub, ok := ws.tradeSubMap.Load(arg); ok {
		newArgs := []string{}
		for _, a := range sub.Args {
			if a != arg {
				newArgs = append(newArgs, a)
			}
		}
		sub.Args = newArgs
	}
	ws.tradeSubMap.Delete(arg)
	ws.commonSubMap.Delete(arg)
}
