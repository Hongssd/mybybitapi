package mybybitapi

func (client *PrivateRestClient) NewSpotMarginTradeSetLeverage() *SpotMarginTradeSetLeverageAPI {
	return &SpotMarginTradeSetLeverageAPI{
		client: client,
		req:    &SpotMarginTradeSetLeverageReq{},
	}
}
func (api *SpotMarginTradeSetLeverageAPI) Do() (*BybitRestRes[SpotMarginTradeSetLeverageRes], error) {
	url := bybitHandlerRequestAPIWithoutPathQueryParam(REST, PrivateRestAPIMap[SpotMarginTradeSetLeverage])
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	return bybitCallAPIWithSecret[SpotMarginTradeSetLeverageRes](api.client.c, url, reqBody, POST)
}

func (client *PrivateRestClient) NewSpotMarginTradeState() *SpotMarginTradeStateAPI {
	return &SpotMarginTradeStateAPI{
		client: client,
		req:    &SpotMarginTradeStateReq{},
	}
}
func (api *SpotMarginTradeStateAPI) Do() (*BybitRestRes[SpotMarginTradeStateRes], error) {
	url := bybitHandlerRequestAPIWithoutPathQueryParam(REST, PrivateRestAPIMap[SpotMarginTradeState])
	return bybitCallAPIWithSecret[SpotMarginTradeStateRes](api.client.c, url, NIL_REQBODY, GET)
}
