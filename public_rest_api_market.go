package mybybitapi

// bybit MarketInstrumentsInfo PublicRest接口 GET 查詢可交易產品的規格信息
func (client *PublicRestClient) NewMarketInstrumentsInfo() *MarketInstrumentsInfoAPI {
	return &MarketInstrumentsInfoAPI{
		client: client,
		req:    &MarketInstrumentsInfoReq{},
	}
}
func (api *MarketInstrumentsInfoAPI) Do() (*BybitRestRes[MarketInstrumentsInfoRes], error) {
	url := bybitHandlerRequestAPIWithPathQueryParam(REST, api.req, PublicRestAPIMap[MarketInstrumentsInfo])
	return bybitCallAPI[MarketInstrumentsInfoRes](api.client.c, url, NIL_REQBODY, GET)
}

// bybit MarketTime PublicRest接口 GET Bybit服務器時間
func (client *PublicRestClient) NewMarketTime() *MarketTimeAPI {
	return &MarketTimeAPI{
		client: client,
		req:    &MarketTimeReq{},
	}
}
func (api *MarketTimeAPI) Do() (*BybitRestRes[MarketTimeRes], error) {
	url := bybitHandlerRequestAPIWithPathQueryParam(REST, api.req, PublicRestAPIMap[MarketTime])
	return bybitCallAPI[MarketTimeRes](api.client.c, url, NIL_REQBODY, GET)
}

// bybit MarketKline PublicRest接口 GET 查詢市場價格K線數據
func (client *PublicRestClient) NewMarketKline() *MarketKlineAPI {
	return &MarketKlineAPI{
		client: client,
		req:    &MarketKlineReq{},
	}
}
func (api *MarketKlineAPI) Do() (*BybitRestRes[MarketKlineRes], error) {
	url := bybitHandlerRequestAPIWithPathQueryParam(REST, api.req, PublicRestAPIMap[MarketKline])
	res, err := bybitCallAPI[MarketKlineMiddle](api.client.c, url, NIL_REQBODY, GET)
	if err != nil {
		return nil, err
	}
	return resConvertTo(res, res.Result.ConvertToRes()), nil
}

// bybit MarketOrderBook PublicRest接口 GET Order Book (深度)
func (client *PublicRestClient) NewMarketOrderBook() *MarketOrderBookAPI {
	return &MarketOrderBookAPI{
		client: client,
		req:    &MarketOrderBookReq{},
	}
}
func (api *MarketOrderBookAPI) Do() (*BybitRestRes[MarketOrderBookRes], error) {
	url := bybitHandlerRequestAPIWithPathQueryParam(REST, api.req, PublicRestAPIMap[MarketOrderBook])
	res, err := bybitCallAPI[MarketOrderBookMiddle](api.client.c, url, NIL_REQBODY, GET)
	if err != nil {
		return nil, err
	}
	return resConvertTo(res, res.Result.ConvertToRes()), nil
}

// bybit MarketTickers PublicRest接口 GET 查詢最新行情信息
func (client *PublicRestClient) NewMarketTickers() *MarketTickersAPI {
	return &MarketTickersAPI{
		client: client,
		req:    &MarketTickersReq{},
	}
}
func (api *MarketTickersAPI) Do() (*BybitRestRes[MarketTickersRes], error) {
	url := bybitHandlerRequestAPIWithPathQueryParam(REST, api.req, PublicRestAPIMap[MarketTickers])
	return bybitCallAPI[MarketTickersRes](api.client.c, url, NIL_REQBODY, GET)
}
