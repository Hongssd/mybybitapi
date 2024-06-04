package mybybitapi

type OrderCreateReq struct {
	Category         *string `json:"category"`                   //string	true	產品類型 統一帳戶: spot, linear, inverse, option 經典帳戶: spot, linear, inverse
	Symbol           *string `json:"symbol"`                     //string	true	合約名稱
	IsLeverage       *int    `json:"isLeverage,omitempty"`       //integer	false	是否借貸. 僅統一帳戶的現貨交易有效. 0(default): 否，則是幣幣訂單, 1: 是，則是槓桿訂單
	Side             *string `json:"side"`                       //string	true	Buy, Sell
	OrderType        *string `json:"orderType"`                  //string	true	訂單類型. Market, Limit
	Qty              *string `json:"qty"`                        //string	true	訂單數量 統一帳戶 現貨: 可以通過設置marketUnit來表示市價單qty的單位, 市價買單默認是quoteCoin, 市價賣單默認是baseCoin 期貨和期權: 總是以base coin作為qty的單位 經典帳戶 現貨: 市價買單的qty總是以quote coin為單位, 其他情況下, qty都是以base coin為單位 期貨: qty總是以base coin為單位 期貨: 如果傳入qty="0"以及reduceOnly="true", 則可以全部平掉對應合約的倉位
	MarketUnit       *string `json:"marketUnit,omitempty"`       //string	false	統一帳戶現貨交易創建市價單時給入參qty指定的單位. 当前不支持止盈/止损和条件单 baseCoin: 比如, 買BTCUSDT, 則"qty"的單位是BTC quoteCoin: 比如, 賣BTCUSDT, 則"qty"的單位是USDT
	Price            *string `json:"price,omitempty"`            //string	false	訂單價格. 市價單將會忽視該字段 請通過該接口確認最低價格和精度要求 如果有持倉, 確保價格優於強平價格
	TriggerDirection *int    `json:"triggerDirection,omitempty"` //integer	false	條件單參數. 用於辨別期望的方向. 1: 當市場價上漲到了triggerPrice時觸發條件單 2: 當市場價下跌到了triggerPrice時觸發條件單 對linear和inverse有效
	OrderFilter      *string `json:"orderFilter,omitempty"`      //string	false	指定訂單品種. Order: 普通單,tpslOrder: 止盈止損單,StopOrder: 條件單. 若不傳, 默認Order 僅對現貨有效
	TriggerPrice     *string `json:"triggerPrice,omitempty"`     //string	false	對於期貨, 是條件單觸發價格參數. 若您希望市場價是要上升後觸發, 確保: triggerPrice > 市場價格 否則, triggerPrice < 市場價格 對於現貨, 這是下止盈止損單或者條件單的觸發價格參數
	TriggerBy        *string `json:"triggerBy,omitempty"`        //string	false	條件單參數. 觸發價格類型. LastPrice, IndexPrice, MarkPrice 僅對linear和inverse有效
	OrderIv          *string `json:"orderIv,omitempty"`          //string	false	隱含波動率. 僅option有效. 按照實際值傳入, e.g., 對於10%, 則傳入0.1. orderIv比price有更高的優先級
	TimeInForce      *string `json:"timeInForce,omitempty"`      //string	false	訂單執行策略 市價單，系統直接使用IOC 若不傳，默認使用GTC
	PositionIdx      *int    `json:"positionIdx,omitempty"`      //integer	false	倉位標識, 用戶不同倉位模式. 該字段對於雙向持倉模式(僅USDT永續和反向期貨有雙向模式)是必傳: 0: 單向持倉 1: 買側雙向持倉 2: 賣側雙向持倉 僅對linear和inverse有效
	OrderLinkId      *string `json:"orderLinkId,omitempty"`      //string	false	用戶自定義訂單Id. category=option時，該參數必傳
	TakeProfit       *string `json:"takeProfit,omitempty"`       //string	false	止盈價格 linear & inverse: 支援統一帳戶和經典帳戶 spot: 僅支持統一帳戶, 創建限價單時, 可以附帶市價止盈止損和限價止盈止損
	StopLoss         *string `json:"stopLoss,omitempty"`         //string	false	止損價格 linear & inverse: 支援統一帳戶和經典帳戶 spot: 僅支持統一帳戶, 創建限價單時, 可以附帶市價止盈止損和限價止盈止損
	TpTriggerBy      *string `json:"tpTriggerBy,omitempty"`      //string	false	觸發止盈的價格類型. MarkPrice, IndexPrice, 默認:LastPrice 僅對linear和inverse有效
	SlTriggerBy      *string `json:"slTriggerBy,omitempty"`      //string	false	觸發止損的價格類型. MarkPrice, IndexPrice, 默認:LastPrice 僅對linear和inverse有效
	ReduceOnly       *bool   `json:"reduceOnly,omitempty"`       //boolean	false	什麼是只減倉? true 將這筆訂單設為只減倉 當減倉時, reduceOnly=true必傳 只減倉單的止盈止損不生效 對linear, inverse和option有效
	CloseOnTrigger   *bool   `json:"closeOnTrigger,omitempty"`   //boolean	false	什麼是觸發後平倉委託?此選項可以確保您的止損單被用於減倉（平倉）而非加倉，並且在可用保證金不足的情況下，取消其他委託，騰出保證金以確保平倉委託的執行. 僅對linear和inverse有效
	SmpType          *string `json:"smpType,omitempty"`          //string	false	Smp執行類型. 什麼是SMP?
	Mmp              *bool   `json:"mmp,omitempty"`              //boolean	false	做市商保護. 僅option有效. true 表示該訂單是做市商保護訂單. 什麼是做市商保護?
	TpslMode         *string `json:"tpslMode,omitempty"`         //string	false	止盈止損模式 Full: 全部倉位止盈止損. 此時, tpOrderType或者slOrderType必須傳Market Partial: 部分倉位止盈止損(下單時沒有size選項, 實際上創建tpsl訂單時, 是按照實際成交的數量來生成止盈止損). 支持創建限價止盈止損. 注意: 創建限價止盈止損時, tpslMode必傳且為Partial 僅對linear和inverse有效
	TpLimitPrice     *string `json:"tpLimitPrice,omitempty"`     //string	false	觸發止盈後轉換為限價單的價格 linear & inverse: 僅作用於當tpslMode=Partial以及tpOrderType=Limit時 spot(UTA): 參數必傳當創建訂單時帶了takeProfit 和 tpOrderType=Limit
	SlLimitPrice     *string `json:"slLimitPrice,omitempty"`     //string	false	觸發止損後轉換為限價單的價格 linear & inverse: 僅作用於當tpslMode=Partial以及slOrderType=Limit時 spot(UTA): 參數必傳當創建訂單時帶了stopLoss 和 slOrderType=Limit
	TpOrderType      *string `json:"tpOrderType,omitempty"`      //string	false	止盈觸發後的訂單類型 linear & inverse: Market(默認), Limit. 對於tpslMode=Full, 僅支持tpOrderType=Market spot(UTA): Market: 當帶了參數"takeProfit", Limit: 當帶了參數"takeProfit" 和 "tpLimitPrice"
	SlOrderType      *string `json:"slOrderType,omitempty"`      //string	false	止損觸發後的訂單類型 linear & inverse: Market(默認), Limit. 對於tpslMode=Full, 僅支持slOrderType=Market spot(UTA): Market: 當帶了參數"stopLoss", Limit: 當帶了參數"stopLoss" 和 "slLimitPrice"
}

