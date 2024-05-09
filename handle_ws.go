package mybybitapi

type CommonWsRes struct {
	Topic string `json:"topic"` //Topic名
	Ts    int64  `json:"ts"`    //行情服務生成數據的時間戳 (毫秒)
	Type  string `json:"type"`  //數據類型. snapshot,delta
}

type WsKlineMiddle struct {
	CommonWsRes
	Data []WsKlineData `json:"data"`
}

type WsKlineData struct {
	Start     int64  `json:"start"`     //開始時間戳 (毫秒)
	End       int64  `json:"end"`       //結束時間戳 (毫秒)
	Interval  string `json:"interval"`  //K線粒度
	Open      string `json:"open"`      //開盤價
	Close     string `json:"close"`     //收盤價
	High      string `json:"high"`      //最高價
	Low       string `json:"low"`       //最低價
	Volume    string `json:"volume"`    //交易量
	Turnover  string `json:"turnover"`  //交易額
	Confirm   bool   `json:"confirm"`   //是否確認
	Timestamp int64  `json:"timestamp"` //蠟燭中最後一筆淨值時間戳 (毫秒)
}

type WsKline struct {
	CommonWsRes
	WsKlineData
}

func handleWsKline(data []byte) (*WsKline, error) {

	wsKlineMiddle := WsKlineMiddle{}
	err := json.Unmarshal(data, &wsKlineMiddle)
	if err != nil {
		return nil, err
	}
	klineData := wsKlineMiddle.Data[0]
	kline := WsKline{
		CommonWsRes: wsKlineMiddle.CommonWsRes,
		WsKlineData: klineData,
	}
	return &kline, nil
}

type WsDepthMiddle struct {
	CommonWsRes
	Cts  int64             `json:"cts"`
	Data WsDepthDataMiddle `json:"data"`
}

type WsDepth struct {
	CommonWsRes
	Cts int64 `json:"cts"`
	WsDepthData
}

type WsDepthDataMiddle struct {
	Symbol string     `json:"s"`
	U      int64      `json:"u"`
	Seq    int64      `json:"seq"`
	Bids   [][]string `json:"b"`
	Asks   [][]string `json:"a"`
}

type WsDepthData struct {
	Symbol string  `json:"s"`
	U      int64   `json:"u"`
	Seq    int64   `json:"seq"`
	Bids   []Books `json:"b"`
	Asks   []Books `json:"a"`
}

func handleWsDepth(data []byte) (*WsDepth, error) {
	wsDepthMiddle := WsDepthMiddle{}
	err := json.Unmarshal(data, &wsDepthMiddle)
	if err != nil {
		return nil, err
	}

	bids := []Books{}
	asks := []Books{}

	for _, b := range wsDepthMiddle.Data.Bids {
		if len(b) != 2 {
			continue
		}
		bids = append(bids, Books{
			Price:    b[0],
			Quantity: b[1],
		})
	}

	for _, a := range wsDepthMiddle.Data.Asks {
		if len(a) != 2 {
			continue
		}
		asks = append(asks, Books{
			Price:    a[0],
			Quantity: a[1],
		})
	}

	wsDepthData := WsDepthData{
		Symbol: wsDepthMiddle.Data.Symbol,
		U:      wsDepthMiddle.Data.U,
		Seq:    wsDepthMiddle.Data.Seq,
		Bids:   bids,
		Asks:   asks,
	}

	wsDepth := WsDepth{
		CommonWsRes: wsDepthMiddle.CommonWsRes,
		Cts:         wsDepthMiddle.Cts,
		WsDepthData: wsDepthData,
	}

	return &wsDepth, nil
}

type WsTrade struct {
	CommonWsRes
	Data []WsTradeData `json:"data"`
}

type WsTradeData struct {
	Timestamp  int64  `json:"T"`   //成交時間戳 (毫秒)
	Symbol     string `json:"s"`   //合約名稱
	Side       string `json:"S"`   //吃單方向. Buy,Sell
	Volume     string `json:"v"`   //成交數量
	Price      string `json:"p"`   //成交價格
	LastChange string `json:"L"`   //價格變化的方向. 期權沒有該字段
	TradeId    string `json:"i"`   //成交Id
	BlockTrade bool   `json:"BT"`  //成交類型是否為大宗交易
	MarkPrice  string `json:"mP"`  //標記價格, 期權的特有字段
	IndexPrice string `json:"iP"`  //指數價格, 期權的特有字段
	MarkIv     string `json:"mIv"` //標記iv, 期權的特有字段
	Iv         string `json:"iv"`  //iv, 期權的特有字段
}

