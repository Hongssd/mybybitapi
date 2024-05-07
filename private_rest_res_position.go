package mybybitapi

type PositionListRes struct {
	Category       string               `json:"category"` //產品類型
	List           []PositionListResRow `json:"list"`
	NextPageCursor string               `json:"nextPageCursor"` //游標，用於翻頁
}

type PositionListResRow struct {
	PositionIdx            int    `json:"positionIdx"`            //倉位標識符, 用于在不同仓位模式下标识仓位
	RiskId                 int    `json:"riskId"`                 //风险限额ID，參見風險限額接口. 注意：若賬戶為組合保證金模式(PM)，該字段返回0，風險限額規則失效
	RiskLimitValue         string `json:"riskLimitValue"`         //當前風險限額ID對應的持倉限制量. 注意：若賬戶為組合保證金模式(PM)，該字段返回""，風險限額規則失效
	Symbol                 string `json:"symbol"`                 //合約名称
	Side                   string `json:"side"`                   //持倉方向，Buy：多头；Sell：空头. 經典帳戶的單向模式下和統一帳戶的反向合約: 空倉時返回None. 統一帳戶(正向合約): 單向或對沖模式空的仓位返回空字符串
	Size                   string `json:"size"`                   //當前倉位的合约數量
	AvgPrice               string `json:"avgPrice"`               //當前倉位的平均入場價格 對於8小時結算的USDC合約倉位, 該字段表示的是平均開倉價格, 不隨著結算而改變
	PositionValue          string `json:"positionValue"`          //仓位的價值
	TradeMode              int    `json:"tradeMode"`              //交易模式。 統一帳戶 (反向合約) & 經典帳戶: 0: 全倉, 1: 逐倉 統一帳戶: 廢棄, 總是 0
	AutoAddMargin          int    `json:"autoAddMargin"`          //是否自動追加保證金. 0: 否, 1: 是. 僅當統一帳戶(除反向合約)開啟了帳戶維度的逐倉保證金模式, 該字段才有意義
	PositionStatus         string `json:"positionStatus"`         //倉位状态. Normal,Liq, Adl
	Leverage               string `json:"leverage"`               //當前倉位的槓桿，仅适用于合约. 注意：若賬戶為組合保證金模式(PM)，該字段返回空字符串，槓桿規則失效
	MarkPrice              string `json:"markPrice"`              //symbol 的最新標記價格
	LiqPrice               string `json:"liqPrice"`               //倉位強平價格， UTA(反向合約) & 普通账户 & UTA(開啟逐倉保證金模式)：是逐倉和全倉持仓的真實價格, 當強平價 <= minPrice或者 強平價 >= maxPrice, 則為""。 統一帳戶(全倉保證金)：是全倉持仓的预估价格（因为统一帳戶模式是按照帳戶維度控制风险率), 當強平價 <= minPrice或者 強平價 >= maxPrice, 則為"" 但是對於組合保證金模式，此字段為空，不會提供強平價格
	BustPrice              string `json:"bustPrice"`              //倉位破產價格. 統一保證金模式返回"", 無倉位破產價格 (不包括統一帳戶下的反向交易)
	PositionIM             string `json:"positionIM"`             //倉位起始保證金. 組合保證金模式(PM)下, 該字段返回為空字符串
	PositionMM             string `json:"positionMM"`             //倉位維持保證金. 組合保證金模式(PM)下, 該字段返回為空字符串
	TpslMode               string `json:"tpslMode"`               //該字段廢棄, 無意義, 總是返回"Full". 期權總是返回""
	PositionBalance        string `json:"positionBalance"`        //倉位保證金 統一帳戶(linear): 僅在逐倉保證金模式下有意義
	TakeProfit             string `json:"takeProfit"`             //止盈價格
	StopLoss               string `json:"stopLoss"`               //止損價格
	TrailingStop           string `json:"trailingStop"`           //追蹤止損（與當前價格的距離）
	SessionAvgPrice        string `json:"sessionAvgPrice"`        //USDC合約平均持倉價格, 會隨著8小時結算而變動
	Delta                  string `json:"delta"`                  //Delta, 期權的獨有字段
	Gamma                  string `json:"gamma"`                  //Gamma, 期權的獨有字段
	Vega                   string `json:"vega"`                   //Vega, 期權的獨有字段
	Theta                  string `json:"theta"`                  //Theta, 期權的獨有字段
	UnrealisedPnl          string `json:"unrealisedPnl"`          //未结盈亏
	CurRealisedPnl         string `json:"curRealisedPnl"`         //當前持倉的已結盈虧
	CumRealisedPnl         string `json:"cumRealisedPnl"`         //累计已结盈亏 期貨: 是從第一次開始有持倉加總的已結盈虧 期權: 總是"", 無意義
	AdlRankIndicator       int    `json:"adlRankIndicator"`       //自動減倉燈.
	IsReduceOnly           bool   `json:"isReduceOnly"`           //僅當Bybit需要降低某個Symbol的風險限額時有用 true: 僅允許減倉操作. 您可以考慮一系列的方式, 比如, 降低risk limit檔位, 或者同檔位修改槓桿或減少倉位, 或者增加保證金, 或者撤單, 這些操作做完後, 可以主動調用確認新的風險限額接口 false(默認): 沒有交易限制, 表示您的倉位在系統調整時處於風險水平之下 僅對逐倉和全倉的期貨倉位有意義
	MmrSysUpdatedTime      string `json:"mmrSysUpdatedTime"`      //僅當Bybit需要降低某個Symbol的風險限額時有用 當isReduceOnly=true: 這個時間戳表示系統強制修改MMR的時間 當isReduceOnly=false: 若不為空, 則表示系統已經完成了MMR調整的時間 僅當系統調整才會賦值, 對於主動的調整, 不會在這裡展示時間戳 默認為"", 但如果曾經這個symbol有過系統降檔的操作, 那麼這裡會顯示上一次操作的時間 僅對逐倉和全倉的期貨倉位有意義
	LeverageSysUpdatedTime string `json:"leverageSysUpdatedTime"` //僅當Bybit需要降低某個Symbol的風險限額時有用 當isReduceOnly=true: 這個時間戳表示系統強制修改槓桿的時間 當isReduceOnly=false: 若不為空, 則表示系統已經完成了槓桿調整的時間 僅當系統調整才會賦值, 對於主動的調整, 不會在這裡展示時間戳 默認為"", 但如果曾經這個symbol有過系統降檔的操作, 那麼這裡會顯示上一次操作的時間 僅對逐倉和全倉的期貨倉位有意義
	CreatedTime            string `json:"createdTime"`            //倉位創建時間
	UpdatedTime            string `json:"updatedTime"`            //倉位數據更新時間
	Seq                    int64  `json:"seq"`                    //序列號, 用於關聯成交和倉位的更新 不同的幣對會存在相同seq, 可以使用seq + symbol來做唯一性識別 如果該幣對從未被交易過, 查詢時則會返回"-1" 對於更新槓桿、更新風險限額等非交易行為, 將會返回上一次成交時更新的seq
}

type PositionSetLeverageRes struct{}

type PositionSwitchIsolatedRes struct{}

type PositionSwitchModeRes struct{}
