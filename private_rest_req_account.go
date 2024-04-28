package mybybitapi

// accountType	true	string	帳戶類型.
// 統一帳戶: UNIFIED(現貨/USDT和USDC永續/期權), CONTRACT(反向)
// 經典帳戶: CONTRACT(期貨), SPOT(現貨)
// coin	false	string	幣種名稱
// 不傳則返回非零資產信息
// 可以傳多個幣進行查詢，以逗號分隔, USDT,USDC
type AccountWalletBalanceReq struct {
	AccountType *string `json:"accountType"` //String	true	string	帳戶類型. 統一帳戶: UNIFIED(現貨/USDT和USDC永續/期權), CONTRACT(反向) 經典帳戶: CONTRACT(期貨), SPOT(現貨)
	Coin        *string `json:"coin"`        //String	false	string	幣種名稱 不傳則返回非零資產信息 可以傳多個幣進行查詢，以逗號分隔, USDT,USDC
}

type AccountWalletBalanceAPI struct {
	client *PrivateRestClient
	req    *AccountWalletBalanceReq
}

// accountType	true	string	帳戶類型. 統一帳戶: UNIFIED(現貨/USDT和USDC永續/期權), CONTRACT(反向) 經典帳戶: CONTRACT(期貨), SPOT(現貨)
func (api *AccountWalletBalanceAPI) AccountType(accountType string) *AccountWalletBalanceAPI {
	api.req.AccountType = GetPointer(accountType)
	return api
}

// coin	false	string	幣種名稱 不傳則返回非零資產信息 可以傳多個幣進行查詢，以逗號分隔, USDT,USDC
func (api *AccountWalletBalanceAPI) Coin(coin string) *AccountWalletBalanceAPI {
	api.req.Coin = GetPointer(coin)
	return api
}