func handleWsTrade(data []byte) (*WsTrade, error) {
	wsTrade := WsTrade{}
	err := json.Unmarshal(data, &wsTrade)
	if err != nil {
		return nil, err
	}
	return &wsTrade, nil
}

type WsOrder struct {
	Id           string        `json:"id"`           //消息id
	Topic        string        `json:"topic"`        //Topic名
	CreationTime int64         `json:"creationTime"` //消息數據創建時間
	Data         []WsOrderData `json:"data"`         //Object
}

type WsOrderData struct {
	Category           string `json:"category"`           //產品類型 統一帳戶: spot, linear, inverse, option 經典帳戶: spot, linear, inverse.
	OrderId            string `json:"orderId"`            //訂單ID
	OrderLinkId        string `json:"orderLinkId"`        //用戶自定義ID
	IsLeverage         string `json:"isLeverage"`         //是否借貸. 僅統一帳戶spot有效. 0: 否, 1: 是. 經典帳戶現貨交易不支持, 總是0
	BlockTradeId       string `json:"blockTradeId"`       //大宗交易訂單Id
	Symbol             string `json:"symbol"`             //合約名稱
	Price              string `json:"price"`              //訂單價格
	Qty                string `json:"qty"`                //訂單數量
	Side               string `json:"side"`               //方向. Buy,Sell
	PositionIdx        int    `json:"positionIdx"`        //倉位標識。用戶不同倉位模式
	OrderStatus        string `json:"orderStatus"`        //訂單狀態
	CreateType         string `json:"createType"`         //訂單創建類型 僅作用於category=linear 或 inverse 現貨、期權不返回該字段
	CancelType         string `json:"cancelType"`         //訂單被取消類型. 經典帳戶現貨交易不支持
	RejectReason       string `json:"rejectReason"`       //拒絕原因. 經典帳戶現貨交易不支持
	AvgPrice           string `json:"avgPrice"`           //訂單平均成交價格 不存在avg price場景的訂單返回"", 以及經典帳戶下部分成交但最終被手動取消的訂單 經典帳戶現貨交易: 該字段不支持, 總是""
	LeavesQty          string `json:"leavesQty"`          //訂單剩餘未成交的數量. 經典帳戶現貨交易不支持
	LeavesValue        string `json:"leavesValue"`        //訂單剩餘未成交的價值. 經典帳戶現貨交易不支持
	CumExecQty         string `json:"cumExecQty"`         //訂單累計成交數量
	CumExecValue       string `json:"cumExecValue"`       //訂單累計成交價值
	CumExecFee         string `json:"cumExecFee"`         //訂單累計成交的手續費 經典帳戶spot: 表示的是最近一次成交的手續費. 升級到統一帳戶後, 您可以使用成交頻道中的execFee字段來獲取每次成交的手續費
	FeeCurrency        string `json:"feeCurrency"`        //現貨交易的手續費幣種. 可以從這裡了解現貨交易的手續費幣種規則
	TimeInForce        string `json:"timeInForce"`        //執行策略
	OrderType          string `json:"orderType"`          //訂單類型. Market,Limit. 對於止盈止損單, 則表示為觸發後的訂單類型
	StopOrderType      string `json:"stopOrderType"`      //條件單類型
	OcoTriggerBy       string `json:"ocoTriggerBy"`       //現貨OCO訂單的觸發類型.OcoTriggerByUnknown, OcoTriggerByTp, OcoTriggerBySl. 經典帳戶現貨不支持該字段
	OrderIv            string `json:"orderIv"`            //隱含波動率
	MarketUnit         string `json:"marketUnit"`         //統一帳戶現貨交易時給入參qty選擇的單位. baseCoin, quoteCoin
	TriggerPrice       string `json:"triggerPrice"`       //觸發價格. 若stopOrderType=TrailingStop, 則這是激活價格. 否則, 它是觸發價格
	TakeProfit         string `json:"takeProfit"`         //止盈價格
	StopLoss           string `json:"stopLoss"`           //止損價格
	TpslMode           string `json:"tpslMode"`           //止盈止損模式 Full: 全部倉位止盈止損, Partial: 部分倉位止盈止損. 現貨不返回該字段, 期權總是返回""
	TpLimitPrice       string `json:"tpLimitPrice"`       //觸發止盈後轉換為限價單的價格
	SlLimitPrice       string `json:"slLimitPrice"`       //觸發止損後轉換為限價單的價格
	TpTriggerBy        string `json:"tpTriggerBy"`        //觸發止盈的價格類型
	SlTriggerBy        string `json:"slTriggerBy"`        //觸發止損的價格類型
	TriggerDirection   int    `json:"triggerDirection"`   //觸發方向. 1: 上漲, 2: 下跌
	TriggerBy          string `json:"triggerBy"`          //觸發價格的觸發類型
	LastPriceOnCreated string `json:"lastPriceOnCreated"` //下單時的市場價格
	ReduceOnly         bool   `json:"reduceOnly"`         //只減倉. true表明這是只減倉單
	CloseOnTrigger     bool   `json:"closeOnTrigger"`     //觸發後平倉委託. 什麼是觸發後平倉委託?
	PlaceType          string `json:"placeType"`          //Place type, option used. iv, price
	SmpType            string `json:"smpType"`            //SMP執行類型
	SmpGroup           int    `json:"smpGroup"`           //所屬Smp組ID. 如果uid不屬於任何組, 則默認為0
	SmpOrderId         string `json:"smpOrderId"`         //觸發此SMP執行的交易對手的 orderID
	CreatedTime        string `json:"createdTime"`        //創建訂單的時間戳 (毫秒)
	UpdatedTime        string `json:"updatedTime"`        //訂單更新的時間戳 (毫秒)
}