type OrderCreateAPI struct {
	client *PrivateRestClient
	req    *OrderCreateReq
}

// category string true 產品類型 統一帳戶: spot, linear, inverse, option 經典帳戶: spot, linear, inverse
func (api *OrderCreateAPI) Category(category string) *OrderCreateAPI {
	api.req.Category = GetPointer(category)
	return api
}

// symbol string true 合約名稱
func (api *OrderCreateAPI) Symbol(symbol string) *OrderCreateAPI {
	api.req.Symbol = GetPointer(symbol)
	return api
}

// isLeverage integer false 是否借貸. 僅統一帳戶的現貨交易有效. 0(default): 否，則是幣幣訂單, 1: 是，則是槓桿訂單
func (api *OrderCreateAPI) IsLeverage(isLeverage int) *OrderCreateAPI {
	api.req.IsLeverage = GetPointer(isLeverage)
	return api
}

// side string true Buy, Sell
func (api *OrderCreateAPI) Side(side string) *OrderCreateAPI {
	api.req.Side = GetPointer(side)
	return api
}

// orderType string true 訂單類型. Market, Limit
func (api *OrderCreateAPI) OrderType(orderType string) *OrderCreateAPI {
	api.req.OrderType = GetPointer(orderType)
	return api
}

// qty string true 訂單數量 統一帳戶 現貨: 可以通過設置marketUnit來表示市價單qty的單位, 市價買單默認是quoteCoin, 市價賣單默認是baseCoin 期貨和期權: 總是以base coin作為qty的單位 經典帳戶 現貨: 市價買單的qty總是以quote coin為單位, 其他情況下, qty都是以base coin為單位 期貨: qty總是以base coin為單位 期貨: 如果傳入qty="0"以及reduceOnly="true", 則可以全部平掉對應合約的倉位
func (api *OrderCreateAPI) Qty(qty string) *OrderCreateAPI {
	api.req.Qty = GetPointer(qty)
	return api
}

// marketUnit string false 統一帳戶現貨交易創建市價單時給入參qty指定的單位. 当前不支持止盈/止损和条件单 baseCoin: 比如, 買BTCUSDT, 則"qty"的單位是BTC quoteCoin: 比如, 賣BTCUSDT, 則"qty"的單位是USDT
func (api *OrderCreateAPI) MarketUnit(marketUnit string) *OrderCreateAPI {
	api.req.MarketUnit = GetPointer(marketUnit)
	return api
}

// price string false 訂單價格. 市價單將會忽視該字段 請通過該接口確認最低價格和精度要求 如果有持倉, 確保價格優於強平價格
func (api *OrderCreateAPI) Price(price string) *OrderCreateAPI {
	api.req.Price = GetPointer(price)
	return api
}

// triggerDirection integer false 條件單參數. 用於辨別期望的方向. 1: 當市場價上漲到了triggerPrice時觸發條件單 2: 當市場價下跌到了triggerPrice時觸發條件單 對linear和inverse有效
func (api *OrderCreateAPI) TriggerDirection(triggerDirection int) *OrderCreateAPI {
	api.req.TriggerDirection = GetPointer(triggerDirection)
	return api
}

// orderFilter string false 指定訂單品種. Order: 普通單,tpslOrder: 止盈止損單,StopOrder: 條件單. 若不傳, 默認Order 僅對現貨有效
func (api *OrderCreateAPI) OrderFilter(orderFilter string) *OrderCreateAPI {
	api.req.OrderFilter = GetPointer(orderFilter)
	return api
}

// triggerPrice string false 對於期貨, 是條件單觸發價格參數. 若您希望市場價是要上升後觸發, 確保: triggerPrice > 市場價格 否則, triggerPrice < 市場價格 對於現貨, 這是下止盈止損單或者條件單的觸發價格參數
func (api *OrderCreateAPI) TriggerPrice(triggerPrice string) *OrderCreateAPI {
	api.req.TriggerPrice = GetPointer(triggerPrice)
	return api
}

// triggerBy string false 條件單參數. 觸發價格類型. LastPrice, IndexPrice, MarkPrice 僅對linear和inverse有效
func (api *OrderCreateAPI) TriggerBy(triggerBy string) *OrderCreateAPI {
	api.req.TriggerBy = GetPointer(triggerBy)
	return api
}

// orderIv string false 隱含波動率. 僅option有效. 按照實際值傳入, e.g., 對於10%, 則傳入0.1. orderIv比price有更高的優先級
func (api *OrderCreateAPI) OrderIv(orderIv string) *OrderCreateAPI {
	api.req.OrderIv = GetPointer(orderIv)
	return api
}

// timeInForce string false 訂單執行策略 市價單，系統直接使用IOC 若不傳，默認使用GTC
func (api *OrderCreateAPI) TimeInForce(timeInForce string) *OrderCreateAPI {
	api.req.TimeInForce = GetPointer(timeInForce)
	return api
}

// positionIdx integer false 倉位標識, 用戶不同倉位模式. 該字段對於雙向持倉模式(僅USDT永續和反向期貨有雙向模式)是必傳: 0: 單向持倉 1: 買側雙向持倉 2: 賣側雙向持倉 僅對linear和inverse有效
func (api *OrderCreateAPI) PositionIdx(positionIdx int) *OrderCreateAPI {
	api.req.PositionIdx = GetPointer(positionIdx)
	return api
}

// orderLinkId string false 用戶自定義訂單Id. category=option時，該參數必傳
func (api *OrderCreateAPI) OrderLinkId(orderLinkId string) *OrderCreateAPI {
	api.req.OrderLinkId = GetPointer(orderLinkId)
	return api
}

