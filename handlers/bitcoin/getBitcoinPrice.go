package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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
func GetBitcoinPriceUSD(ctx *gin.Context) {
	url := "https://api.bitfinex.com/v1/pubticker/btcusd"
	req, err := http.Get(url)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get bitcoin"})
	}
	defer req.Body.Close()
	body, err := io.ReadAll(req.Body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get bitcoin"})
	}
	var bitcoinUsd BitcoinUsd
	err = json.Unmarshal(body, &bitcoinUsd)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get bitcoin"})
	}
	floatNum, err := strconv.ParseFloat(bitcoinUsd.LastPrice, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get bitcoin"})
	}
	ctx.JSON(http.StatusOK, floatNum)
}