func handleWsOrder(data []byte) (*WsOrder, error) {
	wsOrder := WsOrder{}
	err := json.Unmarshal(data, &wsOrder)
	if err != nil {
		return nil, err
	}
	return &wsOrder, nil
}

type WsWallet struct {
	Id           string         `json:"id"`           //消息id
	Topic        string         `json:"topic"`        //Topic名
	CreationTime int64          `json:"creationTime"` //消息數據創建時間
	Data         []WsWalletData `json:"data"`         //Object
}

type WsWalletData struct {
	AccountType            string             `json:"accountType"`            //帳戶類型 統一帳戶: UNIFIED(現貨/USDT和USDC永續/期權), CONTRACT(反向合約) 經典帳戶: CONTRACT, SPOT
	AccountLTV             string             `json:"accountLTV"`             //帳戶LTV account total borrowed size / (account total equity + account total borrowed size 非統一保證金模式和統一帳戶(反向)以及統一帳戶(逐倉模式)，該字段返回為空
	AccountIMRate          string             `json:"accountIMRate"`          //帳戶初始保證金率 Account Total Initial Margin Base Coin / Account Margin Balance Base Coin 非統一保證金模式和統一帳戶(反向)以及統一帳戶(逐倉模式)，該字段返回為空
	AccountMMRate          string             `json:"accountMMRate"`          //帳戶維持保證金率 Account Total Maintenance Margin Base Coin / Account Margin Balance Base Coin 非統一保證金模式和統一帳戶(反向)以及統一帳戶(逐倉模式)，該字段返回為空
	TotalEquity            string             `json:"totalEquity"`            //賬戶維度換算成usd的資產淨值 Account Margin Balance Base Coin + Account Option Value Base Coin 非統一保證金模式和統一帳戶(反向)下，該字段返回為空
	TotalWalletBalance     string             `json:"totalWalletBalance"`     //賬戶維度換算成usd的產錢包餘額 ∑ Asset Wallet Balance By USD value of each asset 非統一保證金模式和統一帳戶(反向)以及統一帳戶(逐倉模式)，該字段返回為空
	TotalMarginBalance     string             `json:"totalMarginBalance"`     //賬戶維度換算成usd的保證金餘額 totalWalletBalance + totalPerpUPL 非統一保證金模式和統一帳戶(反向)以及統一帳戶(逐倉模式)，該字段返回為空
	TotalAvailableBalance  string             `json:"totalAvailableBalance"`  //賬戶維度換算成usd的可用餘額 RM：totalMarginBalance - totalInitialMargin 非統一保證金模式和統一帳戶(反向)以及統一帳戶(逐倉模式)，該字段返回為空
	TotalPerpUPL           string             `json:"totalPerpUPL"`           //賬戶維度換算成usd的永續和USDC交割合約的浮動盈虧 ∑ Each perp and USDC Futures upl by base coin 非統一保證金模式以及統一帳戶(反向)下，該字段返回為空
	TotalInitialMargin     string             `json:"totalInitialMargin"`     //賬戶維度換算成usd的總初始保證金 ∑ Asset Total Initial Margin Base Coin 非統一保證金模式和統一帳戶(反向)以及統一帳戶(逐倉模式)，該字段返回為空
	TotalMaintenanceMargin string             `json:"totalMaintenanceMargin"` //賬戶維度換算成usd的總維持保證金 ∑ Asset Total Maintenance Margin Base Coin 非統一保證金模式和統一帳戶(反向)以及統一帳戶(逐倉模式)，該字段返回為空
	Coin                   []WsWalletDataCoin `json:"coin"`                   //Object. 幣種列表
}

