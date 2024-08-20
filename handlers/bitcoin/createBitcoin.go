package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isaiaspereira307/gowallet/internal/db"
)

// @BasePath /api/v1
// @Summary Create a Bitcoin
// @Description Create a Bitcoin
// @Tags bitcoin
// @Accept json
// @Produce json
// @Param request body CreateBankAccountRequest true "Create Bitcoin Params"
// @Success 200 {object} db.CreateBitcoinParams
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /bitcoins [post]
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