// takeProfit string false 止盈價格 linear & inverse: 支援統一帳戶和經典帳戶 spot: 僅支持統一帳戶, 創建限價單時, 可以附帶市價止盈止損和限價止盈止損
func (api *OrderCreateAPI) TakeProfit(takeProfit string) *OrderCreateAPI {
	api.req.TakeProfit = GetPointer(takeProfit)
	return api
}

// stopLoss string false 止損價格 linear & inverse: 支援統一帳戶和經典帳戶 spot: 僅支持統一帳戶, 創建限價單時, 可以附帶市價止盈止損和限價止盈止損
func (api *OrderCreateAPI) StopLoss(stopLoss string) *OrderCreateAPI {
	api.req.StopLoss = GetPointer(stopLoss)
	return api
}

// tpTriggerBy string false 觸發止盈的價格類型. MarkPrice, IndexPrice, 默認:LastPrice 僅對linear和inverse有效
func (api *OrderCreateAPI) TpTriggerBy(tpTriggerBy string) *OrderCreateAPI {
	api.req.TpTriggerBy = GetPointer(tpTriggerBy)
	return api
}

// slTriggerBy string false 觸發止損的價格類型. MarkPrice, IndexPrice, 默認:LastPrice 僅對linear和inverse有效
func (api *OrderCreateAPI) SlTriggerBy(slTriggerBy string) *OrderCreateAPI {
	api.req.SlTriggerBy = GetPointer(slTriggerBy)
	return api
}

// reduceOnly boolean false 什麼是只減倉? true 將這筆訂單設為只減倉 當減倉時, reduceOnly=true必傳 只減倉單的止盈止損不生效 對linear, inverse和option有效
func (api *OrderCreateAPI) ReduceOnly(reduceOnly bool) *OrderCreateAPI {
	api.req.ReduceOnly = GetPointer(reduceOnly)
	return api
}

// closeOnTrigger boolean false 什麼是觸發後平倉委託?此選項可以確保您的止損單被用於減倉（平倉）而非加倉，並且在可用保證金不足的情況下，取消其他委託，騰出保證金以確保平倉委託的執行. 僅對linear和inverse有效
func (api *OrderCreateAPI) CloseOnTrigger(closeOnTrigger bool) *OrderCreateAPI {
	api.req.CloseOnTrigger = GetPointer(closeOnTrigger)
	return api
}

// smpType string false Smp執行類型. 什麼是SMP?
func (api *OrderCreateAPI) SmpType(smpType string) *OrderCreateAPI {
	api.req.SmpType = GetPointer(smpType)
	return api
}

// mmp boolean false 做市商保護. 僅option有效. true 表示該訂單是做市商保護訂單. 什麼是做市商保護?
func (api *OrderCreateAPI) Mmp(mmp bool) *OrderCreateAPI {
	api.req.Mmp = GetPointer(mmp)
	return api
}

// tpslMode string false 止盈止損模式 Full: 全部倉位止盈止損. 此時, tpOrderType或者slOrderType必須傳Market Partial: 部分倉位止盈止損(下單時沒有size選項, 實際上創建tpsl訂單時, 是按照實際成交的數量來生成止盈止損). 支持創建限價止盈止損. 注意: 創建限價止盈止損時, tpslMode必傳且為Partial 僅對linear和inverse有效
func (api *OrderCreateAPI) TpslMode(tpslMode string) *OrderCreateAPI {
	api.req.TpslMode = GetPointer(tpslMode)
	return api
}

// tpLimitPrice string false 觸發止盈後轉換為限價單的價格 linear & inverse: 僅作用於當tpslMode=Partial以及tpOrderType=Limit時 spot(UTA): 參數必傳當創建訂單時帶了takeProfit 和 tpOrderType=Limit
func (api *OrderCreateAPI) TpLimitPrice(tpLimitPrice string) *OrderCreateAPI {
	api.req.TpLimitPrice = GetPointer(tpLimitPrice)
	return api
}

// slLimitPrice string false 觸發止損後轉換為限價單的價格 linear & inverse: 僅作用於當tpslMode=Partial以及slOrderType=Limit時 spot(UTA): 參數必傳當創建訂單時帶了stopLoss 和 slOrderType=Limit
func (api *OrderCreateAPI) SlLimitPrice(slLimitPrice string) *OrderCreateAPI {
	api.req.SlLimitPrice = GetPointer(slLimitPrice)
	return api
}

// tpOrderType string false 止盈觸發後的訂單類型 linear & inverse: Market(默認), Limit. 對於tpslMode=Full, 僅支持tpOrderType=Market spot(UTA): Market: 當帶了參數"takeProfit", Limit: 當帶了參數"takeProfit" 和 "tpLimitPrice"
func (api *OrderCreateAPI) TpOrderType(tpOrderType string) *OrderCreateAPI {
	api.req.TpOrderType = GetPointer(tpOrderType)
	return api
}

// slOrderType string false 止損觸發後的訂單類型 linear & inverse: Market(默認), Limit. 對於tpslMode=Full, 僅支持slOrderType=Market spot(UTA): Market: 當帶了參數"stopLoss", Limit: 當帶了參數"stopLoss" 和 "slLimitPrice"
func (api *OrderCreateAPI) SlOrderType(slOrderType string) *OrderCreateAPI {
	api.req.SlOrderType = GetPointer(slOrderType)
	return api
}

type OrderCreateBatchReq struct {
	Category *string          `json:"category"` //string	true	產品類型 統一帳戶: spot, linear, inverse, option 經典帳戶: spot, linear, inverse
	Request  []OrderCreateReq `json:"request"`  //array	true	批量訂單請求
}

// 下單時需確保帳戶內有足夠的資金。一旦下單，根據訂單所需資金，您的帳戶資金將在訂單生命週期內凍結相應額度，被凍結的資金額度取決於訂單屬性。
// 每個請求包含的訂單數最大是: 20筆(期权), 10筆(期貨), 10筆(現貨)，返回的數據列表中分成兩個list，訂單創建的列表和創建結果的信息返回，兩個list的訂單的序列是完全保持一致的。
type OrderCreateBatchAPI struct {
	client *PrivateRestClient
	req    *OrderCreateBatchReq
}

func (api *OrderCreateBatchAPI) AddNewOrderCreateReq(orderCreateApi *OrderCreateAPI) *OrderCreateBatchAPI {
	if api.req == nil {
		api.req = &OrderCreateBatchReq{
			Category: orderCreateApi.req.Category,
			Request:  []OrderCreateReq{},
		}
	}
	api.req.Request = append(api.req.Request, *orderCreateApi.req)
	return api
}

