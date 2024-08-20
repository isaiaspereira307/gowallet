package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isaiaspereira307/gowallet/internal/db"

	"github.com/isaiaspereira307/gowallet/configs"
)

var (
	logger *configs.Logger
)

type BitcoinUsd struct {
	Mid       string `json:"mid"`
	Bid       string `json:"bid"`
	Ask       string `json:"ask"`
	LastPrice string `json:"last_price"`
	Low       string `json:"low"`
	High      string `json:"high"`
}

func sendErr(ctx *gin.Context, err error, status int) {
	logger.Errorf("error: %s", err.Error())
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(status, gin.H{"error": err.Error()})
}

func sendSuccess(ctx *gin.Context, operation string, data interface{}) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{"operation": operation, "data": data})
}

type ErrorResponse struct {
	Message   string `json:"message"`
	ErrorCode string `json:"error_code"`
}

type CreateBitcoinResponse struct {
	Message string     `json:"message"`
	Data    db.Bitcoin `json:"data"`
}

type DeleteBitcoinResponse struct {
	Message string     `json:"message"`
	Data    db.Bitcoin `json:"data"`
}

type ShowBitcoinResponse struct {
	Message string     `json:"message"`
	Data    db.Bitcoin `json:"data"`
}

type UpdateBitcoinResponse struct {
	Message string     `json:"message"`
	Data    db.Bitcoin `json:"data"`
}

type ListBitcoinResponse struct {
	Message string       `json:"message"`
	Data    []db.Bitcoin `json:"data"`
}
