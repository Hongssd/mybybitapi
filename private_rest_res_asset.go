package mybybitapi

// AssetTransferQueryInterTransferList:   PrivateRest接口   //GET 查詢劃轉紀錄 (單帳號內)
type AssetTransferQueryInterTransferListRes struct {
	List           []AssetTransferQueryInterTransferListResRow `json:"list"`
	NextPageCursor string                                      `json:"nextPageCursor"` //string	游標，用於翻頁
}

type AssetTransferQueryInterTransferListResRow struct {
	TransferId      string `json:"transferId"`      //劃轉Id
	Coin            string `json:"coin"`            //劃轉幣種
	Amount          string `json:"amount"`          //劃轉金額
	FromAccountType string `json:"fromAccountType"` //劃出賬戶類型
	ToAccountType   string `json:"toAccountType"`   //劃入賬戶類型
	Timestamp       string `json:"timestamp"`       //劃轉創建時間戳 (毫秒)
	Status          string `json:"status"`          //劃轉狀態
}

// AssetTransferQueryTransferCoinList:      PrivateRest接口   //GET 帳戶類型間可劃轉的幣種
type AssetTransferQueryTransferCoinListRes struct {
	List []string `json:"list"` //array	幣種數組
}

// AssetTransferInterTransfer:                  PrivateRest接口  //POST 劃轉 (單帳號內)
type AssetTransferInterTransferRes struct {
	TransferId string `json:"transferId"` //string	UUID
	Status     string `json:"status"`     //string	劃轉狀態  STATUS_UNKNOWN  SUCCESS  PENDING   FAILED
}

// AssetTransferQueryAccountCoinsBalance:  PrivateRest接口    //GET 查詢賬戶所有幣種餘額
type AssetTransferQueryAccountCoinsBalanceRes struct {
	AccountType string `json:"accountType"` //string	賬戶類型
	MemberId    string `json:"memberId"`    //string	用戶ID
	Balance     []struct {
		Coin            string `json:"coin"`            //string	幣種類型
		WalletBalance   string `json:"walletBalance"`   //string	錢包余額
		TransferBalance string `json:"transferBalance"` //string	可划余額
		Bonus           string `json:"bonus"`           //string	体验金
	} `json:"balance"`
}

// AssetTransferQueryAccountCoinBalance:    PrivateRest接口    //GET 查詢帳戶單個幣種餘額
type AssetTransferQueryAccountCoinBalanceRes struct {
	AccountType string `json:"accountType"` //string	賬戶類型
	BizType     int    `json:"bizType"`     //integer	帳戶業務子類型
	AccountId   string `json:"accountId"`   //string	賬戶ID
	MemberId    string `json:"memberId"`    //string	用戶ID
	Balance     struct {
		Coin                  string `json:"coin"`                  //string	幣種類型
		WalletBalance         string `json:"walletBalance"`         //string	錢包余額
		TransferBalance       string `json:"transferBalance"`       //string	可划余額
		Bonus                 string `json:"bonus"`                 //string	可用金額中包含的体验金
		TransferSafeAmount    string `json:"transferSafeAmount"`    //string	可劃轉的安全限額. 若不查詢，則返回""
		LtvTransferSafeAmount string `json:"ltvTransferSafeAmount"` //string	機構借貸用戶的可劃轉餘額. 若不查詢，則返回""
	} `json:"balance"`
}

// AssetTithdrawWithdrawableAmount:              PrivateRest接口  //GET 查詢延遲提幣凍結金額
type AssetTithdrawWithdrawableAmountRes struct {
	LimitAmountUsd     string `json:"limitAmountUsd"` //string	延遲提幣凍結金額 (USD)
	WithdrawableAmount struct {
		SPOT struct { //Object	現貨錢包, 若該錢包被移除, 則不會返回該對象
			Coin               string `json:"coin"`               //string	幣種名稱
			WithdrawableAmount string `json:"withdrawableAmount"` //string	可提現金額
			AvailableBalance   string `json:"availableBalance"`   //string	可用餘額
		} `json:"SPOT"`
		FUND struct { //Object	資金錢包
			Coin               string `json:"coin"`               //string	幣種名稱
			WithdrawableAmount string `json:"withdrawableAmount"` //string	可提現金額
			AvailableBalance   string `json:"availableBalance"`   //string	可用餘額
		} `json:"FUND"`
	} `json:"withdrawableAmount"`
}
