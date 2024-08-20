package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

// @BasePath /api/v1
// @Summary Show Bitcoin Price USD
// @Description Show an Bitcoin Price USD
// @Tags bitcoin
// @Accept json
// @Produce json
// @Success 200 {object} float64
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /bitcoin-price [get]
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