func (api *OrderCreateBatchAPI) SetOrderList(orderCreateApiList []*OrderCreateAPI) *OrderCreateBatchAPI {
	if len(orderCreateApiList) == 0 {
		return api
	}
	if api.req == nil {
		api.req = &OrderCreateBatchReq{
			Category: orderCreateApiList[0].req.Category,
			Request:  []OrderCreateReq{},
		}
	}
	for _, orderCreateApi := range orderCreateApiList {
		if *orderCreateApi.req.Category != *api.req.Category {
			continue
		}
		api.req.Request = append(api.req.Request, *orderCreateApi.req)
	}
	return api
}

type OrderAmendReq struct {
	Category     *string `json:"category"`     //string	true	產品類型 統一帳戶: linear, inverse, spot, option 經典帳戶: linear, inverse, spot
	Symbol       *string `json:"symbol"`       //string	true	合約名稱
	OrderId      *string `json:"orderId"`      //string	false	訂單Id. orderId和orderLinkId必傳其中一個
	OrderLinkId  *string `json:"orderLinkId"`  //string	false	用戶自定義訂單Id. orderId和orderLinkId必傳其中一個
	OrderIv      *string `json:"orderIv"`      //string	false	隱含波動率. 僅option有效. 按照實際值傳入, e.g., 對於10%, 則傳入0.1
	TriggerPrice *string `json:"triggerPrice"` //string	false	對於期貨, 是條件單觸發價格參數. 若您希望市場價是要上升後觸發, 確保: triggerPrice > 市場價格 否則, triggerPrice < 市場價格 對於現貨, 這是下止盈止損單或者條件單的觸發價格參數
	Qty          *string `json:"qty"`          //string	false	修改後的訂單數量. 若不修改，請不要傳該字段
	Price        *string `json:"price"`        //string	false	修改後的訂單價格. 若不修改，請不要傳該字段
	TpslMode     *string `json:"tpslMode"`     //string	false	止盈止損模式 Full: 全部倉位止盈止損. 此時, tpOrderType或者slOrderType必須傳Market Partial: 部分倉位止盈止損. 支持創建限價止盈止損. 注意: 創建限價止盈止損時, tpslMode必傳且為Partial 僅對linear和inverse有效
	TakeProfit   *string `json:"takeProfit"`   //string	false	修改後的止盈價格. 當傳"0"時, 表示取消當前訂單上設置的止盈. 若不修改，請不要傳該字段 適用於 spot(UTA), linear, inverse
	StopLoss     *string `json:"stopLoss"`     //string	false	修改後的止損價格. 當傳"0"時, 表示取消當前訂單上設置的止損. 若不修改，請不要傳該字段 適用於 spot(UTA), linear, inverse
	TpTriggerBy  *string `json:"tpTriggerBy"`  //string	false	止盈價格觸發類型. 若下單時未設置該值，則調用該接口修改止盈價格時，該字段必傳
	SlTriggerBy  *string `json:"slTriggerBy"`  //string	false	止損價格觸發類型. 若下單時未設置該值，則調用該接口修改止損價格時，該字段必傳
	TriggerBy    *string `json:"triggerBy"`    //string	false	觸發價格的觸發類型
	TpLimitPrice *string `json:"tpLimitPrice"` //string	false	*觸發止盈後轉換為限價單的價格. 當且僅當原始訂單下單時創建的是部分止盈止損限價單, 本字段才有效 適用於 spot(UTA), linear, inverse
	SlLimitPrice *string `json:"slLimitPrice"` //string	false	*觸發止損後轉換為限價單的價格. 當且僅當原始訂單下單時創建的是部分止盈止損限價單, 本字段才有效 適用於 spot(UTA), linear, inverse
}

type OrderAmendAPI struct {
	client *PrivateRestClient
	req    *OrderAmendReq
}

// category string true 產品類型 統一帳戶: linear, inverse, spot, option 經典帳戶: linear, inverse, spot
func (api *OrderAmendAPI) Category(category string) *OrderAmendAPI {
	api.req.Category = GetPointer(category)
	return api
}

// symbol string true 合約名稱
func (api *OrderAmendAPI) Symbol(symbol string) *OrderAmendAPI {
	api.req.Symbol = GetPointer(symbol)
	return api
}

// orderId string false 訂單Id. orderId和orderLinkId必傳其中一個
func (api *OrderAmendAPI) OrderId(orderId string) *OrderAmendAPI {
	api.req.OrderId = GetPointer(orderId)
	return api
}

// orderLinkId string false 用戶自定義訂單Id. orderId和orderLinkId必傳其中一個
func (api *OrderAmendAPI) OrderLinkId(orderLinkId string) *OrderAmendAPI {
	api.req.OrderLinkId = GetPointer(orderLinkId)
	return api
}

// orderIv string false 隱含波動率. 僅option有效. 按照實際值傳入, e.g., 對於10%, 則傳入0.1
func (api *OrderAmendAPI) OrderIv(orderIv string) *OrderAmendAPI {
	api.req.OrderIv = GetPointer(orderIv)
	return api
}

// triggerPrice string false 對於期貨, 是條件單觸發價格參數. 若您希望市場價是要上升後觸發, 確保: triggerPrice > 市場價格 否則, triggerPrice < 市場價格 對於現貨, 這是下止盈止損單或者條件單的觸發價格參數
func (api *OrderAmendAPI) TriggerPrice(triggerPrice string) *OrderAmendAPI {
	api.req.TriggerPrice = GetPointer(triggerPrice)
	return api
}

// qty string false 修改後的訂單數量. 若不修改，請不要傳該字段
func (api *OrderAmendAPI) Qty(qty string) *OrderAmendAPI {
	api.req.Qty = GetPointer(qty)
	return api
}

// price string false 修改後的訂單價格. 若不修改，請不要傳該字段
func (api *OrderAmendAPI) Price(price string) *OrderAmendAPI {
	api.req.Price = GetPointer(price)
	return api
}

// tpslMode string false 止盈止損模式 Full: 全部倉位止盈止損. 此時, tpOrderType或者slOrderType必須傳Market Partial: 部分倉位止盈止損. 支持創建限價止盈止損. 注意: 創建限價止盈止損時, tpslMode必傳且為Partial 僅對linear和inverse有效
func (api *OrderAmendAPI) TpslMode(tpslMode string) *OrderAmendAPI {
	api.req.TpslMode = GetPointer(tpslMode)
	return api
}

