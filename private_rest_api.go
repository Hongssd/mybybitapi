package mybybitapi

type PrivateRestAPI int

const (
	//Account
	AccountInfo                     PrivateRestAPI = iota //查詢帳戶信息
	AccountWalletBalance                                  //查詢錢包餘額
	AccountFeeRate                                        //查詢手續費率
	AccountUpgradeToUta                                   //升級至UTA Pro
	AccountSetMarginMode                                  //設置保證金模式(帳戶)
	AccountWithdrawal                                     //查詢可劃轉餘額(统一账户)
	AccountSetCollateralSwitch                            //設置抵押品幣種
	AccountSetCollateralSwitchBatch                       //批量設置抵押品幣種
	AccountCollateralInfo                                 //查詢抵押品幣種

	//Position
	PositionList           //查詢持倉 (實時)
	PositionSetLeverage    //設置槓桿
	PositionSwitchIsolated //切換全倉/逐倉保證金(交易對)
	PositionSwitchMode     //切換持倉模式
	//Order
	OrderCreate        //創建委託單
	OrderCreateBatch   //批量創建委託單
	OrderAmend         //修改委託單
	OrderAmendBatch    //批量修改委託單
	OrderCancel        //撤銷委託單
	OrderCancelBatch   //批量撤銷委託單
	OrderCancelAll     //撤銷所有訂單
	OrderRealtime      //查詢實時委託單
	OrderHistory       //查詢歷史訂單 (2年)
	OrderExecutionList //查詢成交紀錄 (2年)
	//asset
	AssetTransferQueryInterTransferList   //查詢劃轉紀錄 (單帳號內)
	AssetTransferQueryTransferCoinList    //帳戶類型間可劃轉的幣種
	AssetTransferInterTransfer            //劃轉 (單帳號內)
	AssetTransferQueryAccountCoinsBalance //查詢賬戶所有幣種余額
	AssetTransferQueryAccountCoinBalance  //查詢帳戶單個幣種餘額
	AssetTithdrawWithdrawableAmount       //查詢延遲提幣凍結金額

	//margin
	SpotMarginTradeSetLeverage //全倉槓桿設置用戶最大槓桿倍數
	SpotMarginTradeState       //查詢統一帳戶下槓桿交易的開關狀態和槓桿倍數
)

var PrivateRestAPIMap = map[PrivateRestAPI]string{
	AccountInfo:                     "/v5/account/info",                        // GET 查詢帳戶信息
	AccountWalletBalance:            "/v5/account/wallet-balance",              //GET 查詢錢包餘額
	AccountFeeRate:                  "/v5/account/fee-rate",                    //GET 查詢手續費率
	AccountUpgradeToUta:             "/v5/account/upgrade-to-uta",              //POST 升級至UTA Pro
	AccountSetMarginMode:            "/v5/account/set-margin-mode",             //POST 設置保證金模式(帳戶)
	AccountWithdrawal:               "/v5/account/withdrawal",                  //GET 查詢可劃轉餘額(统一账户)
	AccountSetCollateralSwitch:      "/v5/account/set-collateral-switch",       //POST 設置抵押品幣種
	AccountSetCollateralSwitchBatch: "/v5/account/set-collateral-switch-batch", //POST 批量設置抵押品幣種
	AccountCollateralInfo:           "/v5/account/collateral-info",             //GET 查詢抵押品幣種

	PositionList:           "/v5/position/list",            //GET 查詢持倉 (實時)
	PositionSetLeverage:    "/v5/position/set-leverage",    //POST 設置槓桿（統一帳戶覆蓋範圍: USDT永續 / USDC永續 / USDC交割 / 反向合約）
	PositionSwitchIsolated: "/v5/position/switch-isolated", //POST 切換全倉/逐倉保證金(交易對)
	PositionSwitchMode:     "/v5/position/switch-mode",     //POST 切換持倉模式

	OrderCreate:        "/v5/order/create",       //POST 創建委託單
	OrderCreateBatch:   "/v5/order/create-batch", //POST 批量創建委託單
	OrderAmend:         "/v5/order/amend",        //POST 修改委託單
	OrderAmendBatch:    "/v5/order/amend-batch",  //POST 批量修改委託單
	OrderCancel:        "/v5/order/cancel",       //POST 撤銷委託單
	OrderCancelBatch:   "/v5/order/cancel-batch", //POST 批量撤銷委託單
	OrderCancelAll:     "/v5/order/cancel-all",   //POST 撤銷所有訂單
	OrderRealtime:      "/v5/order/realtime",     //GET 查詢實時委託單
	OrderHistory:       "/v5/order/history",      //GET 查詢歷史訂單 (2年)
	OrderExecutionList: "/v5/execution/list",     //GET 查詢成交紀錄 (2年)

	AssetTransferQueryInterTransferList:   "/v5/asset/transfer/query-inter-transfer-list",   //GET 查詢劃轉紀錄 (單帳號內)
	AssetTransferQueryTransferCoinList:    "/v5/asset/transfer/query-transfer-coin-list",    //GET 帳戶類型間可劃轉的幣種
	AssetTransferInterTransfer:            "/v5/asset/transfer/inter-transfer",              //POST 劃轉 (單帳號內)
	AssetTransferQueryAccountCoinsBalance: "/v5/asset/transfer/query-account-coins-balance", //GET 查詢賬戶所有幣種餘額
	AssetTransferQueryAccountCoinBalance:  "/v5/asset/transfer/query-account-coin-balance",  //GET 查詢帳戶單個幣種餘額
	AssetTithdrawWithdrawableAmount:       "/v5/asset/withdraw/withdrawable-amount",         //GET 查詢延遲提幣凍結金額

	// margin
	SpotMarginTradeSetLeverage: "/v5/spot-margin-trade/set-leverage", // POST 全倉槓桿設置用戶最大槓桿倍數, 支持區間 [2, 10]
	SpotMarginTradeState:       "/v5/spot-margin-trade/state",        // 查詢統一帳戶下槓桿交易的開關狀態和槓桿倍數
}
