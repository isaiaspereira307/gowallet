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

type CreateInvestmentResponse struct {
	Message string        `json:"message"`
	Data    db.Investment `json:"data"`
}

type DeleteInvestmentResponse struct {
	Message string        `json:"message"`
	Data    db.Investment `json:"data"`
}

type ShowInvestmentResponse struct {
	Message string        `json:"message"`
	Data    db.Investment `json:"data"`
}

type UpdateInvestmentResponse struct {
	Message string        `json:"message"`
	Data    db.Investment `json:"data"`
}

type ListInvestmentResponse struct {
	Message string          `json:"message"`
	Data    []db.Investment `json:"data"`
}
