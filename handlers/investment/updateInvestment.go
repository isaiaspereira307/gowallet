package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/isaiaspereira307/gowallet/internal/db"
)

// @BasePath /api/v1
// @Summary Update an Investment
// @Description Update an Investment
// @Tags investment
// @Accept json
// @Produce json
// @Param request body UpdateInvestmentRequest true "Update Investment Request"
// @Success 200 {object} UpdateInvestmentResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /investment [put]
func UpdateInvestment(ctx *gin.Context) {
	var req UpdateInvestmentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	err := req.Validate()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newInvestment := db.UpdateInvestmentParams{
		ID:        req.ID,
		Amount:    req.Amount,
		UpdatedAt: time.Now(),
	}

	err = queries.UpdateInvestment(ctx, newInvestment)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update investment"})
		return
	}

	ctx.JSON(http.StatusOK, req)
}
