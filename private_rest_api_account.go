package mybybitapi

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
