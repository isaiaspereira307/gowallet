package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isaiaspereira307/gowallet/internal/db"
)

// @BasePath /api/v1
// @Summary Update an Bitcoin
// @Description Update an Bitcoin
// @Tags bitcoin
// @Accept json
// @Produce json
// @Param request body UpdateBitcoinRequest true "Update Bitcoin Request"
// @Success 200 {object} UpdateBitcoinResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /bitcoin [put]
func UpdateBitcoin(ctx *gin.Context) {
	var req UpdateBitcoinRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	err := req.Validate()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newBitcoin := db.UpdateBitcoinParams{
		ID:            req.ID,
		PurchasePrice: req.PurchasePrice,
		Quantity:      req.Quantity,
		PurchaseDate:  req.PurchaseDate,
	}

	err = queries.UpdateBitcoin(ctx, newBitcoin)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update bitcoin"})
		return
	}

	ctx.JSON(http.StatusOK, req)
}
