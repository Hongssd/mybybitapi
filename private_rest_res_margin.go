package mybybitapi

type SpotMarginTradeSetLeverageRes struct {
	RetCode    int      `json:"retCode"`
	RetMsg     string   `json:"retMsg"`
	Result     struct{} `json:"result"`
	RetExtInfo struct{} `json:"retExtInfo"`
	Time       int64    `json:"time"`
}

// spotLeverage	string	槓桿倍數. 如果處於關閉狀態的話, 則返回 ""
// spotMarginMode	string	開關狀態. 1: 開啟, 0: 關閉
// effectiveLeverage	string	實際借貸槓桿倍數。 精度保留2位小數，向下截取
type SpotMarginTradeStateRes struct {
	SpotLeverage      string `json:"spotLeverage"`      // 槓桿倍數. 如果處於關閉狀態的話, 則返回 ""
	SpotMarginMode    string `json:"spotMarginMode"`    // 開關狀態. 1: 開啟, 0: 關閉
	EffectiveLeverage string `json:"effectiveLeverage"` // 實際借貸槓桿倍數。 精度保留2位小數，向下截取
}
