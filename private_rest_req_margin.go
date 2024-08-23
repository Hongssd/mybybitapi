package mybybitapi

type SpotMarginTradeSetLeverageReq struct {
	Leverage string `json:"leverage"` // 槓桿倍數 (整數), 支持區間 [2, 10]
}
type SpotMarginTradeSetLeverageAPI struct {
	client *PrivateRestClient
	req    *SpotMarginTradeSetLeverageReq
}

func (api *SpotMarginTradeSetLeverageAPI) Leverage(leverage string) *SpotMarginTradeSetLeverageAPI {
	api.req.Leverage = leverage
	return api
}

type SpotMarginTradeStateReq struct{}
type SpotMarginTradeStateAPI struct {
	client *PrivateRestClient
	req    *SpotMarginTradeStateReq
}
