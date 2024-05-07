package mybybitapi

// bybit OrderCreate PrivateRest接口 POST 創建委託單
func (client *PrivateRestClient) NewOrderCreate() *OrderCreateAPI {
	return &OrderCreateAPI{
		client: client,
		req:    &OrderCreateReq{},
	}
}
func (api *OrderCreateAPI) Do() (*BybitRestRes[OrderCreateRes], error) {
	url := bybitHandlerRequestAPIWithoutPathQueryParam(REST, PrivateRestAPIMap[OrderCreate])
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	return bybitCallAPIWithSecret[OrderCreateRes](api.client.c, url, reqBody, POST)
}

// bybit OrderCreateBatch PrivateRest接口 POST 批量創建委託單
func (client *PrivateRestClient) NewOrderCreateBatch() *OrderCreateBatchAPI {
	return &OrderCreateBatchAPI{
		client: client,
		req:    nil,
	}
}
func (api *OrderCreateBatchAPI) Do() (*BybitRestRes[OrderCreateBatchRes], error) {
	url := bybitHandlerRequestAPIWithoutPathQueryParam(REST, PrivateRestAPIMap[OrderCreateBatch])
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	return bybitCallAPIWithSecret[OrderCreateBatchRes](api.client.c, url, reqBody, POST)
}

// bybit OrderAmend PrivateRest接口 POST 修改委託單
func (client *PrivateRestClient) NewOrderAmend() *OrderAmendAPI {
	return &OrderAmendAPI{
		client: client,
		req:    &OrderAmendReq{},
	}
}
func (api *OrderAmendAPI) Do() (*BybitRestRes[OrderAmendRes], error) {
	url := bybitHandlerRequestAPIWithoutPathQueryParam(REST, PrivateRestAPIMap[OrderAmend])
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	return bybitCallAPIWithSecret[OrderAmendRes](api.client.c, url, reqBody, POST)
}

// bybit OrderAmendBatch PrivateRest接口 POST 批量修改委託單
func (client *PrivateRestClient) NewOrderAmendBatch() *OrderAmendBatchAPI {
	return &OrderAmendBatchAPI{
		client: client,
		req:    nil,
	}
}
func (api *OrderAmendBatchAPI) Do() (*BybitRestRes[OrderAmendBatchRes], error) {
	url := bybitHandlerRequestAPIWithoutPathQueryParam(REST, PrivateRestAPIMap[OrderAmendBatch])
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	return bybitCallAPIWithSecret[OrderAmendBatchRes](api.client.c, url, reqBody, POST)
}

// bybit OrderCancel PrivateRest接口 POST 撤銷委託單
func (client *PrivateRestClient) NewOrderCancel() *OrderCancelAPI {
	return &OrderCancelAPI{
		client: client,
		req:    &OrderCancelReq{},
	}
}
func (api *OrderCancelAPI) Do() (*BybitRestRes[OrderCancelRes], error) {
	url := bybitHandlerRequestAPIWithoutPathQueryParam(REST, PrivateRestAPIMap[OrderCancel])
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	return bybitCallAPIWithSecret[OrderCancelRes](api.client.c, url, reqBody, POST)
}

// bybit OrderCancelBatch PrivateRest接口 POST 批量撤銷委託單
func (client *PrivateRestClient) NewOrderCancelBatch() *OrderCancelBatchAPI {
	return &OrderCancelBatchAPI{
		client: client,
		req:    nil,
	}
}
func (api *OrderCancelBatchAPI) Do() (*BybitRestRes[OrderCancelBatchRes], error) {
	url := bybitHandlerRequestAPIWithoutPathQueryParam(REST, PrivateRestAPIMap[OrderCancelBatch])
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	return bybitCallAPIWithSecret[OrderCancelBatchRes](api.client.c, url, reqBody, POST)
}

// bybit OrderCancelAll PrivateRest接口 POST 撤銷所有訂單
func (client *PrivateRestClient) NewOrderCancelAll() *OrderCancelAllAPI {
	return &OrderCancelAllAPI{
		client: client,
		req:    &OrderCancelAllReq{},
	}
}
func (api *OrderCancelAllAPI) Do() (*BybitRestRes[OrderCancelAllRes], error) {
	url := bybitHandlerRequestAPIWithoutPathQueryParam(REST, PrivateRestAPIMap[OrderCancelAll])
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	return bybitCallAPIWithSecret[OrderCancelAllRes](api.client.c, url, reqBody, POST)
}

// bybit OrderRealtime PrivateRest接口 GET 查詢實時委託
func (client *PrivateRestClient) NewOrderRealtime() *OrderRealtimeAPI {
	return &OrderRealtimeAPI{
		client: client,
		req:    &OrderRealtimeReq{},
	}
}
func (api *OrderRealtimeAPI) Do() (*BybitRestRes[OrderRealtimeRes], error) {
	url := bybitHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestAPIMap[OrderRealtime])
	return bybitCallAPIWithSecret[OrderRealtimeRes](api.client.c, url, NIL_REQBODY, GET)
}

// bybit OrderHistory PrivateRest接口 GET 查詢歷史訂單 (2年)
func (client *PrivateRestClient) NewOrderHistory() *OrderHistoryAPI {
	return &OrderHistoryAPI{
		client: client,
		req:    &OrderHistoryReq{},
	}
}
func (api *OrderHistoryAPI) Do() (*BybitRestRes[OrderHistoryRes], error) {
	url := bybitHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestAPIMap[OrderHistory])
	return bybitCallAPIWithSecret[OrderHistoryRes](api.client.c, url, NIL_REQBODY, GET)
}

// bybit OrderExecutionList PrivateRest接口 GET 查詢成交紀錄 (2年)
func (client *PrivateRestClient) NewOrderExecutionList() *OrderExecutionListAPI {
	return &OrderExecutionListAPI{
		client: client,
		req:    &OrderExecutionListReq{},
	}
}
func (api *OrderExecutionListAPI) Do() (*BybitRestRes[OrderExecutionListRes], error) {
	url := bybitHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestAPIMap[OrderExecutionList])
	return bybitCallAPIWithSecret[OrderExecutionListRes](api.client.c, url, NIL_REQBODY, GET)
}