// takeProfit string false 修改後的止盈價格. 當傳"0"時, 表示取消當前訂單上設置的止盈. 若不修改，請不要傳該字段 適用於 spot(UTA), linear, inverse
func (api *OrderAmendAPI) TakeProfit(takeProfit string) *OrderAmendAPI {
	api.req.TakeProfit = GetPointer(takeProfit)
	return api
}

// stopLoss string false 修改後的止損價格. 當傳"0"時, 表示取消當前訂單上設置的止損. 若不修改，請不要傳該字段 適用於 spot(UTA), linear, inverse
func (api *OrderAmendAPI) StopLoss(stopLoss string) *OrderAmendAPI {
	api.req.StopLoss = GetPointer(stopLoss)
	return api
}

// tpTriggerBy string false 止盈價格觸發類型. 若下單時未設置該值，則調用該接口修改止盈價格時，該字段必傳
func (api *OrderAmendAPI) TpTriggerBy(tpTriggerBy string) *OrderAmendAPI {
	api.req.TpTriggerBy = GetPointer(tpTriggerBy)
	return api
}

// slTriggerBy string false 止損價格觸發類型. 若下單時未設置該值，則調用該接口修改止損價格時，該字段必傳
func (api *OrderAmendAPI) SlTriggerBy(slTriggerBy string) *OrderAmendAPI {
	api.req.SlTriggerBy = GetPointer(slTriggerBy)
	return api
}

// triggerBy string false 觸發價格的觸發類型
func (api *OrderAmendAPI) TriggerBy(triggerBy string) *OrderAmendAPI {
	api.req.TriggerBy = GetPointer(triggerBy)
	return api
}

// tpLimitPrice string false *觸發止盈後轉換為限價單的價格. 當且僅當原始訂單下單時創建的是部分止盈止損限價單, 本字段才有效 適用於 spot(UTA), linear, inverse
func (api *OrderAmendAPI) TpLimitPrice(tpLimitPrice string) *OrderAmendAPI {
	api.req.TpLimitPrice = GetPointer(tpLimitPrice)
	return api
}

// slLimitPrice string false *觸發止損後轉換為限價單的價格. 當且僅當原始訂單下單時創建的是部分止盈止損限價單, 本字段才有效 適用於 spot(UTA), linear, inverse
func (api *OrderAmendAPI) SlLimitPrice(slLimitPrice string) *OrderAmendAPI {
	api.req.SlLimitPrice = GetPointer(slLimitPrice)
	return api
}

type OrderAmendBatchReq struct {
	Category *string         `json:"category"` //string	true	產品類型 統一帳戶: linear, inverse, spot, option 經典帳戶: linear, inverse, spot
	Request  []OrderAmendReq `json:"request"`  //array	true	批量訂單請求
}

type OrderAmendBatchAPI struct {
	client *PrivateRestClient
	req    *OrderAmendBatchReq
}

func (api *OrderAmendBatchAPI) AddNewOrderAmendReq(orderAmendApi *OrderAmendAPI) *OrderAmendBatchAPI {
	if api.req == nil {
		api.req = &OrderAmendBatchReq{
			Category: orderAmendApi.req.Category,
			Request:  []OrderAmendReq{},
		}
	}
	api.req.Request = append(api.req.Request, *orderAmendApi.req)
	return api
}

func (api *OrderAmendBatchAPI) SetOrderList(orderAmendApiList []*OrderAmendAPI) *OrderAmendBatchAPI {
	if len(orderAmendApiList) == 0 {
		return api
	}
	if api.req == nil {
		api.req = &OrderAmendBatchReq{
			Category: orderAmendApiList[0].req.Category,
			Request:  []OrderAmendReq{},
		}
	}
	for _, orderAmendApi := range orderAmendApiList {
		if *orderAmendApi.req.Category != *api.req.Category {
			continue
		}
		api.req.Request = append(api.req.Request, *orderAmendApi.req)
	}
	return api
}

type OrderCancelReq struct {
	Category    *string `json:"category"`    //string	true	產品類型 統一帳戶: spot, linear, option 經典帳戶: spot, linear, inverse
	Symbol      *string `json:"symbol"`      //string	true	合約名稱
	OrderId     *string `json:"orderId"`     //string	false	訂單Id. orderId和orderLinkId必傳其中一個
	OrderLinkId *string `json:"orderLinkId"` //string	false	用戶自定義訂單Id. orderId和orderLinkId必傳其中一個
	OrderFilter *string `json:"orderFilter"` //string	false	僅spot有效. Order: 普通單,tpslOrder: 止盈止損單,StopOrder: 條件單. 若不傳, 默認是Order
}

type OrderCancelAPI struct {
	client *PrivateRestClient
	req    *OrderCancelReq
}

// category string true 產品類型 統一帳戶: spot, linear, option 經典帳戶: spot, linear, inverse
func (api *OrderCancelAPI) Category(category string) *OrderCancelAPI {
	api.req.Category = GetPointer(category)
	return api
}

// symbol string true 合約名稱
func (api *OrderCancelAPI) Symbol(symbol string) *OrderCancelAPI {
	api.req.Symbol = GetPointer(symbol)
	return api
}

// orderId string false 訂單Id. orderId和orderLinkId必傳其中一個
func (api *OrderCancelAPI) OrderId(orderId string) *OrderCancelAPI {
	api.req.OrderId = GetPointer(orderId)
	return api
}

// orderLinkId string false 用戶自定義訂單Id. orderId和orderLinkId必傳其中一個
func (api *OrderCancelAPI) OrderLinkId(orderLinkId string) *OrderCancelAPI {
	api.req.OrderLinkId = GetPointer(orderLinkId)
	return api
}

// orderFilter string false 僅spot有效. Order: 普通單,tpslOrder: 止盈止損單,StopOrder: 條件單. 若不傳, 默認是Order
func (api *OrderCancelAPI) OrderFilter(orderFilter string) *OrderCancelAPI {
	api.req.OrderFilter = GetPointer(orderFilter)
	return api
}

type OrderCancelBatchReq struct {
	Category *string          `json:"category"` //string	true	產品類型 統一帳戶: spot, linear, option 經典帳戶: spot, linear, inverse
	Request  []OrderCancelReq `json:"request"`  //array	true	批量訂單請求
}

type OrderCancelBatchAPI struct {
	client *PrivateRestClient
	req    *OrderCancelBatchReq
}

