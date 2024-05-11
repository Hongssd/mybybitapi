package mybybitapi

// category	true	string	产品类型. spot,linear,inverse,option
// symbol	false	string	合约名称
// status	false	string	交易对状态筛选  现货/期货 仅有Trading状态
// baseCoin	false	string	交易货币. 仅对期货/期权有效 对于期权, 默认返回baseCoin为BTC的
// limit	false	integer	每页数量限制. [1, 1000]. 默认: 500
// cursor	false	string	游标，用于翻页
type MarketInstrumentsInfoReq struct {
	Category *string `json:"category"` //String	true	string	产品类型. spot,linear,inverse,option
	Symbol   *string `json:"symbol"`   //String	false	string	合约名称
	Status   *string `json:"status"`   //String	false	string	交易对状态筛选  现货/期货 仅有Trading状态
	BaseCoin *string `json:"baseCoin"` //String	false	string	交易货币. 仅对期货/期权有效 对于期权, 默认返回baseCoin为BTC的
	Limit    *int    `json:"limit"`    //String	false	integer	每页数量限制. [1, 1000]. 默认: 500
	Cursor   *string `json:"cursor"`   //String	false	string	游标，用于翻页
}

type MarketInstrumentsInfoAPI struct {
	client *PublicRestClient
	req    *MarketInstrumentsInfoReq
}

// category	true	string	产品类型. spot,linear,inverse,option
func (api *MarketInstrumentsInfoAPI) Category(category string) *MarketInstrumentsInfoAPI {
	api.req.Category = GetPointer(category)
	return api
}

// symbol	false	string	合约名称
func (api *MarketInstrumentsInfoAPI) Symbol(symbol string) *MarketInstrumentsInfoAPI {
	api.req.Symbol = GetPointer(symbol)
	return api
}

// status	false	string	交易对状态筛选  现货/期货 仅有Trading状态
func (api *MarketInstrumentsInfoAPI) Status(status string) *MarketInstrumentsInfoAPI {
	api.req.Status = GetPointer(status)
	return api
}

// baseCoin	false	string	交易货币. 仅对期货/期权有效 对于期权, 默认返回baseCoin为BTC的
func (api *MarketInstrumentsInfoAPI) BaseCoin(baseCoin string) *MarketInstrumentsInfoAPI {
	api.req.BaseCoin = GetPointer(baseCoin)
	return api
}

// limit	false	integer	每页数量限制. [1, 1000]. 默认: 500
func (api *MarketInstrumentsInfoAPI) Limit(limit int) *MarketInstrumentsInfoAPI {
	api.req.Limit = &limit
	return api
}

// cursor	false	string	游标，用于翻页
func (api *MarketInstrumentsInfoAPI) Cursor(cursor string) *MarketInstrumentsInfoAPI {
	api.req.Cursor = GetPointer(cursor)
	return api
}

type MarketTimeReq struct {
}

type MarketTimeAPI struct {
	client *PublicRestClient
	req    *MarketTimeReq
}

// category	false	string	產品類型. spot,linear,inverse 當category不指定時, 默認是linear
// symbol	true	string	合約名稱
// interval	true	string	時間粒度. 1,3,5,15,30,60,120,240,360,720,D,M,W
// start	false	integer	開始時間戳 (毫秒)
// end	false	integer	結束時間戳 (毫秒)
// limit	false	integer	每頁數量限制. [1, 1000]. 默認: 200
type MarketKlineReq struct {
	Category *string `json:"category"` //String	false	string	產品類型. spot,linear,inverse 當category不指定時, 默認是linear
	Symbol   *string `json:"symbol"`   //String	true	string	合約名稱
	Interval *string `json:"interval"` //String	true	string	時間粒度. 1,3,5,15,30,60,120,240,360,720,D,M,W
	Start    *int64  `json:"start"`    //String	false	integer	開始時間戳 (毫秒)
	End      *int64  `json:"end"`      //String	false	integer	結束時間戳 (毫秒)
	Limit    *int    `json:"limit"`    //String	false	integer	每頁數量限制. [1, 1000]. 默認: 200
}

type MarketKlineAPI struct {
	client *PublicRestClient
	req    *MarketKlineReq
}

