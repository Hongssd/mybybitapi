package mybybitapi

import "fmt"

func getDepthSubscribeArg(symbol, depth string) string {
	return fmt.Sprintf("orderbook.%s.%s", depth, symbol)
}

// 订阅单个深度 如: "BTCUSDT","1"
func (ws *PublicWsStreamClient) SubscribeDepth(symbol, depth string) (*Subscription[WsDepth], error) {
	return ws.SubscribeDepthMultiple([]string{symbol}, depth)
}

// 批量订阅深度 如: ["BTCUSDT","ETHUSDT"],"1"
func (ws *PublicWsStreamClient) SubscribeDepthMultiple(symbols []string, depth string) (*Subscription[WsDepth], error) {
	args := []string{}
	for _, s := range symbols {
		arg := getDepthSubscribeArg(s, depth)
		args = append(args, arg)
	}
	doSub, err := ws.subscribe(SUBSCRIBE, args)
	if err != nil {
		return nil, err
	}
	err = ws.catchSubscribeResult(doSub)
	if err != nil {
		return nil, err
	}
	log.Infof("SubscribeDepth Success: args:%v", doSub.Args)
	sub := &Subscription[WsDepth]{
		SubId:      doSub.SubId,
		Op:         SUBSCRIBE,
		Args:       doSub.Args,
		resultChan: make(chan WsDepth, 50),
		errChan:    make(chan error),
		closeChan:  make(chan struct{}),
		Ws:         &ws.WsStreamClient,
	}
	for _, arg := range args {
		ws.depthSubMap.Store(arg, sub)
		ws.commonSubMap.Store(arg, doSub)
	}
	return sub, nil
}

// 取消订阅单个深度 如: "BTCUSDT","1"
func (ws *PublicWsStreamClient) UnSubscribeDepth(symbol, depth string) error {
	return ws.UnSubscribeDepthMultiple([]string{symbol}, depth)
}

// 批量取消订阅深度 如: ["BTCUSDT","ETHUSDT"],"1"
func (ws *PublicWsStreamClient) UnSubscribeDepthMultiple(symbols []string, depth string) error {
	args := []string{}
	for _, s := range symbols {
		arg := getDepthSubscribeArg(s, depth)
		args = append(args, arg)

	}
	doSub, err := ws.subscribe(UNSUBSCRIBE, args)
	if err != nil {
		return err
	}
	err = ws.catchSubscribeResult(doSub)
	if err != nil {
		return err
	}
	log.Infof("UnSubscribeDepth Success: args:%v", doSub.Args)

	for _, arg := range args {
		ws.deferUnSubscribeDepth(arg)
	}
	return nil
}

func (ws *PublicWsStreamClient) deferUnSubscribeDepth(arg string) {
	if sub, ok := ws.commonSubMap.Load(arg); ok {
		newArgs := []string{}
		for _, a := range sub.Args {
			if a != arg {
				newArgs = append(newArgs, a)
			}
		}
		sub.Args = newArgs
	}
	if sub, ok := ws.depthSubMap.Load(arg); ok {
		newArgs := []string{}
		for _, a := range sub.Args {
			if a != arg {
				newArgs = append(newArgs, a)
			}
		}
		sub.Args = newArgs
	}
	ws.depthSubMap.Delete(arg)
	ws.commonSubMap.Delete(arg)
}
