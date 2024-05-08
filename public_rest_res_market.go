package mybybitapi

import "github.com/shopspring/decimal"

type InstrumentsInfoResRow struct {
	Symbol        string        `json:"symbol"`        // 合約名稱
	BaseCoin      string        `json:"baseCoin"`      // 交易幣種
	QuoteCoin     string        `json:"quoteCoin"`     // 報價幣種
	SettleCoin    string        `json:"settleCoin"`    // 結算幣種
	Status        string        `json:"status"`        // 合約狀態
	LotSizeFilter LotSizeFilter `json:"lotSizeFilter"` // 數量屬性
	PriceFilter   PriceFilter   `json:"priceFilter"`   // 價格屬性
	LinearDetails
	OptionsDetails
	SpotDetails
}

type MarketInstrumentsInfoRes struct {
	Category       string                  `json:"category"`       // 產品類型
	NextPageCursor string                  `json:"nextPageCursor"` // 游標，用於翻頁
	List           []InstrumentsInfoResRow `json:"list"`           // 共有参数列表
}

// LinearDetails
type LinearDetails struct {
	UnifiedMarginTrade bool           `json:"unifiedMarginTrade"` // 是否支持統一保證金交易
	FundingInterval    int            `json:"fundingInterval"`    // 資金費率結算週期 (分鐘)
	CopyTrading        string         `json:"copyTrading"`        // 當前交易對是否支持帶單交易
	UpperFundingRate   string         `json:"upperFundingRate"`   // 資金費率上限
	LowerFundingRate   string         `json:"lowerFundingRate"`   // 資金費率下限
	ContractType       string         `json:"contractType"`       // 合約類型
	LaunchTime         string         `json:"launchTime"`         // 發佈時間 (ms)
	DeliveryTime       string         `json:"deliveryTime"`       // 交割時間 (ms). 僅對交割合約有效
	DeliveryFeeRate    string         `json:"deliveryFeeRate"`    // 交割費率. 僅對交割合約有效
	PriceScale         string         `json:"priceScale"`         // 價格精度
	LeverageFilter     LeverageFilter `json:"leverageFilter"`     // 槓桿屬性
}

// OptionsDetails
type OptionsDetails struct {
	OptionsType     string `json:"optionsType"`     // 期權類型. Call,Put
	LaunchTime      string `json:"launchTime"`      // 發佈時間 (ms)
	DeliveryTime    string `json:"deliveryTime"`    // 交割時間 (ms)
	DeliveryFeeRate string `json:"deliveryFeeRate"` // 交割費率
}

// SpotDetails
type SpotDetails struct {
	Innovation     string         `json:"innovation"`     // 是否屬於創新區交易對. `0`: 否, `1`: 是
	MarginTrading  string         `json:"marginTrading"`  // 該幣對是否支持槓桿交易
	RiskParameters RiskParameters `json:"riskParameters"` // 價格限制參數
}

type LeverageFilter struct {
	MinLeverage  string `json:"minLeverage"`  // 最小槓桿
	MaxLeverage  string `json:"maxLeverage"`  // 最大槓桿
	LeverageStep string `json:"leverageStep"` // 修改槓桿的步長
}

type PriceFilter struct {
	MinPrice string `json:"minPrice"` // 訂單最小價格
	MaxPrice string `json:"maxPrice"` // 訂單最大價格
	TickSize string `json:"tickSize"` // 修改價格的步長
}

type LotSizeFilter struct {
	MaxOrderQty         string `json:"maxOrderQty"`         // 單筆限價或PostOnly單最大下單量
	MinOrderQty         string `json:"minOrderQty"`         // 單筆訂單最小下單量
	MaxMktOrderQty      string `json:"maxMktOrderQty"`      // 單筆市價單最大下單量
	QtyStep             string `json:"qtyStep"`             // 修改下單量的步長
	PostOnlyMaxOrderQty string `json:"postOnlyMaxOrderQty"` // 廢棄, 請參照maxOrderQty
	MinNotionalValue    string `json:"minNotionalValue"`    // 訂單最小名義價值
	BasePrecision       string `json:"basePrecision"`       // 交易幣種精度
	QuotePrecision      string `json:"quotePrecision"`      // 報價幣種精度
	MinOrderAmt         string `json:"minOrderAmt"`         // 單筆訂單最小訂單額
	MaxOrderAmt         string `json:"maxOrderAmt"`         // 單筆訂單最大訂單額
}

