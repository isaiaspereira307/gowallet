package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/isaiaspereira307/gowallet/internal/db"
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

func GetBitcoin(ctx *gin.Context, queries *db.Queries) {
	id := ctx.Param("id")
	idInt32, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	bitcoin, err := queries.GetBitcoinById(ctx, int32(idInt32))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get bitcoin"})
		return
	}

	ctx.JSON(http.StatusOK, bitcoin)
}

func CreateBitcoin(ctx *gin.Context, queries *db.Queries) {
	var req db.CreateBitcoinParams
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	err := queries.CreateBitcoin(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create bitcoin"})
		return
	}

	ctx.JSON(http.StatusCreated, req)
}

func UpdateBitcoin(ctx *gin.Context, queries *db.Queries) {
	id := ctx.Param("id")
	idInt32, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var req db.UpdateBitcoinParams
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	req.ID = int32(idInt32)
	err = queries.UpdateBitcoin(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update bitcoin"})
		return
	}

	ctx.JSON(http.StatusOK, req)
}

func DeleteBitcoin(ctx *gin.Context, queries *db.Queries) {
	id := ctx.Param("id")
	idInt32, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	err = queries.DeleteBitcoin(ctx, int32(idInt32))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete bitcoin"})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