type WsWalletDataCoin struct {
	Coin                string `json:"coin"`                //幣種名稱，例如 BTC，ETH，USDT，USDC
	Equity              string `json:"equity"`              //當前幣種的資產淨值
	UsdValue            string `json:"usdValue"`            //當前幣種折算成 usd 的價值，如果該幣種不能作為保證金的抵押品，則該數值為0
	WalletBalance       string `json:"walletBalance"`       //當前幣種的錢包餘額
	Free                string `json:"free"`                //經典帳戶現貨錢包的可用餘額. 經典帳戶現貨錢包的獨有字段
	Locked              string `json:"locked"`              //現貨掛單凍結金額
	SpotHedgingQty      string `json:"spotHedgingQty"`      //用於組合保證金(PM)現貨對衝的數量, 截斷至8為小數, 默認為0 統一帳戶的獨有字段
	BorrowAmount        string `json:"borrowAmount"`        //當前幣種的已用借貸額度
	AvailableToBorrow   string `json:"availableToBorrow"`   //由於母子共享借貸限額, 該字段已廢棄, 總是返回"". 請通過查詢抵押品信息接口查詢availableToBorrow
	AvailableToWithdraw string `json:"availableToWithdraw"` //當前幣種的可劃轉提現金額
	AccruedInterest     string `json:"accruedInterest"`     //當前幣種的預計要在下一個利息週期收取的利息金額
	TotalOrderIM        string `json:"totalOrderIM"`        //以當前幣種結算的訂單委託預佔用保證金. 組合保證金模式下，該字段返回空字符串
	TotalPositionIM     string `json:"totalPositionIM"`     //以當前幣種結算的所有倉位起始保證金求和 + 所有倉位的預佔用平倉手續費. 組合保證金模式下，該字段返回空字符串
	TotalPositionMM     string `json:"totalPositionMM"`     //以當前幣種結算的所有倉位維持保證金求和. 組合保證金模式下，該字段返回空字符串
	UnrealisedPnl       string `json:"unrealisedPnl"`       //以當前幣種結算的所有倉位的未結盈虧之和
	CumRealisedPnl      string `json:"cumRealisedPnl"`      //以當前幣種結算的所有倉位的累計已結盈虧之和
	Bonus               string `json:"bonus"`               //體驗金. UNIFIED帳戶的獨有字段
	MarginCollateral    bool   `json:"marginCollateral"`    //是否可作為保證金抵押幣種(平台維度), true: 是. false: 否 當marginCollateral=false時, 則collateralSwitch無意義
	CollateralSwitch    bool   `json:"collateralSwitch"`    //用戶是否開啟保證金幣種抵押(用戶維度), true: 是. false: 否 僅當marginCollateral=true時, 才能主動選擇開關抵押
}

