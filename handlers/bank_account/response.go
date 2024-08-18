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

type CreateBankAccountResponse struct {
	Message string         `json:"message"`
	Data    db.BankAccount `json:"data"`
}

type DeleteBankAccountResponse struct {
	Message string         `json:"message"`
	Data    db.BankAccount `json:"data"`
}

type ShowBankAccountResponse struct {
	Message string         `json:"message"`
	Data    db.BankAccount `json:"data"`
}

type UpdateBankAccountResponse struct {
	Message string         `json:"message"`
	Data    db.BankAccount `json:"data"`
}

type ListBankAccountResponse struct {
	Message string           `json:"message"`
	Data    []db.BankAccount `json:"data"`
}
