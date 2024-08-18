package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/isaiaspereira307/gowallet/internal/db"
)

func GetTransactionByAccount(ctx *gin.Context, queries *db.Queries) {
	accountID := ctx.Param("account_id")
	accountIDInt32, err := strconv.ParseInt(accountID, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid account_id"})
		return
	}
	transactions, err := queries.GetTransactionsByBankAccountId(ctx, int32(accountIDInt32))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get transactions"})
		return
	}

	ctx.JSON(http.StatusOK, transactions)
}