func (api *OrderCancelBatchAPI) AddNewOrderCancelReq(orderCancelApi *OrderCancelAPI) *OrderCancelBatchAPI {
	if api.req == nil {
		api.req = &OrderCancelBatchReq{
			Category: orderCancelApi.req.Category,
			Request:  []OrderCancelReq{},
		}
	}
	api.req.Request = append(api.req.Request, *orderCancelApi.req)
	return api
}

func (api *OrderCancelBatchAPI) SetOrderList(orderCancelApiList []*OrderCancelAPI) *OrderCancelBatchAPI {
	if len(orderCancelApiList) == 0 {
		return api
	}
	if api.req == nil {
		api.req = &OrderCancelBatchReq{
			Category: orderCancelApiList[0].req.Category,
			Request:  []OrderCancelReq{},
		}
	}
	for _, orderCancelApi := range orderCancelApiList {
		if *orderCancelApi.req.Category != *api.req.Category {
			continue
		}
		api.req.Request = append(api.req.Request, *orderCancelApi.req)
	}
	return api
}

type OrderCancelAllReq struct {
	Category      *string `json:"category"`      //string	true	產品類型 統一帳戶: spot, linear, inverse, option 經典帳戶: spot, linear, inverse
	Symbol        *string `json:"symbol"`        //string	false	合約名稱. 對於linear & inverse: 若不傳baseCoin和settleCoin, 該字段必傳
	BaseCoin      *string `json:"baseCoin"`      //string	false	交易幣種 對於經典帳戶下的linear & inverse: 當通過baseCoin來全部撤單時，會將linear和inverse訂單全部撤掉。若不傳symbol和baseCoin, 則該字段必傳 對於統一帳戶下的linear & inverse: 當通過baseCoin來全部撤單時，會將對應category的訂單全部撤掉。若不傳symbol和baseCoin, 則該字段必傳 對於經典帳戶的現貨: 該字段無效
	SettleCoin    *string `json:"settleCoin"`    //string	false	結算幣種 對於linear & inverse: 該字段必傳, 若不傳symbol和baseCoin 該字段不支持spot
	OrderFilter   *string `json:"orderFilter"`   //string	false	當category=spot, 該字段可以傳Order(普通單), tpslOrder(止盈止損單), StopOrder(條件單), OcoOrder, BidirectionalTpslOrder(現貨雙向止盈止損訂單). 若不傳, 則默認是撤掉Order單 當category=linear 或者 inverse, 該字段可以傳Order(普通單), StopOrder(條件單, 包括止盈止損單和追蹤出場單). 若不傳, 則所有類型的訂單都會被撤掉 當category=option, 該字段可以傳Order, 不管傳與不傳, 都是撤掉所有訂單
	StopOrderType *string `json:"stopOrderType"` //string	false	條件單類型, Stop 僅用於當category=linear 或者 inverse以及orderFilter=StopOrder時, 若想僅取消條件單 (不包括止盈止損單和追蹤出場單), 則可以傳入該字段
}

type OrderCancelAllAPI struct {
	client *PrivateRestClient
	req    *OrderCancelAllReq
}

// category string true 產品類型 統一帳戶: spot, linear, inverse, option 經典帳戶: spot, linear, inverse
func (api *OrderCancelAllAPI) Category(category string) *OrderCancelAllAPI {
	api.req.Category = GetPointer(category)
	return api
}

// symbol string false 合約名稱. 對於linear & inverse: 若不傳baseCoin和settleCoin, 該字段必傳
func (api *OrderCancelAllAPI) Symbol(symbol string) *OrderCancelAllAPI {
	api.req.Symbol = GetPointer(symbol)
	return api
}

// baseCoin string false 交易幣種 對於經典帳戶下的linear & inverse: 當通過baseCoin來全部撤單時，會將linear和inverse訂單全部撤掉。若不傳symbol和baseCoin, 則該字段必傳 對於統一帳戶下的linear & inverse: 當通過baseCoin來全部撤單時，會將對應category的訂單全部撤掉。若不傳symbol和baseCoin, 則該字段必傳 對於經典帳戶的現貨: 該字段無效
func (api *OrderCancelAllAPI) BaseCoin(baseCoin string) *OrderCancelAllAPI {
	api.req.BaseCoin = GetPointer(baseCoin)
	return api
}

// settleCoin string false 結算幣種 對於linear & inverse: 該字段必傳, 若不傳symbol和baseCoin 該字段不支持spot
func (api *OrderCancelAllAPI) SettleCoin(settleCoin string) *OrderCancelAllAPI {
	api.req.SettleCoin = GetPointer(settleCoin)
	return api
}

// orderFilter string false 當category=spot, 該字段可以傳Order(普通單), tpslOrder(止盈止損單), StopOrder(條件單), OcoOrder, BidirectionalTpslOrder(現貨雙向止盈止損訂單). 若不傳, 則默認是撤掉Order單 當category=linear 或者 inverse, 該字段可以傳Order(普通單), StopOrder(條件單, 包括止盈止損單和追蹤出場單). 若不傳, 則所有類型的訂單都會被撤掉 當category=option, 該字段可以傳Order, 不管傳與不傳, 都是撤掉所有訂單\
func (api *OrderCancelAllAPI) OrderFilter(orderFilter string) *OrderCancelAllAPI {
	api.req.OrderFilter = GetPointer(orderFilter)
	return api
}

// stopOrderType string false 條件單類型, Stop 僅用於當category=linear 或者 inverse以及orderFilter=StopOrder時, 若想僅取消條件單 (不包括止盈止損單和追蹤出場單), 則可以傳入該字段
func (api *OrderCancelAllAPI) StopOrderType(stopOrderType string) *OrderCancelAllAPI {
	api.req.StopOrderType = GetPointer(stopOrderType)
	return api
}

type OrderRealtimeReq struct {
	Category    *string `json:"category"`    //string	true	產品類型 統一帳戶: spot, linear, inverse, option 經典帳戶: spot, linear, inverse
	Symbol      *string `json:"symbol"`      //string	false	合約名稱. 對於linear, symbol, baseCoin 和 settleCoin必傳其中一個
	BaseCoin    *string `json:"baseCoin"`    //string	false	交易幣種. 支持linear, inverse和option. 對於若不傳，則返回期權下所有活動委託單
	SettleCoin  *string `json:"settleCoin"`  //string	false	結算幣種 linear: symbol 和 settleCoin必傳其中一個 spot: 該字段無效
	OrderId     *string `json:"orderId"`     //string	false	訂單Id
	OrderLinkId *string `json:"orderLinkId"` //string	false	用戶自定義訂單Id
	OpenOnly    *int    `json:"openOnly"`    //integer	false	統一帳戶 & 經典帳戶0(默認): 僅查詢活動委託訂單 統一帳戶: 1, 統一帳戶(反向)和經典帳戶: 2
	OrderFilter *string `json:"orderFilter"` //string	false	Order: 活動單, StopOrder: 條件單, 支持現貨和期貨, tpslOrder: 止盈止損單, 僅現貨有效, OcoOrder: OCO訂單, BidirectionalTpslOrder: 現貨(UTA)雙向止盈止損訂單
	Limit       *int    `json:"limit"`       //integer	false	每頁數量限制. [1, 50]. 默認: 20
	Cursor      *string `json:"cursor"`      //string	false	游標，用於翻頁
}

