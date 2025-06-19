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

// bybit AccountWithdrawal PrivateRest接口 GET 查詢可劃轉餘額(统一账户)
func (client *PrivateRestClient) NewAccountWithdrawal() *AccountWithdrawalAPI {
	return &AccountWithdrawalAPI{
		client: client,
		req:    &AccountWithdrawalReq{},
	}
}
func (api *AccountWithdrawalAPI) Do() (*BybitRestRes[AccountWithdrawalRes], error) {
	url := bybitHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestAPIMap[AccountWithdrawal])
	return bybitCallAPIWithSecret[AccountWithdrawalRes](api.client.c, url, NIL_REQBODY, GET)
}

// bybit AccountSetCollateralSwitch PrivateRest接口 POST 設置抵押品幣種
func (client *PrivateRestClient) NewAccountSetCollateralSwitch() *AccountSetCollateralSwitchAPI {
	return &AccountSetCollateralSwitchAPI{
		client: client,
		req:    &AccountSetCollateralSwitchReq{},
	}
}
func (api *AccountSetCollateralSwitchAPI) Do() (*BybitRestRes[AccountSetCollateralSwitchRes], error) {
	url := bybitHandlerRequestAPIWithoutPathQueryParam(REST, PrivateRestAPIMap[AccountSetCollateralSwitch])
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	return bybitCallAPIWithSecret[AccountSetCollateralSwitchRes](api.client.c, url, reqBody, POST)
}

// bybit AccountSetCollateralSwitchBatch PrivateRest接口 POST 批量設置抵押品幣種
func (client *PrivateRestClient) NewAccountSetCollateralSwitchBatch() *AccountSetCollateralSwitchBatchAPI {
	return &AccountSetCollateralSwitchBatchAPI{
		client: client,
		req:    &AccountSetCollateralSwitchBatchReq{},
	}
}
func (api *AccountSetCollateralSwitchBatchAPI) Do() (*BybitRestRes[AccountSetCollateralSwitchBatchRes], error) {
	url := bybitHandlerRequestAPIWithoutPathQueryParam(REST, PrivateRestAPIMap[AccountSetCollateralSwitchBatch])
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	return bybitCallAPIWithSecret[AccountSetCollateralSwitchBatchRes](api.client.c, url, reqBody, POST)
}

// bybit AccountCollateralInfo PrivateRest接口 GET 查詢抵押品幣種
func (client *PrivateRestClient) NewAccountCollateralInfo() *AccountCollateralInfoAPI {
	return &AccountCollateralInfoAPI{
		client: client,
		req:    &AccountCollateralInfoReq{},
	}
}
func (api *AccountCollateralInfoAPI) Do() (*BybitRestRes[AccountCollateralInfoRes], error) {
	url := bybitHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestAPIMap[AccountCollateralInfo])
	return bybitCallAPIWithSecret[AccountCollateralInfoRes](api.client.c, url, NIL_REQBODY, GET)
}