package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isaiaspereira307/gowallet/internal/db"
)

// @BasePath /api/v1
// @Summary Update an FixedExpense
// @Description Update an FixedExpense
// @Tags fixed expense
// @Accept json
// @Produce json
// @Param request body UpdateFixedExpenseRequest true "Update FixedExpense Request"
// @Success 200 {object} UpdateFixedExpenseResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /fixed-expense [put]
func UpdateFixedExpense(ctx *gin.Context) {
	var req UpdateFixedExpenseRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	err := req.Validate()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newFixedExpense := db.UpdateFixedExpenseParams{
		ID:          req.ID,
		Amount:      req.Amount,
		Description: req.Description,
	}
	err = queries.UpdateFixedExpense(ctx, newFixedExpense)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update fixed expense"})
		return
	}

	ctx.JSON(http.StatusOK, req)
}
