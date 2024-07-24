package mybybitapi

//AssetTransferQueryInterTransferList   //查詢劃轉紀錄 (單帳號內)

type AssetTransferQueryInterTransferListRes struct {
	List           []AssetTransferQueryInterTransferListResRow `json:"list"`
	NextPageCursor string                                      `json:"nextPageCursor"`
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

//AssetTransferQueryTransferCoinList    //帳戶類型間可劃轉的幣種

type AssetTransferQueryTransferCoinListRes struct {
	List []string `json:"list"`
}

// AssetTransferInterTransfer            //劃轉 (單帳號內)
// transferId	string	UUID
// status	string	劃轉狀態
// STATUS_UNKNOWN
// SUCCESS
// PENDING
// FAILED
type AssetTransferInterTransferRes struct {
	TransferId string `json:"transferId"`
	Status     string `json:"status"`
}



// AssetTransferQueryAccountCoinsBalance //查詢賬戶所有幣種余額
// accountType	string	賬戶類型
// memberId	string	用戶ID
// balance	Object
// > coin	string	幣種類型
// > walletBalance	string	錢包余額
// > transferBalance	string	可划余額
// > bonus	string	体验金
type AssetTransferQueryAccountCoinsBalanceRes struct {
	AccountType string `json:"accountType"`
	MemberId    string `json:"memberId"`
	Balance     []struct {
		Coin            string `json:"coin"`
		WalletBalance   string `json:"walletBalance"`
		TransferBalance string `json:"transferBalance"`
		Bonus           string `json:"bonus"`
	} `json:"balance"`
}

// AssetTransferQueryAccountCoinBalance  //查詢帳戶單個幣種餘額
// accountType	string	賬戶類型
// bizType	integer	帳戶業務子類型
// accountId	string	賬戶ID
// memberId	string	用戶ID
// balance	Object
// > coin	string	幣種類型
// > walletBalance	string	錢包余額
// > transferBalance	string	可划余額
// > bonus	string	可用金額中包含的体验金
// > transferSafeAmount	string	可劃轉的安全限額. 若不查詢，則返回""
// > ltvTransferSafeAmount	string	機構借貸用戶的可劃轉餘額. 若不查詢，則返回""
type AssetTransferQueryAccountCoinBalanceRes struct {
	AccountType string `json:"accountType"`
	BizType     int    `json:"bizType"`
	AccountId   string `json:"accountId"`
	MemberId    string `json:"memberId"`
	Balance     struct {
		Coin                  string `json:"coin"`
		WalletBalance         string `json:"walletBalance"`
		TransferBalance       string `json:"transferBalance"`
		Bonus                 string `json:"bonus"`
		TransferSafeAmount    string `json:"transferSafeAmount"`
		LtvTransferSafeAmount string `json:"ltvTransferSafeAmount"`
	} `json:"balance"`
}


// AssetTithdrawWithdrawableAmount       //查詢延遲提幣凍結金額
// limitAmountUsd	string	延遲提幣凍結金額 (USD)
// withdrawableAmount	Object
// > SPOT	Object	現貨錢包, 若該錢包被移除, 則不會返回該對象
// >> coin	string	幣種名稱
// >> withdrawableAmount	string	可提現金額
// >> availableBalance	string	可用餘額
// > FUND	Object	資金錢包
// >> coin	string	幣種名稱
// >> withdrawableAmount	string	可提現金額
// >> availableBalance	string	可用餘額

type AssetTithdrawWithdrawableAmountRes struct {
	LimitAmountUsd     string `json:"limitAmountUsd"`
	WithdrawableAmount struct {
		SPOT struct {
			Coin               string `json:"coin"`
			WithdrawableAmount string `json:"withdrawableAmount"`
			AvailableBalance   string `json:"availableBalance"`
		} `json:"SPOT"`
		FUND struct {
			Coin               string `json:"coin"`
			WithdrawableAmount string `json:"withdrawableAmount"`
			AvailableBalance   string `json:"availableBalance"`
		} `json:"FUND"`
	} `json:"withdrawableAmount"`
}
