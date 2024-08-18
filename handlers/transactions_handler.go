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

func GetTransaction(ctx *gin.Context, queries *db.Queries) {
	id := ctx.Param("id")
	idInt32, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	transactions, err := queries.GetTransactionById(ctx, int32(idInt32))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get transactions"})
		return
	}

	ctx.JSON(http.StatusOK, transactions)
}

func CreateTransaction(ctx *gin.Context, queries *db.Queries) {
	var req db.CreateTransactionParams
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	err := queries.CreateTransaction(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create transaction"})
		return
	}

	ctx.JSON(http.StatusCreated, req)
}

func UpdateTransaction(ctx *gin.Context, queries *db.Queries) {
	id := ctx.Param("id")
	idInt32, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var req db.UpdateTransactionParams
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	req.ID = int32(idInt32)
	err = queries.UpdateTransaction(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update transaction"})
		return
	}

	ctx.JSON(http.StatusOK, req)
}

func DeleteTransaction(ctx *gin.Context, queries *db.Queries) {
	id := ctx.Param("id")
	idInt32, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	err = queries.DeleteTransaction(ctx, int32(idInt32))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete transaction"})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
