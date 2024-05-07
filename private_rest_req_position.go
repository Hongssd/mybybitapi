package mybybitapi

type PositionListReq struct {
	Category   *string `json:"category"`   //string	true	統一帳戶: linear,inverse, option 經典帳戶: linear, inverse
	Symbol     *string `json:"symbol"`     //string	false	合約名稱 若傳了symbol, 則不管是否有倉位都返回該symbol數據 若symbol不傳但傳了settleCoin, 則僅返回有實際倉位的數據
	BaseCoin   *string `json:"baseCoin"`   //string	false	交易幣種. 僅option. 若不傳，則返回期權下所有持倉
	SettleCoin *string `json:"settleCoin"` //string	false	結算幣種. 對於USDT和USDC期貨而言，symbol和settleCon必傳其中一個, 若都傳，則symbol有更高的優先級
	Limit      *int    `json:"limit"`      //integer	false	每頁數量限制. [1, 200]. 默認: 20
	Cursor     *string `json:"cursor"`     //string	false	游標，用於翻頁
}

type PositionListAPI struct {
	client *PrivateRestClient
	req    *PositionListReq
}

// category string true 統一帳戶: linear,inverse, option 經典帳戶: linear, inverse
func (api *PositionListAPI) Category(category string) *PositionListAPI {
	api.req.Category = GetPointer(category)
	return api
}

// symbol string false 合約名稱 若傳了symbol, 則不管是否有倉位都返回該symbol數據 若symbol不傳但傳了settleCoin, 則僅返回有實際倉位的數據
func (api *PositionListAPI) Symbol(symbol string) *PositionListAPI {
	api.req.Symbol = GetPointer(symbol)
	return api
}

// baseCoin string false 交易幣種. 僅option. 若不傳，則返回期權下所有持倉
func (api *PositionListAPI) BaseCoin(baseCoin string) *PositionListAPI {
	api.req.BaseCoin = GetPointer(baseCoin)
	return api
}

// settleCoin string false 結算幣種. 對於USDT和USDC期貨而言，symbol和settleCon必傳其中一個, 若都傳，則symbol有更高的優先級
func (api *PositionListAPI) SettleCoin(settleCoin string) *PositionListAPI {
	api.req.SettleCoin = GetPointer(settleCoin)
	return api
}

// limit integer false 每頁數量限制. [1, 200]. 默認: 20
func (api *PositionListAPI) Limit(limit int) *PositionListAPI {
	api.req.Limit = GetPointer(limit)
	return api
}

// cursor string false 游標，用於翻頁
func (api *PositionListAPI) Cursor(cursor string) *PositionListAPI {
	api.req.Cursor = GetPointer(cursor)
	return api
}

type PositionSetLeverageReq struct {
	Category     *string `json:"category"`     //string	true	產品類型 統一帳戶: linear, inverse 經典帳戶: linear, inverse
	Symbol       *string `json:"symbol"`       //string	true	合約名稱
	BuyLeverage  *string `json:"buyLeverage"`  //string	true	[1, 風險限額允許的最大槓桿數] 單倉模式: 經典帳戶和統一帳戶的buyLeverage 必須等於sellLeverage 雙倉模式: 經典帳戶和統一帳戶(逐倉模式)buyLeverage可以與sellLeverage不想等; 統一帳戶(全倉模式)的buyLeverage 必須等於sellLeverage
	SellLeverage *string `json:"sellLeverage"` //string	true	[1, 風險限額允許的最大槓桿數]
}

type PositionSetLeverageAPI struct {
	client *PrivateRestClient
	req    *PositionSetLeverageReq
}

// category string true 產品類型 統一帳戶: linear, inverse 經典帳戶: linear, inverse
func (api *PositionSetLeverageAPI) Category(category string) *PositionSetLeverageAPI {
	api.req.Category = GetPointer(category)
	return api
}

// symbol string true 合約名稱
func (api *PositionSetLeverageAPI) Symbol(symbol string) *PositionSetLeverageAPI {
	api.req.Symbol = GetPointer(symbol)
	return api
}

// buyLeverage string true [1, 風險限額允許的最大槓桿數] 單倉模式: 經典帳戶和統一帳戶的buyLeverage 必須等於sellLeverage 雙倉模式: 經典帳戶和統一帳戶(逐倉模式)buyLeverage可以與sellLeverage不想等; 統一帳戶(全倉模式)的buyLeverage 必須等於sellLeverage
func (api *PositionSetLeverageAPI) BuyLeverage(buyLeverage string) *PositionSetLeverageAPI {
	api.req.BuyLeverage = GetPointer(buyLeverage)
	return api
}

