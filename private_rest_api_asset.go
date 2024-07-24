package mybybitapi

// bybit AssetTransferQueryInterTransferList PrivateRest接口 GET 查詢劃轉紀錄 (單帳號內)
func (client *PrivateRestClient) NewAssetTransferQueryInterTransferList() *AssetTransferQueryInterTransferListAPI {
	return &AssetTransferQueryInterTransferListAPI{
		client: client,
		req:    &AssetTransferQueryInterTransferListReq{},
	}
}
func (api *AssetTransferQueryInterTransferListAPI) Do() (*BybitRestRes[AssetTransferQueryInterTransferListRes], error) {
	url := bybitHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestAPIMap[AssetTransferQueryInterTransferList])
	return bybitCallAPIWithSecret[AssetTransferQueryInterTransferListRes](api.client.c, url, NIL_REQBODY, GET)
}

// bybit AssetTransferQueryTransferCoinList:  PrivateRest接口 //GET 帳戶類型間可劃轉的幣種
func (client *PrivateRestClient) NewAssetTransferQueryTransferCoinList() *AssetTransferQueryTransferCoinListAPI {
	return &AssetTransferQueryTransferCoinListAPI{
		client: client,
		req:    &AssetTransferQueryTransferCoinListReq{},
	}
}
func (api *AssetTransferQueryTransferCoinListAPI) Do() (*BybitRestRes[AssetTransferQueryTransferCoinListRes], error) {
	url := bybitHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestAPIMap[AssetTransferQueryTransferCoinList])
	return bybitCallAPIWithSecret[AssetTransferQueryTransferCoinListRes](api.client.c, url, NIL_REQBODY, GET)
}

// bybit AssetTransferInterTransfer:  PrivateRest接口 //POST 劃轉 (單帳號內)
func (client *PrivateRestClient) NewAssetTransferInterTransfer() *AssetTransferInterTransferAPI {
	return &AssetTransferInterTransferAPI{
		client: client,
		req:    &AssetTransferInterTransferReq{},
	}
}
func (api *AssetTransferInterTransferAPI) Do() (*BybitRestRes[AssetTransferInterTransferRes], error) {
	url := bybitHandlerRequestAPIWithoutPathQueryParam(REST, PrivateRestAPIMap[AssetTransferInterTransfer])
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	return bybitCallAPIWithSecret[AssetTransferInterTransferRes](api.client.c, url, reqBody, POST)
}

func (client *PrivateRestClient) NewAssetTransferQueryAccountCoinsBalance() *AssetTransferQueryAccountCoinsBalanceAPI {
	return &AssetTransferQueryAccountCoinsBalanceAPI{
		client: client,
		req:    &AssetTransferQueryAccountCoinsBalanceReq{},
	}
}

func (api *AssetTransferQueryAccountCoinsBalanceAPI) Do() (*BybitRestRes[AssetTransferQueryAccountCoinsBalanceRes], error) {
	url := bybitHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestAPIMap[AssetTransferQueryAccountCoinsBalance])
	return bybitCallAPIWithSecret[AssetTransferQueryAccountCoinsBalanceRes](api.client.c, url, NIL_REQBODY, GET)
}

func (client *PrivateRestClient) NewAssetTransferQueryAccountCoinBalance() *AssetTransferQueryAccountCoinBalanceAPI {
	return &AssetTransferQueryAccountCoinBalanceAPI{
		client: client,
		req:    &AssetTransferQueryAccountCoinBalanceReq{},
	}
}

func (api *AssetTransferQueryAccountCoinBalanceAPI) Do() (*BybitRestRes[AssetTransferQueryAccountCoinBalanceRes], error) {
	url := bybitHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestAPIMap[AssetTransferQueryAccountCoinBalance])
	return bybitCallAPIWithSecret[AssetTransferQueryAccountCoinBalanceRes](api.client.c, url, NIL_REQBODY, GET)
}

func (client *PrivateRestClient) NewAssetTithdrawWithdrawableAmount() *AssetTithdrawWithdrawableAmountAPI {
	return &AssetTithdrawWithdrawableAmountAPI{
		client: client,
		req:    &AssetTithdrawWithdrawableAmountReq{},
	}
}

func (api *AssetTithdrawWithdrawableAmountAPI) Do() (*BybitRestRes[AssetTithdrawWithdrawableAmountRes], error) {
	url := bybitHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestAPIMap[AssetTithdrawWithdrawableAmount])
	return bybitCallAPIWithSecret[AssetTithdrawWithdrawableAmountRes](api.client.c, url, NIL_REQBODY, GET)
}