// category	false	string	產品類型. spot,linear,inverse 當category不指定時, 默認是linear
func (api *MarketKlineAPI) Category(category string) *MarketKlineAPI {
	api.req.Category = GetPointer(category)
	return api
}

// symbol	true	string	合約名稱
func (api *MarketKlineAPI) Symbol(symbol string) *MarketKlineAPI {
	api.req.Symbol = GetPointer(symbol)
	return api
}

// interval	true	string	時間粒度. 1,3,5,15,30,60,120,240,360,720,D,M,W
func (api *MarketKlineAPI) Interval(interval string) *MarketKlineAPI {
	api.req.Interval = GetPointer(interval)
	return api
}

// start	false	integer	開始時間戳 (毫秒)
func (api *MarketKlineAPI) Start(start int64) *MarketKlineAPI {
	api.req.Start = &start
	return api
}

// end	false	integer	結束時間戳 (毫秒)
func (api *MarketKlineAPI) End(end int64) *MarketKlineAPI {
	api.req.End = &end
	return api
}

// limit	false	integer	每頁數量限制. [1, 1000]. 默認: 200
func (api *MarketKlineAPI) Limit(limit int) *MarketKlineAPI {
	api.req.Limit = &limit
	return api
}

type MarketOrderBookReq struct {
	Category *string `json:"category"` //String	true	string	產品類型. spot, linear, inverse, option
	Symbol   *string `json:"symbol"`   //String	true	string	合約名稱
	Limit    *int    `json:"limit"`    //String	false	integer	深度限制. spot: [1, 200], 默認: 1. linear&inverse: [1, 500],默認: 25. option: [1, 25],默認: 1.
}

type MarketOrderBookAPI struct {
	client *PublicRestClient
	req    *MarketOrderBookReq
}

// category	true	string	產品類型. spot, linear, inverse, option
func (api *MarketOrderBookAPI) Category(category string) *MarketOrderBookAPI {
	api.req.Category = GetPointer(category)
	return api
}

// symbol	true	string	合約名稱
func (api *MarketOrderBookAPI) Symbol(symbol string) *MarketOrderBookAPI {
	api.req.Symbol = GetPointer(symbol)
	return api
}

// limit	false	integer	深度限制. spot: [1, 200], 默認: 1. linear&inverse: [1, 500],默認: 25. option: [1, 25],默認: 1
func (api *MarketOrderBookAPI) Limit(limit int) *MarketOrderBookAPI {
	api.req.Limit = &limit
	return api
}

// category	true	string	產品類型. spot,linear,inverse,option
// symbol	false	string	合約名稱
// baseCoin	false	string	交易幣種. 僅option有效, baseCoin和symbol必傳其中一個
// expDate	false	string	到期日. 舉例, 25DEC22. 僅option有效
type MarketTickersReq struct {
	Category *string `json:"category"` //String	true	string	產品類型. spot,linear,inverse,option
	Symbol   *string `json:"symbol"`   //String	false	string	合約名稱
	BaseCoin *string `json:"baseCoin"` //String	false	string	交易幣種. 僅option有效, baseCoin和symbol必傳其中一個
	ExpDate  *string `json:"expDate"`  //String	false	string	到期日. 舉例, 25DEC22. 僅option有效
}

type MarketTickersAPI struct {
	client *PublicRestClient
	req    *MarketTickersReq
}

// category	true	string	產品類型. spot,linear,inverse,option
func (api *MarketTickersAPI) Category(category string) *MarketTickersAPI {
	api.req.Category = GetPointer(category)
	return api
}

// symbol	false	string	合約名稱
func (api *MarketTickersAPI) Symbol(symbol string) *MarketTickersAPI {
	api.req.Symbol = GetPointer(symbol)
	return api
}

// baseCoin	false	string	交易幣種. 僅option有效, baseCoin和symbol必傳其中一個
func (api *MarketTickersAPI) BaseCoin(baseCoin string) *MarketTickersAPI {
	api.req.BaseCoin = GetPointer(baseCoin)
	return api
}

// expDate	false	string	到期日. 舉例, 25DEC22. 僅option有效
func (api *MarketTickersAPI) ExpDate(expDate string) *MarketTickersAPI {
	api.req.ExpDate = GetPointer(expDate)
	return api
}
