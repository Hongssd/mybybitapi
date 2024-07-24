package mybybitapi

type AssetTransferQueryInterTransferListReq struct {
	TransferId *string `json:"transferId,omitempty"` //String	false	string	UUID. 使用創建劃轉時用的UUID
	Coin       *string `json:"coin,omitempty"`       //String	false	string	幣種
	Status     *string `json:"status,omitempty"`     //String	false	string	劃轉狀態
	StartTime  *int    `json:"startTime,omitempty"`  //String	false	integer	開始時間戳 (毫秒) 注意: 實際查詢時是秒級維度生效	當startTime & endTime都不傳入時, API默認返回30天的數據
	EndTime    *int    `json:"endTime,omitempty"`    //String	false	integer	結束時間戳 (毫秒) 注意: 實際查詢時是秒級維度生效
	Limit      *int    `json:"limit,omitempty"`      //String	false	integer	每頁數量限制. [1, 50]. 默認: 20
	Cursor     *string `json:"cursor,omitempty"`     //String	false	string	游標，用於翻頁
}

type AssetTransferQueryInterTransferListAPI struct {
	client *PrivateRestClient
	req    *AssetTransferQueryInterTransferListReq
}

func (api *AssetTransferQueryInterTransferListAPI) TransferId(transferId string) *AssetTransferQueryInterTransferListAPI {
	api.req.TransferId = GetPointer(transferId)
	return api
}

func (api *AssetTransferQueryInterTransferListAPI) Coin(coin string) *AssetTransferQueryInterTransferListAPI {
	api.req.Coin = GetPointer(coin)
	return api
}

func (api *AssetTransferQueryInterTransferListAPI) Status(status string) *AssetTransferQueryInterTransferListAPI {
	api.req.Status = GetPointer(status)
	return api
}

func (api *AssetTransferQueryInterTransferListAPI) StartTime(startTime int) *AssetTransferQueryInterTransferListAPI {
	api.req.StartTime = GetPointer(startTime)
	return api
}

func (api *AssetTransferQueryInterTransferListAPI) EndTime(endTime int) *AssetTransferQueryInterTransferListAPI {
	api.req.EndTime = GetPointer(endTime)
	return api
}

func (api *AssetTransferQueryInterTransferListAPI) Limit(limit int) *AssetTransferQueryInterTransferListAPI {
	api.req.Limit = GetPointer(limit)
	return api
}

func (api *AssetTransferQueryInterTransferListAPI) Cursor(cursor string) *AssetTransferQueryInterTransferListAPI {
	api.req.Cursor = GetPointer(cursor)
	return api
}

type AssetTransferQueryTransferCoinListReq struct {
	FromAccountType *string `json:"fromAccountType"` //String	true	string	劃出帳戶類型
	ToAccountType   *string `json:"toAccountType"`   //String	true	string	劃入帳戶類型
}

type AssetTransferQueryTransferCoinListAPI struct {
	client *PrivateRestClient
	req    *AssetTransferQueryTransferCoinListReq
}

func (api *AssetTransferQueryTransferCoinListAPI) FromAccountType(fromAccountType string) *AssetTransferQueryTransferCoinListAPI {
	api.req.FromAccountType = GetPointer(fromAccountType)
	return api
}

func (api *AssetTransferQueryTransferCoinListAPI) ToAccountType(toAccountType string) *AssetTransferQueryTransferCoinListAPI {
	api.req.ToAccountType = GetPointer(toAccountType)
	return api
}

// AssetTransferInterTransfer:          "/v5/asset/transfer/inter-transfer",            //POST 劃轉 (單帳號內)
type AssetTransferInterTransferReq struct {
	TransferId      *string `json:"transferId"`      //String	true	string	UUID. 請自行手動生成UUID
	Coin            *string `json:"coin"`            //String	true	string	幣種
	Amount          *string `json:"amount"`          //String	true	string	劃入數量
	FromAccountType *string `json:"fromAccountType"` //String	true	string	轉出賬戶類型
	ToAccountType   *string `json:"toAccountType"`   //String	true	string	轉入賬戶類型
}

