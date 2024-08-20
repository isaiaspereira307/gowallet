package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/isaiaspereira307/gowallet/internal/db"
)

// @BasePath /api/v1
// @Summary Update an FixedExpense
// @Description Update an FixedExpense
// @Tags fixed expense
// @Accept json
// @Produce json
// @Param id query string true "FixedExpense ID"
// @Param request body UpdateFixedExpenseRequest true "Update FixedExpense Request"
// @Success 200 {object} UpdateFixedExpenseResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /fixed_expenses [put]
func UpdateFixedExpense(ctx *gin.Context, queries *db.Queries) {
	id := ctx.Param("id")
	idInt32, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var req db.UpdateFixedExpenseParams
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	req.ID = int32(idInt32)
	err = queries.UpdateFixedExpense(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update fixed expense"})
		return
	}

	ctx.JSON(http.StatusOK, req)
}