func handleWsWallet(data []byte) (*WsWallet, error) {
	wsWallet := WsWallet{}
	err := json.Unmarshal(data, &wsWallet)
	if err != nil {
		return nil, err
	}
	return &wsWallet, nil
}

type WsPosition struct {
	Id           string           `json:"id"`           //消息id
	Topic        string           `json:"topic"`        //Topic名
	CreationTime int64            `json:"creationTime"` //消息數據創建時間
	Data         []WsPositionData `json:"data"`         //Object
}

type WsPositionData struct {
	Category               string `json:"category"`               //產品類型
	Symbol                 string `json:"symbol"`                 //合約名稱
	Side                   string `json:"side"`                   //持倉方向. Buy,Sell
	Size                   string `json:"size"`                   //持倉大小
	PositionIdx            int    `json:"positionIdx"`            //倉位標識
	TradeMode              int    `json:"tradeMode"`              //交易模式 0: 全倉, 1: 逐倉
	PositionValue          string `json:"positionValue"`          //倉位價值
	RiskId                 int    `json:"riskId"`                 //風險限額id 注意: 組合保證金模式下，該字段返回0，風險限額規則失效
	RiskLimitValue         string `json:"riskLimitValue"`         //風險限額id對應的風險限額度 注意: 組合保證金模式下，該字段返回空字符串，風險限額規則失效
	EntryPrice             string `json:"entryPrice"`             //入場價
	MarkPrice              string `json:"markPrice"`              //標記價
	Leverage               string `json:"leverage"`               //槓桿 注意: 組合保證金模式下，該字段返回""，槓桿規則失效
	PositionBalance        string `json:"positionBalance"`        //倉位保證金 注意: 組合保證金模式(PM)下, 該字段返回為空字符串
	AutoAddMargin          int    `json:"autoAddMargin"`          //是否自動追加保證金 0: 否, 1: 是
	PositionMM             string `json:"positionMM"`             //倉位維持保證金 注意: 組合保證金模式下，該字段返回空字符串
	PositionIM             string `json:"positionIM"`             //倉位初始保證金 注意: 組合保證金模式下，該字段返回空字符串
	LiqPrice               string `json:"liqPrice"`               //倉位強平價格
	BustPrice              string `json:"bustPrice"`              //預估破產價
	TpslMode               string `json:"tpslMode"`               //該字段廢棄, 無意義, 總是返回"Full"
	TakeProfit             string `json:"takeProfit"`             //止盈價格
	StopLoss               string `json:"stopLoss"`               //止損價格
	TrailingStop           string `json:"trailingStop"`           //追蹤止損
	UnrealisedPnl          string `json:"unrealisedPnl"`          //未結盈虧
	SessionAvgPrice        string `json:"sessionAvgPrice"`        //USDC合約平均持倉價格
	Delta                  string `json:"delta"`                  //Delta 期權的獨有字段
	Gamma                  string `json:"gamma"`                  //Gamma 期權的獨有字段
	Vega                   string `json:"vega"`                   //Vega 期權的獨有字段
	Theta                  string `json:"theta"`                  //Theta 期權的獨有字段
	CurRealisedPnl         string `json:"curRealisedPnl"`         //當前持倉的已結盈虧
	CumRealisedPnl         string `json:"cumRealisedPnl"`         //累計已结盈亏 期貨: 是從第一次開始有持倉加總的已結盈虧 期權: 它是本次持倉的加總已結盈虧
	PositionStatus         string `json:"positionStatus"`         //倉位狀態 Normal,Liq, Adl
	AdlRankIndicator       int    `json:"adlRankIndicator"`       //自動減倉燈
	IsReduceOnly           bool   `json:"isReduceOnly"`           //僅當Bybit需要降低某個Symbol的風險限額時有用 true: 僅允許減倉操作. false: 沒有交易限制 表示您的倉位在系統調整時處於風險水平之下 僅對逐倉和全倉的期貨倉位有意義
	MmrSysUpdatedTime      string `json:"mmrSysUpdatedTime"`      //僅當Bybit需要降低某個Symbol的風險限額時有用 當isReduceOnly=true: 這個時間戳表示系統強制修改MMR的時間 當isReduceOnly=false: 若不為空, 則表示系統已經完成了MMR調整的時間 僅當系統調整才會賦值, 對於主動的調整, 不會在這裡展示時間戳 默認為"", 但如果曾經這個symbol有過系統降檔的操作, 那麼這裡會顯示上一次操作的時間 僅對逐倉和全倉的期貨倉位有意義
	LeverageSysUpdatedTime string `json:"leverageSysUpdatedTime"` //僅當Bybit需要降低某個Symbol的風險限額時有用 當isReduceOnly=true: 這個時間戳表示系統強制修改槓桿的時間 當isReduceOnly=false: 若不為空, 則表示系統已經完成了槓桿調整的時間 僅當系統調整才會賦值, 對於主動的調整, 不會在這裡展示時間戳 默認為"", 但如果曾經這個symbol有過系統降檔的操作, 那麼這裡會顯示上一次操作的時間 僅對逐倉和全倉的期貨倉位有意義
	CreatedTime            string `json:"createdTime"`            //倉位創建時間戳 (毫秒)
	UpdatedTime            string `json:"updatedTime"`            //倉位數據更新時間戳 (毫秒)
	Seq                    int64  `json:"seq"`                    //序列號, 用於關聯成交和倉位的更新 不同的幣對會存在相同seq, 可以使用seq + symbol來做唯一性識別 如果該幣對從未被交易過, 查詢時則會返回"-1" 對於更新槓桿、更新風險限額等非交易行為, 將會返回上一次成交時更新的seq
}