type AssetTransferInterTransferAPI struct {
	client *PrivateRestClient
	req    *AssetTransferInterTransferReq
}

func (api *AssetTransferInterTransferAPI) TransferId(transferId string) *AssetTransferInterTransferAPI {
	api.req.TransferId = GetPointer(transferId)
	return api
}

func (api *AssetTransferInterTransferAPI) Coin(coin string) *AssetTransferInterTransferAPI {
	api.req.Coin = GetPointer(coin)
	return api
}

func (api *AssetTransferInterTransferAPI) Amount(amount string) *AssetTransferInterTransferAPI {
	api.req.Amount = GetPointer(amount)
	return api
}

func (api *AssetTransferInterTransferAPI) FromAccountType(fromAccountType string) *AssetTransferInterTransferAPI {
	api.req.FromAccountType = GetPointer(fromAccountType)
	return api
}

func (api *AssetTransferInterTransferAPI) ToAccountType(toAccountType string) *AssetTransferInterTransferAPI {
	api.req.ToAccountType = GetPointer(toAccountType)
	return api
}

// AssetTransferQueryAccountCoinsBalance: "/v5/asset/transfer/query-account-coins-balance", //GET 查詢賬戶所有幣種餘額
// AssetTransferQueryAccountCoinBalance:  "/v5/asset/transfer/query-account-coin-balance",  //GET 查詢帳戶單個幣種餘額
// AssetTithdrawWithdrawableAmount:       "/v5/asset/withdraw/withdrawable-amount",         //GET 查詢延遲提幣凍結金額

// memberId	false	string	用戶ID. 當使用母帳號api key查詢子帳戶的幣種餘額時，該字段必傳
// accountType	true	string	賬戶類型
// coin	false	string	幣種類型
// withBonus	false	integer	是否查詢體驗金. 0(默認)：不查詢; 1：查詢

type AssetTransferQueryAccountCoinsBalanceReq struct {
	MemberId    *string `json:"memberId"`    //String	false	string	用戶ID. 當使用母帳號api key查詢子帳戶的幣種餘額時，該字段必傳
	AccountType *string `json:"accountType"` //String	true	string	賬戶類型
	Coin        *string `json:"coin"`        //String	false	string	幣種類型
	WithBonus   *int    `json:"withBonus"`   //String	false	integer	是否查詢體驗金. 0(默認)：不查詢; 1：查詢
}

type AssetTransferQueryAccountCoinsBalanceAPI struct {
	client *PrivateRestClient
	req    *AssetTransferQueryAccountCoinsBalanceReq
}

func (api *AssetTransferQueryAccountCoinsBalanceAPI) MemberId(memberId string) *AssetTransferQueryAccountCoinsBalanceAPI {
	api.req.MemberId = GetPointer(memberId)
	return api
}

func (api *AssetTransferQueryAccountCoinsBalanceAPI) AccountType(accountType string) *AssetTransferQueryAccountCoinsBalanceAPI {
	api.req.AccountType = GetPointer(accountType)
	return api
}

func (api *AssetTransferQueryAccountCoinsBalanceAPI) Coin(coin string) *AssetTransferQueryAccountCoinsBalanceAPI {
	api.req.Coin = GetPointer(coin)
	return api
}

func (api *AssetTransferQueryAccountCoinsBalanceAPI) WithBonus(withBonus int) *AssetTransferQueryAccountCoinsBalanceAPI {
	api.req.WithBonus = GetPointer(withBonus)
	return api
}

