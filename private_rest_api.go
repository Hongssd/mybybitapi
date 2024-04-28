package mybybitapi

type PrivateRestAPI int

const (
	//Account
	AccountWalletBalance PrivateRestAPI = iota //查詢錢包餘額
)

var PrivateRestAPIMap = map[PrivateRestAPI]string{
	AccountWalletBalance: "/v5/account/wallet-balance", //GET 查詢錢包餘額
}
