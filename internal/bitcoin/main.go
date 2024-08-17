package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type BitcoinUsd struct {
	Mid       string `json:"mid"`
	Bid       string `json:"bid"`
	Ask       string `json:"ask"`
	LastPrice string `json:"last_price"`
	Low       string `json:"low"`
	High      string `json:"high"`
}

func GetBitcoinPriceUSD() float64 {
	url := "https://api.bitfinex.com/v1/pubticker/btcusd"
	req, err := http.Get(url)
	if err != nil {
		return 0.0
	}
	defer req.Body.Close()
	body, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err.Error())
	}
	var bitcoinUsd BitcoinUsd
	err = json.Unmarshal(body, &bitcoinUsd)
	if err != nil {
		panic(err.Error())
	}
	floatNum, err := strconv.ParseFloat(bitcoinUsd.LastPrice, 64)
	if err != nil {
		fmt.Println("Erro na conversão:", err)
	}
	return floatNum
}
