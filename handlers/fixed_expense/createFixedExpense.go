package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isaiaspereira307/gowallet/internal/db"
)

func CreateFixedExpense(ctx *gin.Context, queries *db.Queries) {
	var req db.CreateFixedExpenseParams
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	err := queries.CreateFixedExpense(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create fixed expense"})
		return
	}

	ctx.JSON(http.StatusCreated, req)
}
