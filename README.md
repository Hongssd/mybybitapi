# mybybitapi
[![Go 1.19.0](https://img.shields.io/badge/Go-1.19.0-brightgreen.svg)](https://github.com/Hongssd/mybybitapi)
[![Contributor Victor](https://img.shields.io/badge/contributor-Victor-blue.svg)](https://github.com/Hongssd/mybybitapi)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://github.com/Hongssd/mybybitapi/blob/master/LICENSE)

# Table of Contents
- [Require](#Require)
- [Installation](#Installation)
- [Public Rest Examples](#Public-Rest-Examples)
- [Private Rest Examples](#Private-Rest-Examples)
- [Public Websocket Examples](#Public-Websocket-Examples)
- [Private Websocket Examples](#Private-Websocket-Examples)
- [Trade Websocket Examples](#Trade-Websocket-Examples)
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


# Public Rest Examples
## Kline
```go
client := mybybitapi.NewRestClient("YOUR_API_KEY","YOUR_API_KEY").PublicRestClient()
res, err := client.NewMarketKline().Category("spot").Interval("1").Symbol("BTCUSDT").Limit(10).Do()
if err != nil {
    panic(err)
}
fmt.Println(res)
```
## OrderBook
```go
client := mybybitapi.NewRestClient("YOUR_API_KEY","YOUR_API_KEY").PublicRestClient()
res, err := client.NewMarketOrderBook().Category("spot").Symbol("BTCUSDT").Limit(20).Do()
if err != nil {
    panic(err)
}
fmt.Println(res)
```

# Private Rest Examples

## Wallet Balance
```go
client := mybybitapi.NewRestClient("YOUR_API_KEY","YOUR_API_KEY").PrivateRestClient()
res, err := client.NewAccountWalletBalance().AccountType("UNIFIED").Do()
if err != nil {
    panic(err)
}
fmt.Println(res)
```

## Order Create
```go
client := mybybitapi.NewRestClient("YOUR_API_KEY","YOUR_API_KEY").PrivateRestClient()
res, err := client.NewOrderCreate().
	Category("linear").Symbol("BTCUSDT").Side("Buy").OrderType("Limit").Qty("0.1").Price("10000").
	Do()
if err != nil {
    panic(err)
}
fmt.Println(res)
```

## Batch Order Create
```go
client := mybybitapi.NewRestClient("YOUR_API_KEY","YOUR_API_KEY").PrivateRestClient()
res, err := client.NewOrderCreateBatch().
        AddNewOrderCreateReq(client.NewOrderCreate().
			Category("linear").Symbol("BTCUSDT").Side("Buy").OrderType("Limit").Qty("0.1").Price("10000")).
        AddNewOrderCreateReq(client.NewOrderCreate().
			Category("linear").Symbol("BTCUSDT").Side("Buy").OrderType("Limit").Qty("0.1").Price("11000")).
        Do()
if err != nil {
    panic(err)
}
fmt.Println(res)
```

# Public Websocket Examples

## Subscribe Kline
```go
wsclient := mybybitapi.NewPublicSpotWsStreamClient()
//Open Websocket Connect
err := wsclient.OpenConn()
if err != nil {
    panic(err)
}
//Subscribe Kline
sub, err := wsclient.SubscribeKlineMultiple([]string{"BTCUSDT", "ETHUSDT"}, []string{"1", "5", "30"})
if err != nil {
    panic(err)
}


//Listen for klines 
for {
    select {
    case err := <-sub.ErrChan():
        panic(err)
    case result := <-sub.ResultChan():
        fmt.Println(result)
    case <-sub.CloseChan():
        fmt.Println("subscribe closed")
    }
}
```


# Private Websocket Examples

## Subscribe Order
```go
client := mybybitapi.NewRestClient("YOUR_API_KEY","YOUR_API_KEY")
wsclient := mybybitapi.NewPrivateWsStreamClient()

//Open Websocket Connect
err := wsclient.OpenConn()
if err != nil {
    panic(err)
}

//Auth To Bybit
err = wsclient.Auth(client)
if err != nil {
    panic(err)
}


//Subscribe All-In-One Topic: order
sub, err := wsclient.SubscribeOrder("all")
if err != nil {
    panic(err)
}


//Listen for orders 
for {
    select {
    case err := <-sub.ErrChan():
        panic(err)
    case result := <-sub.ResultChan():
        fmt.Println(result)
    case <-sub.CloseChan():
        fmt.Println("subscribe closed")
    }
}

```
# Trade Websocket Examples

## Create Order

```go
client := mybybitapi.NewRestClient("YOUR_API_KEY","YOUR_API_KEY")
privateClient := client.PrivateRestClient()
wsclient := mybybitapi.NewTradeWsStreamClient()
//Open Websocket Connect
err := wsclient.OpenConn()
if err != nil {
    panic(err)
}

//Auth To Bybit
err = wsclient.Auth(client)
if err != nil {
    panic(err)
}

//Websocket order create
orderCreateRes, err := wsclient.CreateOrder(
    privateClient.NewOrderCreate().
    Category("linear").Symbol("BTCUSDT").Side("Buy").OrderType("Limit").Qty("0.1").Price("10000"))
if err != nil {
    panic(err)
}
fmt.Println(orderCreateRes)
```