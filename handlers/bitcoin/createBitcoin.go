package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isaiaspereira307/gowallet/internal/db"
)

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
