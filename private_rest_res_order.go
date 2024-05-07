package mybybitapi

type OrderCommonRes struct {
	OrderId     string `json:"orderId"`     //訂單ID
	OrderLinkId string `json:"orderLinkId"` //用戶自定義訂單ID
}

type OrderCreateRes OrderCommonRes
type OrderAmendRes OrderCommonRes
type OrderCancelRes OrderCommonRes

type OrderCreateBatchRes struct {
	List []OrderCreateBatchResRow `json:"list"`
}

type OrderCreateBatchResRow struct {
	Category    string `json:"category"`    //產品類型
	Symbol      string `json:"symbol"`      //合約名稱
	OrderId     string `json:"orderId"`     //訂單Id
	OrderLinkId string `json:"orderLinkId"` //用戶自定義訂單Id
	CreateAt    string `json:"createAt"`    //訂單創建時間 (毫秒
}

type OrderAmendBatchRes struct {
	List []OrderAmendBatchResRow `json:"list"`
}

type OrderAmendBatchResRow struct {
	Category    string `json:"category"`    //產品類型
	Symbol      string `json:"symbol"`      //合約名稱
	OrderId     string `json:"orderId"`     //訂單Id
	OrderLinkId string `json:"orderLinkId"` //用戶自定義訂單Id
}

type OrderCancelBatchRes struct {
	List []OrderCancelBatchResRow `json:"list"`
}

type OrderCancelBatchResRow struct {
	Category    string `json:"category"`    //產品類型
	Symbol      string `json:"symbol"`      //合約名稱
	OrderId     string `json:"orderId"`     //訂單Id
	OrderLinkId string `json:"orderLinkId"` //用戶自定義訂單Id
}

type OrderCancelAllRes struct {
	List    []OrderCancelAllResRow `json:"list"`    //array	Object
	Success string                 `json:"success"` //string	"1": 成功, "0": 失敗
}

type OrderCancelAllResRow struct {
	OrderId     string `json:"orderId"`     //訂單Id
	OrderLinkId string `json:"orderLinkId"` //用戶自定義訂單Id
}

type OrderQueryCommon struct {
	OrderId            string `json:"orderId"`            //訂單Id
	OrderLinkId        string `json:"orderLinkId"`        //用戶自定義Id
	BlockTradeId       string `json:"blockTradeId"`       //Paradigm大宗交易Id
	Symbol             string `json:"symbol"`             //合約名稱
	Price              string `json:"price"`              //訂單價格
	Qty                string `json:"qty"`                //訂單數量
	Side               string `json:"side"`               //方向. Buy,Sell
	IsLeverage         string `json:"isLeverage"`         //是否借貸. 僅統一帳戶spot有效. 0: 否, 1: 是. 經典帳戶現貨交易不支持, 總是0
	PositionIdx        int    `json:"positionIdx"`        //倉位標識。用戶不同倉位模式
	OrderStatus        string `json:"orderStatus"`        //訂單狀態
	CreateType         string `json:"createType"`         //訂單創建類型
	CancelType         string `json:"cancelType"`         //訂單被取消類型
	RejectReason       string `json:"rejectReason"`       //拒絕原因. 經典帳戶現貨交易不支持
	AvgPrice           string `json:"avgPrice"`           //訂單平均成交價格
	LeavesQty          string `json:"leavesQty"`          //訂單剩餘未成交的數量. 經典帳戶現貨交易不支持
	LeavesValue        string `json:"leavesValue"`        //訂單剩餘未成交的價值. 經典帳戶現貨交易不支持
	CumExecQty         string `json:"cumExecQty"`         //訂單累計成交數量
	CumExecValue       string `json:"cumExecValue"`       //訂單累計成交價值. 經典帳戶現貨交易不支持
	CumExecFee         string `json:"cumExecFee"`         //訂單累計成交的手續費. 經典帳戶現貨交易不支持
	TimeInForce        string `json:"timeInForce"`        //執行策略
	OrderType          string `json:"orderType"`          //訂單類型. Market,Limit. 對於止盈止損單, 則表示為觸發後的訂單類型
	StopOrderType      string `json:"stopOrderType"`      //條件單類型
	OrderIv            string `json:"orderIv"`            //隱含波動率
	MarketUnit         string `json:"marketUnit"`         //統一帳戶現貨交易時給入參qty選擇的單位. baseCoin, quoteCoin
	TriggerPrice       string `json:"triggerPrice"`       //觸發價格. 若stopOrderType=TrailingStop, 則這是激活價格. 否則, 它是觸發價格
	TakeProfit         string `json:"takeProfit"`         //止盈價格
	StopLoss           string `json:"stopLoss"`           //止損價格
	TpslMode           string `json:"tpslMode"`           //止盈止損模式 Full: 全部倉位止盈止損, Partial: 部分倉位止盈止損. 現貨不返回該字段, 期權總是返回""
	OcoTriggerBy       string `json:"ocoTriggerBy"`       //現貨OCO訂單的觸發類型.OcoTriggerByUnknown, OcoTriggerByTp, OcoTriggerBySl. 經典帳戶現貨不支持該字段
	TpLimitPrice       string `json:"tpLimitPrice"`       //觸發止盈後轉換為限價單的價格
	SlLimitPrice       string `json:"slLimitPrice"`       //觸發止損後轉換為限價單的價格
	TpTriggerBy        string `json:"tpTriggerBy"`        //觸發止盈的價格類型
	SlTriggerBy        string `json:"slTriggerBy"`        //觸發止損的價格類型
	TriggerDirection   int    `json:"triggerDirection"`   //觸發方向. 1: 上漲, 2: 下跌
	TriggerBy          string `json:"triggerBy"`          //觸發價格的觸發類型
	LastPriceOnCreated string `json:"lastPriceOnCreated"` //下單時的市場價格
	ReduceOnly         bool   `json:"reduceOnly"`         //只減倉. true表明這是只減倉單
	CloseOnTrigger     bool   `json:"closeOnTrigger"`     //觸發後平倉委託. 什麼是觸發後平倉委託?
	PlaceType          string `json:"placeType"`          //下單類型, 僅期權使用. iv, price
	SmpType            string `json:"smpType"`            //SMP執行類型
	SmpGroup           int    `json:"smpGroup"`           //所屬Smp組ID. 如果uid不屬於任何組, 則默認為0
	SmpOrderId         string `json:"smpOrderId"`         //觸發此SMP執行的交易對手的 orderID
	CreatedTime        string `json:"createdTime"`        //創建訂單的時間戳 (毫秒)
	UpdatedTime        string `json:"updatedTime"`        //訂單更新的時間戳 (毫秒)
}