// memberId	false	string	用戶Id. 當查詢子帳號的餘額時，該字段必傳
// toMemberId	false	string	劃入帳戶UID. 當查詢不同uid間劃轉時, 該字段必傳
// accountType	true	string	帳戶類型
// toAccountType	false	string	劃入帳戶類型. 當查詢不同帳戶類型間的劃轉時, 該字段必傳
// coin	true	string	幣種
// withBonus	false	integer	是否查詢體驗金. 0(默認): 不查詢,1: 查詢.
// withTransferSafeAmount	false	integer	是否查詢延遲提幣安全限額  // 0(默認)：否, 1：是// 什麼是延遲提幣?
// withLtvTransferSafeAmount	false	integer	特別用於機構借貸用戶, 可以查詢風險水平內的可劃轉餘額0(default)：false, 1：true  此時toAccountType字段必傳
type AssetTransferQueryAccountCoinBalanceReq struct {
	MemberId                  *string `json:"memberId"`                  //String	false	string	用戶ID. 當使用母帳號api key查詢子帳戶的幣種餘額時，該字段必傳
	ToMemberId                *string `json:"toMemberId"`                //String	false	string	劃入帳戶UID. 當查詢不同uid間劃轉時, 該字段必傳
	AccountType               *string `json:"accountType"`               //String	true	string	帳戶類型
	Coin                      *string `json:"coin"`                      //String	true	string	幣種
	WithBonus                 *int    `json:"withBonus"`                 //String	false	integer	是否查詢體驗金. 0(默認): 不查詢,1: 查詢.
	WithTransferSafeAmount    *int    `json:"withTransferSafeAmount"`    //String	false	integer	是否查詢延遲提幣安全限額  // 0(默認)：否, 1：是// 什麼是延遲提幣?
	WithLtvTransferSafeAmount *int    `json:"withLtvTransferSafeAmount"` //String	false	integer	特別用於機構借貸用戶, 可以查詢風險水平內的可劃轉餘額0(default)：false, 1：true  此時toAccountType字段必傳
}

type AssetTransferQueryAccountCoinBalanceAPI struct {
	client *PrivateRestClient
	req    *AssetTransferQueryAccountCoinBalanceReq
}

func (api *AssetTransferQueryAccountCoinBalanceAPI) MemberId(memberId string) *AssetTransferQueryAccountCoinBalanceAPI {
	api.req.MemberId = GetPointer(memberId)
	return api
}

func (api *AssetTransferQueryAccountCoinBalanceAPI) ToMemberId(toMemberId string) *AssetTransferQueryAccountCoinBalanceAPI {
	api.req.ToMemberId = GetPointer(toMemberId)
	return api
}

func (api *AssetTransferQueryAccountCoinBalanceAPI) AccountType(accountType string) *AssetTransferQueryAccountCoinBalanceAPI {
	api.req.AccountType = GetPointer(accountType)
	return api
}

func (api *AssetTransferQueryAccountCoinBalanceAPI) Coin(coin string) *AssetTransferQueryAccountCoinBalanceAPI {
	api.req.Coin = GetPointer(coin)
	return api
}

func (api *AssetTransferQueryAccountCoinBalanceAPI) WithBonus(withBonus int) *AssetTransferQueryAccountCoinBalanceAPI {
	api.req.WithBonus = GetPointer(withBonus)
	return api
}

func (api *AssetTransferQueryAccountCoinBalanceAPI) WithTransferSafeAmount(withTransferSafeAmount int) *AssetTransferQueryAccountCoinBalanceAPI {
	api.req.WithTransferSafeAmount = GetPointer(withTransferSafeAmount)
	return api
}

func (api *AssetTransferQueryAccountCoinBalanceAPI) WithLtvTransferSafeAmount(withLtvTransferSafeAmount int) *AssetTransferQueryAccountCoinBalanceAPI {
	api.req.WithLtvTransferSafeAmount = GetPointer(withLtvTransferSafeAmount)
	return api
}

// coin	true	string	幣種敏誠
type AssetTithdrawWithdrawableAmountReq struct {
	Coin *string `json:"coin"` //String	true	string	幣種敏誠
}

type AssetTithdrawWithdrawableAmountAPI struct {
	client *PrivateRestClient
	req    *AssetTithdrawWithdrawableAmountReq
}

func (api *AssetTithdrawWithdrawableAmountAPI) Coin(coin string) *AssetTithdrawWithdrawableAmountAPI {
	api.req.Coin = GetPointer(coin)
	return api
}
