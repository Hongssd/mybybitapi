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
