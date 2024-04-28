package mybybitapi

type PublicRestAPI int

const (
	MarketInstrumentsInfo PublicRestAPI = iota //查詢可交易產品的規格信息
	MarketTime                                 //Bybit服務器時間
	MarketKline                                //查詢市場價格K線數據
	MarketOrderBook                            //Order Book (深度)

)

var PublicRestAPIMap = map[PublicRestAPI]string{
	//GET 查詢可交易產品的規格信息
	MarketInstrumentsInfo: "/v5/market/instruments-info", //GET 查詢可交易產品的規格信息
	MarketTime:            "/v5/market/time",             //GET Bybit服務器時間
	MarketKline:           "/v5/market/kline",            //GET 查詢市場價格K線數據
	MarketOrderBook:       "/v5/market/orderbook",        //GET Order Book (深度)

}
