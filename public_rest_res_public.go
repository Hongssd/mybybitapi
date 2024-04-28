package mybybitapi

type InstrumentsInfoResList struct {
	Symbol             string                        `json:"symbol"`             // 合约名称
	ContractType       string                        `json:"contractType"`       // 合约类型
	Status             string                        `json:"status"`             // 合约状态
	BaseCoin           string                        `json:"baseCoin"`           // 交易币种
	QuoteCoin          string                        `json:"quoteCoin"`          // 报价币种
	LaunchTime         string                        `json:"launchTime"`         // 发布时间 (ms)
	DeliveryTime       string                        `json:"deliveryTime"`       // 交割时间 (ms). 仅对交割合约有效
	DeliveryFeeRate    string                        `json:"deliveryFeeRate"`    // 交割费率. 仅对交割合约有效
	PriceScale         string                        `json:"priceScale"`         // 价格精度
	UnifiedMarginTrade bool                          `json:"unifiedMarginTrade"` // 是否支持统一保证金交易
	FundingInterval    int                           `json:"fundingInterval"`    // 资金费率结算周期 (分钟)
	CopyTrading        string                        `json:"copyTrading"`        // 当前交易对是否支持带单交易
	UpperFundingRate   string                        `json:"upperFundingRate"`   // 资金费率上限
	LowerFundingRate   string                        `json:"lowerFundingRate"`   // 资金费率下限
	SettleCoin         string                        `json:"settleCoin"`         //结算币种
	OptionsType        string                        `json:"optionsType"`        //期权类型. Call,Put
	Innovation         string                        `json:"innovation"`         //是否属于创新区交易对. `0`: 否, `1`: 是
	MarginTrading      string                        `json:"marginTrading"`      //该币对是否支持杠杆交易
	LeverageFilter     InstrumentsInfoLeverageFilter `json:"leverageFilter"`     //杠杆限制
	PriceFilter        InstrumentsInfoPriceFilter    `json:"priceFilter"`        //价格限制
	LotSizeFilter      InstrumentsInfoLotSizeFilter  `json:"lotSizeFilter"`      //下单量限制
	RiskParameters     InstrumentsInfoRiskParameters `json:"riskParameters"`     //风险参数
}

type InstrumentsInfoLeverageFilter struct {
	MinLeverage  string `json:"minLeverage"`  // 最小杠杆
	MaxLeverage  string `json:"maxLeverage"`  // 最大杠杆
	LeverageStep string `json:"leverageStep"` // 修改杠杆的步长
}

type InstrumentsInfoPriceFilter struct {
	MinPrice string `json:"minPrice"` // 订单最小价格
	MaxPrice string `json:"maxPrice"` // 订单最大价格
	TickSize string `json:"tickSize"` // 修改价格的步长
}

type InstrumentsInfoLotSizeFilter struct {
	MaxOrderQty         string `json:"maxOrderQty"`         // 单笔限价或PostOnly单最大下单量
	MaxMktOrderQty      string `json:"maxMktOrderQty"`      // 单笔市价单最大下单量
	MinOrderQty         string `json:"minOrderQty"`         // 单笔订单最小下单量
	QtyStep             string `json:"qtyStep"`             // 修改下单量的步长
	PostOnlyMaxOrderQty string `json:"postOnlyMaxOrderQty"` // 废弃, 请参照maxOrderQty
	MinNotionalValue    string `json:"minNotionalValue"`    // 订单最小名义价值
	BasePrecision       string `json:"basePrecision"`       //交易币种精度
	QuotePrecision      string `json:"quotePrecision"`      //报价币种精度
	MinOrderAmt         string `json:"minOrderAmt"`         //单笔订单最小订单额
	MaxOrderAmt         string `json:"maxOrderAmt"`         //单笔订单最大订单额
}

type InstrumentsInfoRiskParameters struct {
	LimitParameter  string `json:"limitParameter"`  //限价单价格限制. 如果预估成交价与最新成交价的偏差大于限定的百分比，则该笔限价单将会被限制下单。 举例说明，接口返回0.1, 则表示限价为10%，您下的限价单的下单价格不能超过最后成交价格（LTP）的110%； 您卖出的限价单不能低于LTP的90 %。
	MarketParameter string `json:"marketParameter"` //市价单价格限制. 如果预估成交价与最新成交价的偏差大于限定的百分比，市价单将被部分执行. 举例说明，接口返回0.05, 则表示限价为5%. 假如当前MNT /USDT限价为5%，用户下单10 万USDT 市价买入（当前最新成交价为2 USDT）。 凡在2.1 USDT以上成交的部分将被取消。 假设只有价值2万 USDT的MNT可以在2.1 USDT及以下的价格成交，剩余的8万 USDT订单价值将因为偏差超过5%的阈值而被取消。
}

type MarketInstrumentsInfoRes struct {
	Category       string                   `json:"category"`       //产品类型
	NextPageCursor string                   `json:"nextPageCursor"` //游标，用于翻页
	List           []InstrumentsInfoResList `json:"list"`           //Object
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
