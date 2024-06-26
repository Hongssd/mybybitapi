package mybybitapi

// bybit AccountInfo PrivateRest接口 GET 查詢帳戶信息
func (client *PrivateRestClient) NewAccountInfo() *AccountInfoAPI {
	return &AccountInfoAPI{
		client: client,
		req:    &AccountInfoReq{},
	}
}
func (api *AccountInfoAPI) Do() (*BybitRestRes[AccountInfoRes], error) {
	url := bybitHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestAPIMap[AccountInfo])
	return bybitCallAPIWithSecret[AccountInfoRes](api.client.c, url, NIL_REQBODY, GET)
}

// bybit AccountWalletBalance PrivateRest接口 GET 查詢錢包餘額
func (client *PrivateRestClient) NewAccountWalletBalance() *AccountWalletBalanceAPI {
	return &AccountWalletBalanceAPI{
		client: client,
		req:    &AccountWalletBalanceReq{},
	}
}
func (api *AccountWalletBalanceAPI) Do() (*BybitRestRes[AccountWalletBalanceRes], error) {
	url := bybitHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestAPIMap[AccountWalletBalance])
	return bybitCallAPIWithSecret[AccountWalletBalanceRes](api.client.c, url, NIL_REQBODY, GET)
}

// bybit AccountFeeRate PrivateRest接口 GET 查詢手續費率
func (client *PrivateRestClient) NewAccountFeeRate() *AccountFeeRateAPI {
	return &AccountFeeRateAPI{
		client: client,
		req:    &AccountFeeRateReq{},
	}
}
func (api *AccountFeeRateAPI) Do() (*BybitRestRes[AccountFeeRateRes], error) {
	url := bybitHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestAPIMap[AccountFeeRate])
	return bybitCallAPIWithSecret[AccountFeeRateRes](api.client.c, url, NIL_REQBODY, GET)
}

// bybit AccountUpgradeToUta PrivateRest接口 POST 升級為UTA帳戶
func (client *PrivateRestClient) NewAccountUpgradeToUta() *AccountUpgradeToUtaAPI {
	return &AccountUpgradeToUtaAPI{
		client: client,
		req:    &AccountUpgradeToUtaReq{},
	}
}
func (api *AccountUpgradeToUtaAPI) Do() (*BybitRestRes[AccountUpgradeToUtaRes], error) {
	url := bybitHandlerRequestAPIWithoutPathQueryParam(REST, PrivateRestAPIMap[AccountUpgradeToUta])
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	return bybitCallAPIWithSecret[AccountUpgradeToUtaRes](api.client.c, url, reqBody, POST)
}

// bybit AccountSetMarginMode PrivateRest接口 POST 設置保證金模式(帳戶)
func (client *PrivateRestClient) NewAccountSetMarginMode() *AccountSetMarginModeAPI {
	return &AccountSetMarginModeAPI{
		client: client,
		req:    &AccountSetMarginModeReq{},
	}
}
func (api *AccountSetMarginModeAPI) Do() (*BybitRestRes[AccountSetMarginModeRes], error) {
	url := bybitHandlerRequestAPIWithoutPathQueryParam(REST, PrivateRestAPIMap[AccountSetMarginMode])
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	return bybitCallAPIWithSecret[AccountSetMarginModeRes](api.client.c, url, reqBody, POST)
}
