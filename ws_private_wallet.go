package mybybitapi

import "fmt"

func getWalletSubscribeArg() string {
	return fmt.Sprintf("wallet")
}

// 批量订阅钱包推送
func (ws *PrivateWsStreamClient) SubscribeWallet() (*Subscription[WsWallet], error) {
	args := []string{getWalletSubscribeArg()}

	doSub, err := ws.subscribe(SUBSCRIBE, args)
	if err != nil {
		return nil, err
	}
	err = ws.catchSubscribeResult(doSub)
	if err != nil {
		return nil, err
	}
	log.Infof("SubscribeWallet Success: args:%v", doSub.Args)
	sub := &Subscription[WsWallet]{
		SubId:      doSub.SubId,
		Op:         SUBSCRIBE,
		Args:       doSub.Args,
		resultChan: make(chan WsWallet, 50),
		errChan:    make(chan error),
		closeChan:  make(chan struct{}),
		Ws:         &ws.WsStreamClient,
	}
	for _, arg := range args {
		ws.walletSubMap.Store(arg, sub)
		ws.commonSubMap.Store(arg, doSub)
	}
	return sub, nil
}

// 批量取消订阅钱包推送
func (ws *PrivateWsStreamClient) UnSubscribeWallet() error {
	args := []string{getWalletSubscribeArg()}

	doSub, err := ws.subscribe(UNSUBSCRIBE, args)
	if err != nil {
		return err
	}
	err = ws.catchSubscribeResult(doSub)
	if err != nil {
		return err
	}
	log.Infof("UnSubscribeWallet Success: args:%v", doSub.Args)

	for _, arg := range args {
		ws.deferUnSubscribeWallet(arg)
	}
	return nil
}

func (ws *PrivateWsStreamClient) deferUnSubscribeWallet(arg string) {
	if sub, ok := ws.commonSubMap.Load(arg); ok {
		newArgs := []string{}
		for _, a := range sub.Args {
			if a != arg {
				newArgs = append(newArgs, a)
			}
		}
		sub.Args = newArgs
	}
	if sub, ok := ws.walletSubMap.Load(arg); ok {
		newArgs := []string{}
		for _, a := range sub.Args {
			if a != arg {
				newArgs = append(newArgs, a)
			}
		}
		sub.Args = newArgs
	}
	ws.walletSubMap.Delete(arg)
	ws.commonSubMap.Delete(arg)
}
