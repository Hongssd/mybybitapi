package mybybitapi

import "fmt"

// 下单
func (ws *TradeWsStreamClient) CreateOrder(api *OrderCreateAPI) (*WsOrderResult[OrderCreateRes], error) {
	if ws.waitOrderCreateMu.TryLock() {
		defer ws.waitOrderCreateMu.Unlock()
	} else {
		return nil, fmt.Errorf("another order create request is executing")
	}

	if ws.waitOrderCreateResult != nil {
		return nil, fmt.Errorf("another order create request waiting for result")
	}

	orderSend, err := doOrder[OrderCreateReq, OrderCreateRes](&ws.WsStreamClient, ORDER_CREATE, *api.req)
	if err != nil {
		return nil, err
	}

	ws.waitOrderCreateResult = orderSend

	result, err := CatchDoOrderResult(&ws.WsStreamClient, orderSend)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// 撤单
func (ws *TradeWsStreamClient) CancelOrder(api *OrderCancelAPI) (*WsOrderResult[OrderCancelRes], error) {

	if ws.waitOrderCancelMu.TryLock() {
		defer ws.waitOrderCancelMu.Unlock()
	} else {
		return nil, fmt.Errorf("another order cancel request is executing")
	}

	if ws.waitOrderCancelResult != nil {
		return nil, fmt.Errorf("another order cancel request waiting for result")
	}

	orderSend, err := doOrder[OrderCancelReq, OrderCancelRes](&ws.WsStreamClient, ORDER_CANCEL, *api.req)
	if err != nil {
		return nil, err
	}
	ws.waitOrderCancelResult = orderSend
	result, err := CatchDoOrderResult(&ws.WsStreamClient, orderSend)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// 改单
func (ws *TradeWsStreamClient) AmendOrder(api *OrderAmendAPI) (*WsOrderResult[OrderAmendRes], error) {
	if ws.waitOrderAmendMu.TryLock() {
		defer ws.waitOrderAmendMu.Unlock()
	} else {
		return nil, fmt.Errorf("another order amend request is executing")
	}

	if ws.waitOrderAmendResult != nil {
		return nil, fmt.Errorf("another order amend request waiting for result")
	}

	orderSend, err := doOrder[OrderAmendReq, OrderAmendRes](&ws.WsStreamClient, ORDER_AMEND, *api.req)
	if err != nil {
		return nil, err
	}
	ws.waitOrderAmendResult = orderSend
	result, err := CatchDoOrderResult(&ws.WsStreamClient, orderSend)
	if err != nil {
		return nil, err
	}
	return result, nil
}
