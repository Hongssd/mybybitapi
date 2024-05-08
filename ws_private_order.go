package mybybitapi

import "fmt"

func getOrderSubscribeArg(category string) string {
	if category == "" || category == "all" {
		return fmt.Sprintf("order")
	} else {
		return fmt.Sprintf("order.%s", category)
	}

}

// 订阅单个订单推送 如: "spot"
func (ws *PrivateWsStreamClient) SubscribeOrder(category string) (*Subscription[WsOrder], error) {
	return ws.SubscribeOrderMultiple([]string{category})
}

// 批量订阅订单推送 如: ["spot","linear"]
func (ws *PrivateWsStreamClient) SubscribeOrderMultiple(categories []string) (*Subscription[WsOrder], error) {
	args := []string{}

	for _, c := range categories {
		args = append(args, getOrderSubscribeArg(c))
	}

	doSub, err := ws.subscribe(SUBSCRIBE, args)
	if err != nil {
		return nil, err
	}
	err = ws.catchSubscribeResult(doSub)
	if err != nil {
		return nil, err
	}
	log.Infof("SubscribeOrder Success: args:%v", doSub.Args)
	sub := &Subscription[WsOrder]{
		SubId:      doSub.SubId,
		Op:         SUBSCRIBE,
		Args:       doSub.Args,
		resultChan: make(chan WsOrder, 50),
		errChan:    make(chan error),
		closeChan:  make(chan struct{}),
		Ws:         &ws.WsStreamClient,
	}
	for _, arg := range args {
		ws.orderSubMap.Store(arg, sub)
		ws.commonSubMap.Store(arg, doSub)
	}
	return sub, nil
}

// 取消订阅单个订单推送 如: "spot"
func (ws *PrivateWsStreamClient) UnSubscribeOrder(category string) error {
	return ws.UnSubscribeOrderMultiple([]string{category})
}

// 批量取消订阅订单推送 如: ["spot","linear"]
func (ws *PrivateWsStreamClient) UnSubscribeOrderMultiple(categories []string) error {
	args := []string{}
	for _, c := range categories {
		arg := getOrderSubscribeArg(c)
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
	log.Infof("UnSubscribeOrder Success: args:%v", doSub.Args)

	for _, arg := range args {
		ws.deferUnSubscribeOrder(arg)
	}
	return nil
}

func (ws *PrivateWsStreamClient) deferUnSubscribeOrder(arg string) {
	if sub, ok := ws.commonSubMap.Load(arg); ok {
		newArgs := []string{}
		for _, a := range sub.Args {
			if a != arg {
				newArgs = append(newArgs, a)
			}
		}
		sub.Args = newArgs
	}
	if sub, ok := ws.orderSubMap.Load(arg); ok {
		newArgs := []string{}
		for _, a := range sub.Args {
			if a != arg {
				newArgs = append(newArgs, a)
			}
		}
		sub.Args = newArgs
	}
	ws.orderSubMap.Delete(arg)
	ws.commonSubMap.Delete(arg)
}
