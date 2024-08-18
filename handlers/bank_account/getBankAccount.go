package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/isaiaspereira307/gowallet/internal/db"
)

func GetBankAccount(ctx *gin.Context, queries *db.Queries) {
	id := ctx.Param("id")
	idInt32, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	bankAccount, err := queries.GetBankAccountById(ctx, int32(idInt32))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get bank account"})
		return
	}

	ctx.JSON(http.StatusOK, bankAccount)
}