// sellLeverage string true [1, 風險限額允許的最大槓桿數]
func (api *PositionSetLeverageAPI) SellLeverage(sellLeverage string) *PositionSetLeverageAPI {
	api.req.SellLeverage = GetPointer(sellLeverage)
	return api
}

type PositionSwitchIsolatedReq struct {
	Category     *string `json:"category"`     //string	true	產品類型 統一帳戶: inverse 經典帳戶: linear, inverse 這裡category字段不參與業務邏輯，僅做路由使用
	Symbol       *string `json:"symbol"`       //string	true	合約名稱
	TradeMode    *int    `json:"tradeMode"`    //integer	true	0: 全倉. 1: 逐倉
	BuyLeverage  *string `json:"buyLeverage"`  //string	true	買側槓桿倍數. 必須與sellLeverage的值保持相同
	SellLeverage *string `json:"sellLeverage"` //string	true	賣側槓桿倍數. 必須與buyLeverage的值保持相同
}

type PositionSwitchIsolatedAPI struct {
	client *PrivateRestClient
	req    *PositionSwitchIsolatedReq
}

// category string true 產品類型 統一帳戶: inverse 經典帳戶: linear, inverse 這裡category字段不參與業務邏輯，僅做路由使用
func (api *PositionSwitchIsolatedAPI) Category(category string) *PositionSwitchIsolatedAPI {
	api.req.Category = GetPointer(category)
	return api
}

// symbol string true 合約名稱
func (api *PositionSwitchIsolatedAPI) Symbol(symbol string) *PositionSwitchIsolatedAPI {
	api.req.Symbol = GetPointer(symbol)
	return api
}

// tradeMode integer true 0: 全倉. 1: 逐倉
func (api *PositionSwitchIsolatedAPI) TradeMode(tradeMode int) *PositionSwitchIsolatedAPI {
	api.req.TradeMode = GetPointer(tradeMode)
	return api
}

// buyLeverage string true 買側槓桿倍數. 必須與sellLeverage的值保持相同
func (api *PositionSwitchIsolatedAPI) BuyLeverage(buyLeverage string) *PositionSwitchIsolatedAPI {
	api.req.BuyLeverage = GetPointer(buyLeverage)
	return api
}

// sellLeverage string true 賣側槓桿倍數. 必須與buyLeverage的值保持相同
func (api *PositionSwitchIsolatedAPI) SellLeverage(sellLeverage string) *PositionSwitchIsolatedAPI {
	api.req.SellLeverage = GetPointer(sellLeverage)
	return api
}

type PositionSwitchModeReq struct {
	Category *string `json:"category"` //string	true	產品類型 統一帳戶: linear: USDT永續; inverse: 反向交割 經典帳戶: linear, inverse. 這裡category字段不參與業務邏輯，僅做路由使用
	Symbol   *string `json:"symbol"`   //string	false	合約名稱. symbol和coin必須傳其中一個. symbol有更高優先級
	Coin     *string `json:"coin"`     //string	false	結算幣種
	Mode     *int    `json:"mode"`     //integer	true	倉位模式. 0: 單向持倉. 3: 雙向持倉
}

type PositionSwitchModeAPI struct {
	client *PrivateRestClient
	req    *PositionSwitchModeReq
}

// category string true 產品類型 統一帳戶: linear: USDT永續; inverse: 反向交割 經典帳戶: linear, inverse. 這裡category字段不參與業務邏輯，僅做路由使用
func (api *PositionSwitchModeAPI) Category(category string) *PositionSwitchModeAPI {
	api.req.Category = GetPointer(category)
	return api
}

// symbol string false 合約名稱. symbol和coin必須傳其中一個. symbol有更高優先級
func (api *PositionSwitchModeAPI) Symbol(symbol string) *PositionSwitchModeAPI {
	api.req.Symbol = GetPointer(symbol)
	return api
}

// coin string false 結算幣種
func (api *PositionSwitchModeAPI) Coin(coin string) *PositionSwitchModeAPI {
	api.req.Coin = GetPointer(coin)
	return api
}

// mode integer true 倉位模式. 0: 單向持倉. 3: 雙向持倉
func (api *PositionSwitchModeAPI) Mode(mode int) *PositionSwitchModeAPI {
	api.req.Mode = GetPointer(mode)
	return api
}