func handleWsPosition(data []byte) (*WsPosition, error) {
	wsPosition := WsPosition{}
	err := json.Unmarshal(data, &wsPosition)
	if err != nil {
		return nil, err
	}
	return &wsPosition, nil
}

// id	string	消息id
// topic	string	Topic名
// creationTime	number	消息數據創建時間
// data	array	Object
// > category	string	產品類型
// 統一帳戶: spot, linear, iverse, option
// 經典帳戶: spot, linear, inverse.
// > symbol	string	合約名稱
// > isLeverage	string	是否借貸. 僅統一帳戶spot有效. 0: 否, 1: 是. 經典帳戶現貨交易不支持, 總是0
// > orderId	string	訂單ID
// > orderLinkId	string	用戶自定義訂單ID
// > side	string	訂單方向.買：Buy,賣：Sell
// > orderPrice	string	訂單價格. 經典帳戶現貨交易不支持
// > orderQty	string	訂單數量. 經典帳戶現貨交易不支持
// > leavesQty	string	剩餘委託未成交數量. 經典帳戶現貨交易不支持
// > createType	string	訂單創建類型
// 統一帳戶: 僅作用於category=linear 或 inverse
// 經典帳戶: 總是返回""
// 現貨、期權不返回該字段
// > orderType	string	訂單類型. 市價單：Market,限價單：Limit
// > stopOrderType	string	条件单的订单类型。如果该订单不是条件单，则不会返回任何类型. 經典帳戶現貨交易不支持
// > execFee	string	交易手續費. 您可以從這裡了解現貨手續費幣種信息. 經典帳戶現貨交易不支持
// > execId	string	成交Id
// > execPrice	string	成交價格
// > execQty	string	成交數量
// > execType	string	成交類型. 經典帳戶現貨交易不支持
// > execValue	string	成交價值. 經典帳戶現貨交易不支持
// > execTime	string	成交時間（毫秒）
// > isMaker	Bool	是否是 Maker 訂單,true 為 maker 訂單，false 為 taker 訂單
// > feeRate	string	手續費率. 經典帳戶現貨交易不支持
// > tradeIv	string	隱含波動率，僅期權有效
// > markIv	string	標記價格的隱含波動率，僅期權有效
// > markPrice	string	成交執行時，該 symbol 當時的標記價格. 目前僅對期權業務有效
// > indexPrice	string	成交執行時，該 symbol 當時的指數價格，目前僅對期權業務有效
// > underlyingPrice	string	成交執行時，該 symbol 當時的底層資產價格，僅期權有效
// > blockTradeId	string	大宗交易的订单 ID ，使用 paradigm 进行大宗交易时生成的 ID
// > closedSize	string	平倉數量
// > seq	long	序列號, 用於關聯成交和倉位的更新
// 同一時間有多筆成交, seq相同
// 不同的幣對會存在相同seq, 可以使用seq + symbol來做唯一性識別
type WsExecution struct {
	Id           string            `json:"id"`           //消息id
	Topic        string            `json:"topic"`        //Topic名
	CreationTime int64             `json:"creationTime"` //消息數據創建時間
	Data         []WsExecutionData `json:"data"`         //Object
}

