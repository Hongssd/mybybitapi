# mybybitapi
[![Go 1.19.0](https://img.shields.io/badge/Go-1.19.0-brightgreen.svg)](https://github.com/Hongssd/mybybitapi)
[![Contributor Victor](https://img.shields.io/badge/contributor-Victor-blue.svg)](https://github.com/Hongssd/mybybitapi)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://github.com/Hongssd/mybybitapi/LICENSE)

# Table of Contents
- [Require](#Require)
- [Installation](#Installation)
- [Examples](#Examples)

# Require

```go
require (
	github.com/json-iterator/go v1.1.12
	github.com/shopspring/decimal v1.4.0
	github.com/sirupsen/logrus v1.9.3
)
```
# Installation
```shell
go get github.com/Hongssd/mybybitapi
```

# Examples

## Public Instruments Info

```go
client := mybybitapi.NewRestClient("YOUR_API_KEY","YOUR_API_KEY").PublicRestClient()
res, err := client.NewMarketInstrumentsInfo().Category("linear").Do()
if err != nil {
    panic(err)
}
fmt.Println(res)
```

## Public Kline

```go
client := mybybitapi.NewRestClient("YOUR_API_KEY","YOUR_API_KEY").PublicRestClient()
res, err := client.NewMarketKline().Category("spot").Interval("1").Symbol("BTCUSDT").Limit(10).Do()
if err != nil {
    panic(err)
}
fmt.Println(res)
```


## Public OrderBook

```go
client := mybybitapi.NewRestClient("YOUR_API_KEY","YOUR_API_KEY").PublicRestClient()
res, err := client.NewMarketOrderBook().Category("spot").Symbol("BTCUSDT").Limit(20).Do()
if err != nil {
    panic(err)
}
fmt.Println(res)
```

# Private Account Wallet Balance

```go
client := mybybitapi.NewRestClient("YOUR_API_KEY","YOUR_API_KEY").PrivateRestClient()
res, err := client.NewAccountWalletBalance().AccountType("UNIFIED").Do()
if err != nil {
    panic(err)
}
log.Info(res)
```