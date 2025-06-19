package mybybitapi

type AccountInfoReq struct {
}

type AccountInfoAPI struct {
	client *PrivateRestClient
	req    *AccountInfoReq
}

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

type AccountUpgradeToUtaReq struct {
}

type AccountUpgradeToUtaAPI struct {
	client *PrivateRestClient
	req    *AccountUpgradeToUtaReq
}

// setMarginMode	true	string	ISOLATED_MARGIN(逐倉保證金模式), REGULAR_MARGIN（全倉保證金模式）PORTFOLIO_MARGIN（組合保證金模式）默認常規，傳常規則返回設置成功
type AccountSetMarginModeReq struct {
	SetMarginMode *string `json:"setMarginMode"` //true	ISOLATED_MARGIN(逐倉保證金模式), REGULAR_MARGIN（全倉保證金模式）PORTFOLIO_MARGIN（組合保證金模式）默認常規，傳常規則返回設置成功
}

type AccountSetMarginModeAPI struct {
	client *PrivateRestClient
	req    *AccountSetMarginModeReq
}

// setMarginMode true string ISOLATED_MARGIN(逐倉保證金模式), REGULAR_MARGIN（全倉保證金模式）PORTFOLIO_MARGIN（組合保證金模式）默認常規，傳常規則返回設置成功
func (api *AccountSetMarginModeAPI) SetMarginMode(setMarginMode string) *AccountSetMarginModeAPI {
	api.req.SetMarginMode = GetPointer(setMarginMode)
	return api
}

type AccountWithdrawalAPI struct {
	client *PrivateRestClient
	req    *AccountWithdrawalReq
}

// coinName	true	string	幣種名稱, 僅大寫. 支持最多20個幣種批量查詢, 用逗號分隔. BTC,SOL,USDT,USDC
type AccountWithdrawalReq struct {
	CoinName *string `json:"coinName"` //string	true	幣種名稱, 僅大寫. 支持最多20個幣種批量查詢, 用逗號分隔. BTC,SOL,USDT,USDC
}

// coinName string true 幣種名稱, 僅大寫. 支持最多20個幣種批量查詢, 用逗號分隔. BTC,SOL,USDT,USDC
func (api *AccountWithdrawalAPI) CoinName(coinName string) *AccountWithdrawalAPI {
	api.req.CoinName = GetPointer(coinName)
	return api
}

type AccountSetCollateralSwitchAPI struct {
	client *PrivateRestClient
	req    *AccountSetCollateralSwitchReq
}

type AccountSetCollateralSwitchReq struct {
	Coin             *string `json:"coin"`             //string	true	幣種名稱
	CollateralSwitch *string `json:"collateralSwitch"` //string	true	ON: 開啟抵押, OFF: 關閉抵押
}

// coin string true 幣種名稱
func (api *AccountSetCollateralSwitchAPI) Coin(coin string) *AccountSetCollateralSwitchAPI {
	api.req.Coin = GetPointer(coin)
	return api
}

// collateralSwitch string true ON: 開啟抵押, OFF: 關閉抵押
func (api *AccountSetCollateralSwitchAPI) CollateralSwitch(collateralSwitch string) *AccountSetCollateralSwitchAPI {
	api.req.CollateralSwitch = GetPointer(collateralSwitch)
	return api
}

type AccountSetCollateralSwitchBatchAPI struct {
	client *PrivateRestClient
	req    *AccountSetCollateralSwitchBatchReq
}

type AccountSetCollateralSwitchBatchReq struct {
	Request []AccountSetCollateralSwitchReq `json:"request"`
}

func (api *AccountSetCollateralSwitchBatchAPI) AddNewSetCollateralSwitchReq(coin string, collateralSwitch string) *AccountSetCollateralSwitchBatchAPI {
	if api.req == nil {
		api.req = &AccountSetCollateralSwitchBatchReq{
			Request: []AccountSetCollateralSwitchReq{},
		}
	}
	api.req.Request = append(api.req.Request, AccountSetCollateralSwitchReq{
		Coin:             GetPointer(coin),
		CollateralSwitch: GetPointer(collateralSwitch),
	})
	return api
}

func (api *AccountSetCollateralSwitchBatchAPI) SetSetCollateralSwitchReqList(setCollateralSwitchReqList []AccountSetCollateralSwitchReq) *AccountSetCollateralSwitchBatchAPI {
	if len(setCollateralSwitchReqList) == 0 {
		return api
	}
	api.req.Request = setCollateralSwitchReqList
	return api
}

type AccountCollateralInfoAPI struct {
	client *PrivateRestClient
	req    *AccountCollateralInfoReq
}

type AccountCollateralInfoReq struct {
	Currency *string `json:"currency"` //string false 目前所有抵押品的資產幣種
}

// currency string false 目前所有抵押品的資產幣種
func (api *AccountCollateralInfoAPI) Currency(currency string) *AccountCollateralInfoAPI {
	api.req.Currency = GetPointer(currency)
	return api
}

