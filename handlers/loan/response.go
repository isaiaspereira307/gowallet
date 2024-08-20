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

type CreateLoanResponse struct {
	Message string  `json:"message"`
	Data    db.Loan `json:"data"`
}

type DeleteLoanResponse struct {
	Message string  `json:"message"`
	Data    db.Loan `json:"data"`
}

type ShowLoanResponse struct {
	Message string  `json:"message"`
	Data    db.Loan `json:"data"`
}

type UpdateLoanResponse struct {
	Message string  `json:"message"`
	Data    db.Loan `json:"data"`
}

type ListLoanResponse struct {
	Message string    `json:"message"`
	Data    []db.Loan `json:"data"`
}