type OrderRealtimeRes struct {
	Category       string             `json:"category"`       //產品類型
	NextPageCursor string             `json:"nextPageCursor"` //游標，用於翻頁
	List           []OrderQueryCommon `json:"list"`           //array	Object
}

type OrderHistoryRes struct {
	Category       string             `json:"category"`       //產品類型
	NextPageCursor string             `json:"nextPageCursor"` //游標，用於翻頁
	List           []OrderQueryCommon `json:"list"`           //array	Object
}

type OrderExecutionListRes struct {
	Category       string                     `json:"category"`       //產品類型
	NextPageCursor string                     `json:"nextPageCursor"` //游標，用於翻頁
	List           []OrderExecutionListResRow `json:"list"`           //array	Object
}

type OrderExecutionListResRow struct {
	Symbol          string `json:"symbol"`          //合約名稱
	OrderId         string `json:"orderId"`         //訂單Id
	OrderLinkId     string `json:"orderLinkId"`     //用戶自定義訂單id. 經典帳戶現貨交易不支持
	Side            string `json:"side"`            //訂單方向.買： Buy,賣：Sell
	OrderPrice      string `json:"orderPrice"`      //訂單價格
	OrderQty        string `json:"orderQty"`        //訂單數量
	LeavesQty       string `json:"leavesQty"`       //剩餘委託未成交數量. 經典帳戶現貨交易不支持
	CreateType      string `json:"createType"`      //訂單創建類型
	OrderType       string `json:"orderType"`       //訂單類型. 市價單：Market,限價單：Limit
	StopOrderType   string `json:"stopOrderType"`   //条件单的订单类型。如果该订单不是条件单，则可能返回""或者UNKNOWN. 經典帳戶現貨交易不支持
	ExecFee         string `json:"execFee"`         //交易手續費. 您可以從這裡了解現貨手續費幣種信息
	ExecId          string `json:"execId"`          //成交Id
	ExecPrice       string `json:"execPrice"`       //成交價格
	ExecQty         string `json:"execQty"`         //成交數量
	ExecType        string `json:"execType"`        //交易類型. 經典帳戶現貨交易不支持
	ExecValue       string `json:"execValue"`       //成交價值. 經典帳戶現貨交易不支持
	ExecTime        string `json:"execTime"`        //成交時間（毫秒）
	FeeCurrency     string `json:"feeCurrency"`     //現貨手續費幣種 經典帳戶現貨交易不支持
	IsMaker         bool   `json:"isMaker"`         //是否是 Maker 訂單,true 為 maker 訂單，false 為 taker 訂單
	FeeRate         string `json:"feeRate"`         //手續費率. 經典帳戶現貨交易不支持
	TradeIv         string `json:"tradeIv"`         //隱含波動率，僅期權有效
	MarkIv          string `json:"markIv"`          //標記價格的隱含波動率，僅期權有效
	MarkPrice       string `json:"markPrice"`       //成交執行時，該 symbol 當時的標記價格. 經典帳戶現貨交易不支持
	IndexPrice      string `json:"indexPrice"`      //成交執行時，該 symbol 當時的指數價格，目前僅對期權業務有效
	UnderlyingPrice string `json:"underlyingPrice"` //成交執行時，該 symbol 當時的底層資產價格，僅期權有效
	BlockTradeId    string `json:"blockTradeId"`    //大宗交易的订单 ID ，使用 paradigm 进行大宗交易时生成的 ID
	ClosedSize      string `json:"closedSize"`      //平倉數量
	Seq             int    `json:"seq"`             //序列號, 用於關聯成交和倉位的更新
}
