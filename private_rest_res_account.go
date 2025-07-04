package mybybitapi

type AccountInfoRes struct {
	MarginMode          string `json:"marginMode"`
	UpdatedTime         string `json:"updatedTime"`
	UnifiedMarginStatus int    `json:"unifiedMarginStatus"`
	DcpStatus           string `json:"dcpStatus"`
	TimeWindow          int    `json:"timeWindow"`
	SmpGroup            int    `json:"smpGroup"`
	IsMasterTrader      bool   `json:"isMasterTrader"`
	SpotHedgingStatus   string `json:"spotHedgingStatus"`
}

type AccountWalletBalanceRes struct {
	List []AccountWalletBalanceResRow `json:"list"`
}

type AccountWalletBalanceResRow struct {
	AccountType            string                           `json:"accountType"`            //帳戶類型
	AccountLTV             string                           `json:"accountLTV"`             //字段廢棄
	AccountIMRate          string                           `json:"accountIMRate"`          //帳戶初始保證金率  Account Total Initial Margin Base Coin / Account Margin Balance Base Coin. 非統一保證金模式和統一帳戶(反向)以及統一帳戶(逐倉模式)，該字段返回為空。
	AccountMMRate          string                           `json:"accountMMRate"`          //帳戶維持保證金率 Account Total Maintenance Margin Base Coin / Account Margin Balance Base Coin。非統一保證金模式和統一帳戶(反向)以及統一帳戶(逐倉模式)，該字段返回為空。
	TotalEquity            string                           `json:"totalEquity"`            //總凈值為賬戶中每個幣種資產凈值的法幣估值之和。。非統一保證金模式以及統一帳戶(反向)下，該字段返回為空。
	TotalWalletBalance     string                           `json:"totalWalletBalance"`     //賬戶維度換算成usd的產錢包餘額：∑ Asset Wallet Balance By USD value of each asset。非統一保證金模式和統一帳戶(反向)以及統一帳戶(逐倉模式)，該字段返回為空。
	TotalMarginBalance     string                           `json:"totalMarginBalance"`     //賬戶維度換算成usd的保證金餘額：totalWalletBalance + totalPerpUPL。非統一保證金模式和統一帳戶(反向)以及統一帳戶(逐倉模式)，該字段返回為空。
	TotalAvailableBalance  string                           `json:"totalAvailableBalance"`  //賬戶維度換算成usd的可用餘額：RM：totalMarginBalance - totalInitialMargin。非統一保證金模式和統一帳戶(反向)以及統一帳戶(逐倉模式)，該字段返回為空。
	TotalPerpUPL           string                           `json:"totalPerpUPL"`           //賬戶維度換算成usd的永續和USDC交割合約的浮動盈虧：∑ Each perp and USDC Futures upl by base coin。非統一保證金模式以及統一帳戶(反向)下，該字段返回為空。
	TotalInitialMargin     string                           `json:"totalInitialMargin"`     //賬戶維度換算成usd的總初始保證金：∑ Asset Total Initial Margin Base Coin。非統一保證金模式和統一帳戶(反向)以及統一帳戶(逐倉模式)，該字段返回為空。
	TotalMaintenanceMargin string                           `json:"totalMaintenanceMargin"` //賬戶維度換算成usd的總維持保證金: ∑ Asset Total Maintenance Margin Base Coin。非統一保證金模式和統一帳戶(反向)以及統一帳戶(逐倉模式)，該字段返回為空。
	Coin                   []AccountWalletBalanceResRowCoin `json:"coin"`                   //幣種列表

}

type AccountWalletBalanceResRowCoin struct {
	Coin                string `json:"coin"`                //幣種名稱，例如 BTC，ETH，USDT，USDC
	Equity              string `json:"equity"`              //當前幣種的資產淨值
	UsdValue            string `json:"usdValue"`            //當前幣種折算成 usd 的價值
	WalletBalance       string `json:"walletBalance"`       //當前幣種的錢包餘額
	Free                string `json:"free"`                //經典帳戶現貨錢包的可用餘額. 經典帳戶現貨錢包的獨有字段
	Locked              string `json:"locked"`              //現貨掛單凍結金額
	SpotHedgingQty      string `json:"spotHedgingQty"`      //用於組合保證金(PM)現貨對衝的數量, 截斷至8為小數, 默認為0 統一帳戶的獨有字段
	BorrowAmount        string `json:"borrowAmount"`        //當前幣種的已用借貸額度
	AvailableToWithdraw string `json:"availableToWithdraw"` //當前幣種的可劃轉提現金額
	AccruedInterest     string `json:"accruedInterest"`     //當前幣種的預計要在下一個利息週期收取的利息金額
	TotalOrderIM        string `json:"totalOrderIM"`        //以當前幣種結算的訂單委託預佔用保證金. 組合保證金模式下，該字段返回空字符串
	TotalPositionIM     string `json:"totalPositionIM"`     //以當前幣種結算的所有倉位起始保證金求和 + 所有倉位的預佔用平倉手續費. 組合保證金模式下，該字段返回空字符串
	TotalPositionMM     string `json:"totalPositionMM"`     //以當前幣種結算的所有倉位維持保證金求和. 組合保證金模式下，該字段返回空字符串
	UnrealisedPnl       string `json:"unrealisedPnl"`       //以當前幣種結算的所有倉位的未結盈虧之和
	CumRealisedPnl      string `json:"cumRealisedPnl"`      //以當前幣種結算的所有倉位的累計已結盈虧之和
	Bonus               string `json:"bonus"`               //體驗金. UNIFIED帳戶的獨有字段
	MarginCollateral    bool   `json:"marginCollateral"`    //是否可作為保證金抵押幣種(平台維度), true: 是. false: 否
	CollateralSwitch    bool   `json:"collateralSwitch"`    //用戶是否開啟保證金幣種抵押(用戶維度), true: 是. false: 否
}