type OrderRealtimeAPI struct {
	client *PrivateRestClient
	req    *OrderRealtimeReq
}

// category string true 產品類型 統一帳戶: spot, linear, inverse, option 經典帳戶: spot, linear, inverse
func (api *OrderRealtimeAPI) Category(category string) *OrderRealtimeAPI {
	api.req.Category = GetPointer(category)
	return api
}

// symbol string false 合約名稱. 對於linear, symbol, baseCoin 和 settleCoin必傳其中一個
func (api *OrderRealtimeAPI) Symbol(symbol string) *OrderRealtimeAPI {
	api.req.Symbol = GetPointer(symbol)
	return api
}

// baseCoin string false 交易幣種. 支持linear, inverse和option. 對於若不傳，則返回期權下所有活動委託單
func (api *OrderRealtimeAPI) BaseCoin(baseCoin string) *OrderRealtimeAPI {
	api.req.BaseCoin = GetPointer(baseCoin)
	return api
}

// settleCoin string false 結算幣種 linear: symbol 和 settleCoin必傳其中一個 spot: 該字段無效
func (api *OrderRealtimeAPI) SettleCoin(settleCoin string) *OrderRealtimeAPI {
	api.req.SettleCoin = GetPointer(settleCoin)
	return api
}

// orderId string false 訂單Id
func (api *OrderRealtimeAPI) OrderId(orderId string) *OrderRealtimeAPI {
	api.req.OrderId = GetPointer(orderId)
	return api
}

// orderLinkId string false 用戶自定義訂單Id
func (api *OrderRealtimeAPI) OrderLinkId(orderLinkId string) *OrderRealtimeAPI {
	api.req.OrderLinkId = GetPointer(orderLinkId)
	return api
}

// openOnly int false 統一帳戶 & 經典帳戶0(默認): 僅查詢活動委託訂單 統一帳戶: 1, 統一帳戶(反向)和經典帳戶: 2
func (api *OrderRealtimeAPI) OpenOnly(openOnly int) *OrderRealtimeAPI {
	api.req.OpenOnly = GetPointer(openOnly)
	return api
}

// orderFilter string false Order: 活動單, StopOrder: 條件單, 支持現貨和期貨, tpslOrder: 止盈止損單, 僅現貨有效, OcoOrder: OCO訂單, BidirectionalTpslOrder: 現貨(UTA)雙向止盈止損訂單
func (api *OrderRealtimeAPI) OrderFilter(orderFilter string) *OrderRealtimeAPI {
	api.req.OrderFilter = GetPointer(orderFilter)
	return api
}

// limit int false 每頁數量限制. [1, 50]. 默認: 20
func (api *OrderRealtimeAPI) Limit(limit int) *OrderRealtimeAPI {
	api.req.Limit = GetPointer(limit)
	return api
}

// cursor string false 游標，用於翻頁
func (api *OrderRealtimeAPI) Cursor(cursor string) *OrderRealtimeAPI {
	api.req.Cursor = GetPointer(cursor)
	return api
}

type OrderHistoryReq struct {
	Category    *string `json:"category"`    //string	true	產品類型 統一帳戶: spot, linear, inverse, option 經典帳戶: spot, linear, inverse
	Symbol      *string `json:"symbol"`      //string	false	合約名稱
	BaseCoin    *string `json:"baseCoin"`    //string	false	交易幣種. 統一帳戶（反向）以及經典帳戶不支持該字段的查詢
	SettleCoin  *string `json:"settleCoin"`  //string	false	結算幣種. 統一帳戶（反向）以及經典帳戶不支持該字段的查詢
	OrderId     *string `json:"orderId"`     //string	false	訂單ID
	OrderLinkId *string `json:"orderLinkId"` //string	false	用戶自定義訂單ID
	OrderFilter *string `json:"orderFilter"` //string	false	Order: 普通單, StopOrder: 條件單, 支持現貨和期貨, tpslOrder: 現貨止盈止損單, OcoOrder: OCO訂單, BidirectionalTpslOrder: 現貨(UTA)雙向止盈止損訂單
	OrderStatus *string `json:"orderStatus"` //string	false	訂單狀態 經典帳戶spot: 該字段無效 UTA(spot, linear, option): 不傳則默認查詢所有終態訂單 UTA(inverse)和經典帳戶: 不傳則默認查詢活動態+終態的訂單
	StartTime   *int64  `json:"startTime"`   //integer	false	開始時間戳 (毫秒). 經典帳戶現貨不支持使用startTime和endTime startTime 和 endTime都不傳入, 則默認返回最近7天的數據 startTime 和 endTime都傳入的話, 則確保endTime - startTime <= 7天 若只傳startTime，則查詢startTime和startTime+7天的數據 若只傳endTime，則查詢endTime-7天和endTime的數據
	EndTime     *int64  `json:"endTime"`     //integer	false	結束時間戳 (毫秒)
	Limit       *int    `json:"limit"`       //integer	false	每頁數量限制. [1, 50]. 默認: 20
	Cursor      *string `json:"cursor"`      //string	false	游標，用於翻頁
}

type OrderHistoryAPI struct {
	client *PrivateRestClient
	req    *OrderHistoryReq
}

// category string true 產品類型 統一帳戶: spot, linear, inverse, option 經典帳戶: spot, linear, inverse
func (api *OrderHistoryAPI) Category(category string) *OrderHistoryAPI {
	api.req.Category = GetPointer(category)
	return api
}

// symbol string false 合約名稱
func (api *OrderHistoryAPI) Symbol(symbol string) *OrderHistoryAPI {
	api.req.Symbol = GetPointer(symbol)
	return api
}

// baseCoin string false 交易幣種. 統一帳戶（反向）以及經典帳戶不支持該字段的查詢
func (api *OrderHistoryAPI) BaseCoin(baseCoin string) *OrderHistoryAPI {
	api.req.BaseCoin = GetPointer(baseCoin)
	return api
}

