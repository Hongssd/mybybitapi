package mybybitapi

import "fmt"

func getPositionSubscribeArg(category string) string {
	if category == "" || category == "all" {
		return fmt.Sprintf("position")
	} else {
		return fmt.Sprintf("position.%s", category)
	}

}

// 订阅单个仓位推送 如: "spot"
func (ws *PrivateWsStreamClient) SubscribePosition(category string) (*Subscription[WsPosition], error) {
	return ws.SubscribePositionMultiple([]string{category})
}

// 批量订阅仓位推送 如: ["spot","linear"]
func (ws *PrivateWsStreamClient) SubscribePositionMultiple(categories []string) (*Subscription[WsPosition], error) {
	args := []string{}

	for _, c := range categories {
		args = append(args, getPositionSubscribeArg(c))
	}

	doSub, err := ws.subscribe(SUBSCRIBE, args)
	if err != nil {
		return nil, err
	}
	err = ws.catchSubscribeResult(doSub)
	if err != nil {
		return nil, err
	}
	log.Infof("SubscribePosition Success: args:%v", doSub.Args)
	sub := &Subscription[WsPosition]{
		SubId:      doSub.SubId,
		Op:         SUBSCRIBE,
		Args:       doSub.Args,
		resultChan: make(chan WsPosition, 50),
		errChan:    make(chan error),
		closeChan:  make(chan struct{}),
		Ws:         &ws.WsStreamClient,
	}
	for _, arg := range args {
		ws.positionSubMap.Store(arg, sub)
		ws.commonSubMap.Store(arg, doSub)
	}
	return sub, nil
}

// 取消订阅单个仓位推送 如: "spot"
func (ws *PrivateWsStreamClient) UnSubscribePosition(category string) error {
	return ws.UnSubscribePositionMultiple([]string{category})
}

// 批量取消订阅仓位推送 如: ["spot","linear"]
func (ws *PrivateWsStreamClient) UnSubscribePositionMultiple(categories []string) error {
	args := []string{}
	for _, c := range categories {
		arg := getPositionSubscribeArg(c)
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
	log.Infof("UnSubscribePosition Success: args:%v", doSub.Args)

	for _, arg := range args {
		ws.deferUnSubscribePosition(arg)
	}
	return nil
}

func (ws *PrivateWsStreamClient) deferUnSubscribePosition(arg string) {
	if sub, ok := ws.commonSubMap.Load(arg); ok {
		newArgs := []string{}
		for _, a := range sub.Args {
			if a != arg {
				newArgs = append(newArgs, a)
			}
		}
		sub.Args = newArgs
	}
	if sub, ok := ws.positionSubMap.Load(arg); ok {
		newArgs := []string{}
		for _, a := range sub.Args {
			if a != arg {
				newArgs = append(newArgs, a)
			}
		}
		sub.Args = newArgs
	}
	ws.positionSubMap.Delete(arg)
	ws.commonSubMap.Delete(arg)
}
