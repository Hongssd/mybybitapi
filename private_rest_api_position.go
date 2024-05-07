package mybybitapi

// bybit PositionList PrivateRest接口 GET 查詢持倉 (實時)
func (client *PrivateRestClient) NewPositionList() *PositionListAPI {
	return &PositionListAPI{
		client: client,
		req:    &PositionListReq{},
	}
}
func (api *PositionListAPI) Do() (*BybitRestRes[PositionListRes], error) {
	url := bybitHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestAPIMap[PositionList])
	return bybitCallAPIWithSecret[PositionListRes](api.client.c, url, NIL_REQBODY, GET)
}

// bybit PositionSetLeverage PrivateRest接口 POST 設置持倉槓桿
func (client *PrivateRestClient) NewPositionSetLeverage() *PositionSetLeverageAPI {
	return &PositionSetLeverageAPI{
		client: client,
		req:    &PositionSetLeverageReq{},
	}
}
func (api *PositionSetLeverageAPI) Do() (*BybitRestRes[PositionSetLeverageRes], error) {
	url := bybitHandlerRequestAPIWithoutPathQueryParam(REST, PrivateRestAPIMap[PositionSetLeverage])
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	return bybitCallAPIWithSecret[PositionSetLeverageRes](api.client.c, url, reqBody, POST)
}

// bybit PositionSwitchIsolated PrivateRest接口 POST 切換全倉/逐倉保證金(交易對)
func (client *PrivateRestClient) NewPositionSwitchIsolated() *PositionSwitchIsolatedAPI {
	return &PositionSwitchIsolatedAPI{
		client: client,
		req:    &PositionSwitchIsolatedReq{},
	}
}
func (api *PositionSwitchIsolatedAPI) Do() (*BybitRestRes[PositionSwitchIsolatedRes], error) {
	url := bybitHandlerRequestAPIWithoutPathQueryParam(REST, PrivateRestAPIMap[PositionSwitchIsolated])
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	return bybitCallAPIWithSecret[PositionSwitchIsolatedRes](api.client.c, url, reqBody, POST)
}

// bybit PositionSwitchMode PrivateRest接口 POST 切換持倉模式
func (client *PrivateRestClient) NewPositionSwitchMode() *PositionSwitchModeAPI {
	return &PositionSwitchModeAPI{
		client: client,
		req:    &PositionSwitchModeReq{},
	}
}
func (api *PositionSwitchModeAPI) Do() (*BybitRestRes[PositionSwitchModeRes], error) {
	url := bybitHandlerRequestAPIWithoutPathQueryParam(REST, PrivateRestAPIMap[PositionSwitchMode])
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	return bybitCallAPIWithSecret[PositionSwitchModeRes](api.client.c, url, reqBody, POST)
}