// settleCoin string false 結算幣種. 統一帳戶（反向）以及經典帳戶不支持該字段的查詢
func (api *OrderHistoryAPI) SettleCoin(settleCoin string) *OrderHistoryAPI {
	api.req.SettleCoin = GetPointer(settleCoin)
	return api
}

// orderId string false 訂單ID
func (api *OrderHistoryAPI) OrderId(orderId string) *OrderHistoryAPI {
	api.req.OrderId = GetPointer(orderId)
	return api
}

// orderLinkId string false 用戶自定義訂單ID
func (api *OrderHistoryAPI) OrderLinkId(orderLinkId string) *OrderHistoryAPI {
	api.req.OrderLinkId = GetPointer(orderLinkId)
	return api
}

// orderFilter string false Order: 普通單, StopOrder: 條件單, 支持現貨和期貨, tpslOrder: 現貨止盈止損單, OcoOrder: OCO訂單, BidirectionalTpslOrder: 現貨(UTA)雙向止盈止損訂單
func (api *OrderHistoryAPI) OrderFilter(orderFilter string) *OrderHistoryAPI {
	api.req.OrderFilter = GetPointer(orderFilter)
	return api
}

// orderStatus string false 訂單狀態 經典帳戶spot: 該字段無效 UTA(spot, linear, option): 不傳則默認查詢所有終態訂單 UTA(inverse)和經典帳戶: 不傳則默認查詢活動態+終態的訂單
func (api *OrderHistoryAPI) OrderStatus(orderStatus string) *OrderHistoryAPI {
	api.req.OrderStatus = GetPointer(orderStatus)
	return api
}

// startTime int false 開始時間戳 (毫秒). 經典帳戶現貨不支持使用startTime和endTime startTime 和 endTime都不傳入, 則默認返回最近7天的數據 startTime 和 endTime都傳入的話, 則確保endTime - startTime <= 7天 若只傳startTime，則查詢startTime和startTime+7天的數據 若只傳endTime，則查詢endTime-7天和endTime的數據
func (api *OrderHistoryAPI) StartTime(startTime int64) *OrderHistoryAPI {
	api.req.StartTime = GetPointer(startTime)
	return api
}

// endTime int false 結束時間戳 (毫秒)
func (api *OrderHistoryAPI) EndTime(endTime int64) *OrderHistoryAPI {
	api.req.EndTime = GetPointer(endTime)
	return api
}

// limit int false 每頁數量限制. [1, 50]. 默認: 20
func (api *OrderHistoryAPI) Limit(limit int) *OrderHistoryAPI {
	api.req.Limit = GetPointer(limit)
	return api
}

// cursor string false 游標，用於翻頁
func (api *OrderHistoryAPI) Cursor(cursor string) *OrderHistoryAPI {
	api.req.Cursor = GetPointer(cursor)
	return api
}

type OrderExecutionListReq struct {
	Category    *string `json:"category"`    //string	true	產品類型 統一帳戶: spot, linear, inverse, option 經典帳戶: spot, linear, inverse
	Symbol      *string `json:"symbol"`      //string	false	合約名稱
	OrderId     *string `json:"orderId"`     //string	false	訂單Id
	OrderLinkId *string `json:"orderLinkId"` //string	false	用戶自定義訂單id. 經典帳戶不支持該字段查詢
	BaseCoin    *string `json:"baseCoin"`    //string	false	交易幣種. 統一帳戶(反向)和經典帳戶不支持該字段查詢
	StartTime   *int64  `json:"startTime"`   //integer	false	開始時間戳 (毫秒)
	EndTime     *int64  `json:"endTime"`     //integer	false	結束時間戳 (毫秒)
	ExecType    *string `json:"execType"`    //string	false	執行類型. 經典帳戶現貨交易無效
	Limit       *int    `json:"limit"`       //integer	false	每頁數量限制. [1, 100]. 默認: 50
	Cursor      *string `json:"cursor"`      //string	false	游標，用於翻頁
}

type OrderExecutionListAPI struct {
	client *PrivateRestClient
	req    *OrderExecutionListReq
}

// category string true 產品類型 統一帳戶: spot, linear, inverse, option 經典帳戶: spot, linear, inverse
func (api *OrderExecutionListAPI) Category(category string) *OrderExecutionListAPI {
	api.req.Category = GetPointer(category)
	return api
}

// symbol string false 合約名稱
func (api *OrderExecutionListAPI) Symbol(symbol string) *OrderExecutionListAPI {
	api.req.Symbol = GetPointer(symbol)
	return api
}

// orderId string false 訂單Id
func (api *OrderExecutionListAPI) OrderId(orderId string) *OrderExecutionListAPI {
	api.req.OrderId = GetPointer(orderId)
	return api
}

// orderLinkId string false 用戶自定義訂單id. 經典帳戶不支持該字段查詢
func (api *OrderExecutionListAPI) OrderLinkId(orderLinkId string) *OrderExecutionListAPI {
	api.req.OrderLinkId = GetPointer(orderLinkId)
	return api
}

// baseCoin string false 交易幣種. 統一帳戶(反向)和經典帳戶不支持該字段查詢
func (api *OrderExecutionListAPI) BaseCoin(baseCoin string) *OrderExecutionListAPI {
	api.req.BaseCoin = GetPointer(baseCoin)
	return api
}

// startTime int false 開始時間戳 (毫秒)
func (api *OrderExecutionListAPI) StartTime(startTime int64) *OrderExecutionListAPI {
	api.req.StartTime = GetPointer(startTime)
	return api
}

// endTime int false 結束時間戳 (毫秒)
func (api *OrderExecutionListAPI) EndTime(endTime int64) *OrderExecutionListAPI {
	api.req.EndTime = GetPointer(endTime)
	return api
}

// execType string false 執行類型. 經典帳戶現貨交易無效
func (api *OrderExecutionListAPI) ExecType(execType string) *OrderExecutionListAPI {
	api.req.ExecType = GetPointer(execType)
	return api
}

// limit int false 每頁數量限制. [1, 100]. 默認: 50
func (api *OrderExecutionListAPI) Limit(limit int) *OrderExecutionListAPI {
	api.req.Limit = GetPointer(limit)
	return api
}

// cursor string false 游標，用於翻頁
func (api *OrderExecutionListAPI) Cursor(cursor string) *OrderExecutionListAPI {
	api.req.Cursor = GetPointer(cursor)
	return api
}