type AccountFeeRateRes struct {
	Category string                 `json:"category"` //產品類型. spot, option. 期貨不返回該字段
	List     []AccountFeeRateResRow `json:"list"`
}

type AccountFeeRateResRow struct {
	Symbol       string `json:"symbol"`       //合約名稱. 期權總是為""
	BaseCoin     string `json:"baseCoin"`     //交易幣種. SOL, BTC, ETH 期貨不返回該字段 現貨總是返回""
	TakerFeeRate string `json:"takerFeeRate"` //吃單手續費率
	MakerFeeRate string `json:"makerFeeRate"` //掛單手續費率
}

type AccountUpgradeToUtaRes struct {
	UnifiedUpdateStatus string `json:"unifiedUpdateStatus"`
	UnifiedUpdateMsg    struct {
		Msg []string `json:"msg"`
	} `json:"unifiedUpdateMsg"`
}

type AccountSetMarginModeRes struct {
	Reasons []struct {
		ReasonCode string `json:"reasonCode"`
		ReasonMsg  string `json:"reasonMsg"`
	} `json:"reasons"`
}

// availableWithdrawal	string	請求中第一個幣種的可劃轉餘額
// availableWithdrawalMap	Object	每個請求幣種的可劃轉餘額的對象。在映射中，鍵是請求的幣種，值是相應的金額(字符串)
// 例如, "availableWithdrawalMap":{"BTC":"4.54549050","SOL":"33.16713007","XRP":"10805.54548970","ETH":"17.76451865"}
type AccountWithdrawalRes struct {
	AvailableWithdrawal    string            `json:"availableWithdrawal"`
	AvailableWithdrawalMap map[string]string `json:"availableWithdrawalMap"`
}

type AccountSetCollateralSwitchRes struct{}

type AccountSetCollateralSwitchBatchResRow struct {
	Coin             string `json:"coin"`             //幣種名稱
	CollateralSwitch string `json:"collateralSwitch"` //ON: 開啟抵押, OFF: 關閉抵押
}
type AccountSetCollateralSwitchBatchRes struct {
	List []AccountSetCollateralSwitchBatchResRow `json:"list"`
}

type AccountCollateralInfoRes struct {
	List []AccountCollateralInfoResRow `json:"list"`
}

type AccountCollateralInfoResRow struct {
	Currency            string `json:"currency"`            //目前所有抵押品的資產幣種
	HourlyBorrowRate    string `json:"hourlyBorrowRate"`    //每小時藉款利率
	MaxBorrowingAmount  string `json:"maxBorrowingAmount"`  //最大可藉貸額度. 該值由母子帳號共享
	FreeBorrowingLimit  string `json:"freeBorrowingLimit"`  //免息借款額上限
	FreeBorrowAmount    string `json:"freeBorrowAmount"`    //借款總額中免息部分的借款金額
	BorrowAmount        string `json:"borrowAmount"`        //已用借貸額度
	OtherBorrowAmount   string `json:"otherBorrowAmount"`   //其他帳戶的已借貸總額(同一個母帳戶下)
	AvailableToBorrow   string `json:"availableToBorrow"`   //用戶剩餘可藉額度. 該值由母子帳號共享
	Borrowable          bool   `json:"borrowable"`          //是否是可藉貸的幣種, true: 是. false: 否
	FreeBorrowingAmount string `json:"freeBorrowingAmount"` //廢棄字段, 總是返回空字符串, 請參考freeBorrowingLimit
	BorrowUsageRate     string `json:"borrowUsageRate"`     //借貸資金使用率: 母子帳戶加起來的borrowAmount/maxBorrowingAmount. 這是一個真實值, 0.5则表示50%
	MarginCollateral    bool   `json:"marginCollateral"`    //是否可作為保證金抵押幣種(平台維度), true: 是. false: 否
	CollateralSwitch    bool   `json:"collateralSwitch"`    //用戶是否開啟保證金幣種抵押(用戶維度), true: 是. false: 否
	CollateralRatio     string `json:"collateralRatio"`     //由於新的階梯價值率邏輯, 該字段從2025年2月19日開始不再準確。請使用查詢階梯價值率
}