type RiskParameters struct {
	LimitParameter  string `json:"limitParameter"`  // 限價單價格限制
	MarketParameter string `json:"marketParameter"` // 市價單價格限制
}

type MarketTimeRes struct {
	TimeSecond string `json:"timeSecond"` //Bybit服務器時間戳 (秒)
	TimeNano   string `json:"timeNano"`   //Bybit 服務器時間戳 (微秒)
}

type MarketKlineResRow struct {
	StartTime  string `json:"startTime"`  //蠟燭的開始時間戳 (毫秒)
	OpenPrice  string `json:"openPrice"`  //開始價格
	HighPrice  string `json:"highPrice"`  //最高價格
	LowPrice   string `json:"lowPrice"`   //最低價格
	ClosePrice string `json:"closePrice"` //結束價格. 如果蠟燭尚未結束，則表示為最新成交價格
	Volume     string `json:"volume"`     //交易量. 合約單位: 合約的張數. 現貨單位: 幣種的數量
	Turnover   string `json:"turnover"`   //交易額. 單位: 報價貨幣的數量
}

type MarketKlineRes struct {
	Category string              `json:"category"` //產品類型
	Symbol   string              `json:"symbol"`   //合約名稱
	List     []MarketKlineResRow `json:"list"`
}
type MarketKlineMiddleRow [7]interface{}
type MarketKlineMiddle struct {
	Symbol   string                 `json:"symbol"`
	Category string                 `json:"category"`
	List     []MarketKlineMiddleRow `json:"list"`
}

func (middleRow MarketKlineMiddle) ConvertToRes() MarketKlineRes {
	res := MarketKlineRes{
		Category: middleRow.Category,
		Symbol:   middleRow.Symbol,
		List:     make([]MarketKlineResRow, len(middleRow.List)),
	}
	for i, row := range middleRow.List {
		res.List[i] = MarketKlineResRow{
			StartTime:  row[0].(string),
			OpenPrice:  row[1].(string),
			HighPrice:  row[2].(string),
			LowPrice:   row[3].(string),
			ClosePrice: row[4].(string),
			Volume:     row[5].(string),
			Turnover:   row[6].(string),
		}
	}
	return res
}

type MarketOrderBookRes struct {
	Symbol string  `json:"s"`   //合約名稱
	Bids   []Books `json:"b"`   //買方. 按照價格從大到小
	Asks   []Books `json:"a"`   //賣方. 按照價格從小到大
	Ts     int64   `json:"ts"`  //行情服務生成數據時間戳（毫秒）
	U      int64   `json:"u"`   //表示數據連續性的id. 對於期貨, 它和wss推送裡的500檔的u對齊 對於現貨, 它和wss推送裡的200檔的u對齊
	Seq    int64   `json:"seq"` //撮合版本號 該字段可以用於關聯不同檔位的orderbook, 如果值越小, 則說明數據生成越早 期權目前不存在此字段
}
type MarketOrderBookMiddle struct {
	Symbol string        `json:"s"`   //合約名稱
	Bids   []BooksMiddle `json:"b"`   //買方. 按照價格從大到小
	Asks   []BooksMiddle `json:"a"`   //賣方. 按照價格從小到大
	Ts     int64         `json:"ts"`  //行情服務生成數據時間戳（毫秒）
	U      int64         `json:"u"`   //表示數據連續性的id. 對於期貨, 它和wss推送裡的500檔的u對齊 對於現貨, 它和wss推送裡的200檔的u對齊
	Seq    int64         `json:"seq"` //撮合版本號 該字段可以用於關聯不同檔位的orderbook, 如果值越小, 則說明數據生成越早 期權目前不存在此字段
}

type Books struct {
	Price    string `json:"price"`    //价格
	Quantity string `json:"quantity"` //合约张数或交易币的数量
}

func (b *Books) Float64Result() (float64, float64) {
	return stringToFloat64(b.Price), stringToFloat64(b.Quantity)
}

func (b *Books) DecimalResult() (decimal.Decimal, decimal.Decimal) {
	return decimal.RequireFromString(b.Price), decimal.RequireFromString(b.Quantity)
}

