package mybybitapi

import "fmt"

func getExecutionSubscribeArg(category string) string {
	if category == "" || category == "all" {
		return fmt.Sprintf("execution")
	} else {
		return fmt.Sprintf("execution.%s", category)
	}

}

// 订阅单个个人成交推送 如: "spot"
func (ws *PrivateWsStreamClient) SubscribeExecution(category string) (*Subscription[WsExecution], error) {
	return ws.SubscribeExecutionMultiple([]string{category})
}

// 批量订阅个人成交推送 如: ["spot","linear"]
func (ws *PrivateWsStreamClient) SubscribeExecutionMultiple(categories []string) (*Subscription[WsExecution], error) {
	args := []string{}

	for _, c := range categories {
		args = append(args, getExecutionSubscribeArg(c))
	}

	id, err := generateReqId()
	if err != nil {
		return nil, err
	}
	sub := &Subscription[WsExecution]{
		SubId:      id,
		Op:         SUBSCRIBE,
		Args:       args,
		resultChan: make(chan WsExecution, 50),
		errChan:    make(chan error),
		closeChan:  make(chan struct{}),
		Ws:         &ws.WsStreamClient,
	}
	for _, arg := range args {
		ws.executionSubMap.Store(arg, sub)
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
		log.Infof("SubscribeExecution Success: args:%v", doSub.Args)
	}()

	for _, arg := range args {
		ws.commonSubMap.Store(arg, doSub)
	}
	return sub, nil
}

// 取消订阅单个个人成交推送 如: "spot"
func (ws *PrivateWsStreamClient) UnSubscribeExecution(category string) error {
	return ws.UnSubscribeExecutionMultiple([]string{category})
}

// 批量取消订阅个人成交推送 如: ["spot","linear"]
func (ws *PrivateWsStreamClient) UnSubscribeExecutionMultiple(categories []string) error {
	args := []string{}
	for _, c := range categories {
		arg := getExecutionSubscribeArg(c)
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
	log.Infof("UnSubscribeExecution Success: args:%v", doSub.Args)

	for _, arg := range args {
		ws.deferUnSubscribeExecution(arg)
	}
	return nil
}

func (ws *PrivateWsStreamClient) deferUnSubscribeExecution(arg string) {
	if sub, ok := ws.commonSubMap.Load(arg); ok {
		newArgs := []string{}
		for _, a := range sub.Args {
			if a != arg {
				newArgs = append(newArgs, a)
			}
		}
		sub.Args = newArgs
	}
	if sub, ok := ws.executionSubMap.Load(arg); ok {
		newArgs := []string{}
		for _, a := range sub.Args {
			if a != arg {
				newArgs = append(newArgs, a)
			}
		}
		sub.Args = newArgs
	}
	ws.executionSubMap.Delete(arg)
	ws.commonSubMap.Delete(arg)
}
