package mybybitapi

type SpotMarginTradeSetLeverageRes struct {
	RetCode    int      `json:"retCode"`
	RetMsg     string   `json:"retMsg"`
	Result     struct{} `json:"result"`
	RetExtInfo struct{} `json:"retExtInfo"`
	Time       int64    `json:"time"`
}