type BooksMiddle [2]interface{}

func (middle MarketOrderBookMiddle) ConvertToRes() MarketOrderBookRes {
	res := MarketOrderBookRes{
		Symbol: middle.Symbol,
		Bids:   make([]Books, len(middle.Bids)),
		Asks:   make([]Books, len(middle.Asks)),
		Ts:     middle.Ts,
		U:      middle.U,
		Seq:    middle.Seq,
	}
	for i, bid := range middle.Bids {
		res.Bids[i] = Books{
			Price:    bid[0].(string),
			Quantity: bid[1].(string),
		}
	}
	for i, ask := range middle.Asks {
		res.Asks[i] = Books{
			Price:    ask[0].(string),
			Quantity: ask[1].(string),
		}
	}
	return res
}

type MarketTickersResRow struct {
	//common 公有参数
	Symbol       string `json:"symbol"`       //合約名稱
	Bid1Price    string `json:"bid1Price"`    //買1價
	Bid1Size     string `json:"bid1Size"`     //買1價的數量
	Ask1Price    string `json:"ask1Price"`    //賣1價
	Ask1Size     string `json:"ask1Size"`     //賣1價的數量
	LastPrice    string `json:"lastPrice"`    //最新市場成交價
	PrevPrice24H string `json:"prevPrice24h"` //24小時前的整點市價
	Price24HPcnt string `json:"price24hPcnt"` //市場價格相對24h前變化的百分比
	HighPrice24H string `json:"highPrice24h"` //最近24小時的最高價
	LowPrice24H  string `json:"lowPrice24h"`  //最近24小時的最低價
	Turnover24H  string `json:"turnover24h"`  //最近24小時成交額
	Volume24H    string `json:"volume24h"`    //最近24小時成交量

	MarketTickersLinear
	MarketTickersOption
	MarketTickersSpot
}

// linear 独特参数
type MarketTickersLinear struct {
	IndexPrice             string `json:"indexPrice"`             //指數價格
	MarkPrice              string `json:"markPrice"`              //標記價格
	PrevPrice1H            string `json:"prevPrice1h"`            //1小時前的整點市價
	OpenInterest           string `json:"openInterest"`           //未平倉合約的數量
	OpenInterestValue      string `json:"openInterestValue"`      //未平倉合約的價值
	FundingRate            string `json:"fundingRate"`            //資金費率
	NextFundingTime        string `json:"nextFundingTime"`        //下次結算資金費用的時間 (毫秒)
	PredictedDeliveryPrice string `json:"predictedDeliveryPrice"` //預計交割價格. 交割前30分鐘有值
	BasisRate              string `json:"basisRate"`              //交割合約基差率
	DeliveryFeeRate        string `json:"deliveryFeeRate"`        //交割費率
	DeliveryTime           string `json:"deliveryTime"`           //交割時間戳 (毫秒)
	Basis                  string `json:"basis"`                  //交割合約基差
}

// option 独特参数
type MarketTickersOption struct {
	Bid1Iv          string `json:"bid1Iv"`          //買1價對應的iv
	Ask1Iv          string `json:"ask1Iv"`          //賣1價對應的iv
	MarkIv          string `json:"markIv"`          //標記價格對應的iv
	UnderlyingPrice string `json:"underlyingPrice"` //底層資產的價格
	TotalVolume     string `json:"totalVolume"`     //總成交量
	TotalTurnover   string `json:"totalTurnover"`   //總成交額
	Delta           string `json:"delta"`           //Delta
	Gamma           string `json:"gamma"`           //Gamma
	Vega            string `json:"vega"`            //Vega
	Theta           string `json:"theta"`           //Theta
	Change24H       string `json:"change24h"`       //過去24小時的變化
}

// spot 独特参数
type MarketTickersSpot struct {
	UsdIndexPrice string `json:"usdIndexPrice"` //USD指數價格 用於計算統一帳戶裡資產折算成USD價值的價格 若幣種不屬於抵押品幣種, 則返回空字符串 只有那些幣對名是"XXX/USDT"或者"XXX/USDC"有值
}

type MarketTickersRes struct {
	Category string                `json:"category"` //產品類型
	List     []MarketTickersResRow `json:"list"`
}
