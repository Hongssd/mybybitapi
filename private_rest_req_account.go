package mybybitapi

type AccountWalletBalanceReq struct {
	AccountType *string `json:"accountType"` //string	true	帳戶類型. 統一帳戶: UNIFIED(現貨/USDT和USDC永續/期權), CONTRACT(反向) 經典帳戶: CONTRACT(期貨), SPOT(現貨)
	Coin        *string `json:"coin"`        //string	false	幣種名稱 不傳則返回非零資產信息 可以傳多個幣進行查詢，以逗號分隔, USDT,USDC
}

type AccountWalletBalanceAPI struct {
	client *PrivateRestClient
	req    *AccountWalletBalanceReq
}

// accountType string true 帳戶類型. 統一帳戶: UNIFIED(現貨/USDT和USDC永續/期權), CONTRACT(反向) 經典帳戶: CONTRACT(期貨), SPOT(現貨)
func (api *AccountWalletBalanceAPI) AccountType(accountType string) *AccountWalletBalanceAPI {
	api.req.AccountType = GetPointer(accountType)
	return api
}

// coin string false 幣種名稱 不傳則返回非零資產信息 可以傳多個幣進行查詢，以逗號分隔, USDT,USDC
func (api *AccountWalletBalanceAPI) Coin(coin string) *AccountWalletBalanceAPI {
	api.req.Coin = GetPointer(coin)
	return api
}

type AccountFeeRateReq struct {
	Category *string `json:"category"` //string	true	產品類型. spot, linear, inverse, option
	Symbol   *string `json:"symbol"`   //string	false	合約名稱. 僅spot, linear, inverse有效
	BaseCoin *string `json:"baseCoin"` //string	false	交易幣種. SOL, BTC, ETH.僅option有效
}

type AccountFeeRateAPI struct {
	client *PrivateRestClient
	req    *AccountFeeRateReq
}

// category string true 產品類型. spot, linear, inverse, option
func (api *AccountFeeRateAPI) Category(category string) *AccountFeeRateAPI {
	api.req.Category = GetPointer(category)
	return api
}

// symbol string false 合約名稱. 僅spot, linear, inverse有效
func (api *AccountFeeRateAPI) Symbol(symbol string) *AccountFeeRateAPI {
	api.req.Symbol = GetPointer(symbol)
	return api
}

// baseCoin string false 交易幣種. SOL, BTC, ETH.僅option有效
func (api *AccountFeeRateAPI) BaseCoin(baseCoin string) *AccountFeeRateAPI {
	api.req.BaseCoin = GetPointer(baseCoin)
	return api
}