type WsExecutionData struct {
	Category        string `json:"category"`        //產品類型 統一帳戶: spot, linear, inverse, option 經典帳戶: spot, linear, inverse
	Symbol          string `json:"symbol"`          //合約名稱
	IsLeverage      string `json:"isLeverage"`      //是否借貸. 僅統一帳戶spot有效. 0: 否, 1: 是. 經典帳戶現貨交易不支持, 總是0
	OrderId         string `json:"orderId"`         //訂單ID
	OrderLinkId     string `json:"orderLinkId"`     //用戶自定義訂單ID
	Side            string `json:"side"`            //訂單方向.買：Buy,賣：Sell
	OrderPrice      string `json:"orderPrice"`      //訂單價格. 經典帳戶現貨交易不支持
	OrderQty        string `json:"orderQty"`        //訂單數量. 經典帳戶現貨交易不支持
	LeavesQty       string `json:"leavesQty"`       //剩餘委託未成交數量. 經典帳戶現貨交易不支持
	CreateType      string `json:"createType"`      //訂單創建類型 統一帳戶: 僅作用於category=linear 或 inverse 經典帳戶: 總是返回"" 現貨、期權不返回該字段
	OrderType       string `json:"orderType"`       //訂單類型. 市價單：Market,限價單：Limit
	StopOrderType   string `json:"stopOrderType"`   //条件单的订单类型。如果该订单不是条件单，则不会返回任何类型. 經典帳戶現貨交易不支持
	ExecFee         string `json:"execFee"`         //交易手續費. 您可以從這裡了解現貨手續費幣種信息. 經典帳戶現貨交易不支持
	ExecId          string `json:"execId"`          //成交Id
	ExecPrice       string `json:"execPrice"`       //成交價格
	ExecQty         string `json:"execQty"`         //成交數量
	ExecType        string `json:"execType"`        //成交類型. 經典帳戶現貨交易不支持
	ExecValue       string `json:"execValue"`       //成交價值
	ExecTime        string `json:"execTime"`        //成交時間（毫秒）
	IsMaker         bool   `json:"isMaker"`         //是否是 Maker 訂單,true 為 maker 訂單，false 為 taker 訂單
	FeeRate         string `json:"feeRate"`         //手續費率. 經典帳戶現貨交易不支持
	TradeIv         string `json:"tradeIv"`         //隱含波動率，僅期權有效
	MarkIv          string `json:"markIv"`          //標記價格的隱含波動率，僅期權有效
	MarkPrice       string `json:"markPrice"`       //成交執行時，該 symbol 當時的標記價格. 目前僅對期權業務有效
	IndexPrice      string `json:"indexPrice"`      //成交執行時，該 symbol 當時的指數價格，目前僅對期權業務有效
	UnderlyingPrice string `json:"underlyingPrice"` //成交執行時，該 symbol 當時的底層資產價格，僅期權有效
	BlockTradeId    string `json:"blockTradeId"`    //大宗交易的订单 ID ，使用 paradigm 进行大宗交易时生成的 ID
	ClosedSize      string `json:"closedSize"`      //平倉數量
	Seq             int64  `json:"seq"`             //序列號, 用於關聯成交和倉位的更新 同一時間有多筆成交, seq相同 不同的幣對會存在相同seq, 可以使用seq + symbol來做唯一性識別
}

func handleWsExecution(data []byte) (*WsExecution, error) {
	wsExecution := WsExecution{}
	err := json.Unmarshal(data, &wsExecution)
	if err != nil {
		return nil, err
	}
	return &wsExecution, nil
}

func handleWsDoOrderResult[T OrderResType](data []byte) (*WsOrderResult[T], error) {
	wsOrderResult := WsOrderResult[T]{}
	err := json.Unmarshal(data, &wsOrderResult)
	if err != nil {
		return nil, err
	}
	return &wsOrderResult, nil
}
