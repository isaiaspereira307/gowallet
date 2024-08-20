package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/isaiaspereira307/gowallet/internal/db"
)

// @BasePath /api/v1
// @Summary Delete an FixedExpense
// @Description Delete an FixedExpense
// @Tags fixed expense
// @Accept json
// @Produce json
// @Param id query string true "Delete FixedExpense Param"
// @Success 200 {object} DeleteFixedExpenseResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /fixed_expenses [delete]
func DeleteFixedExpense(ctx *gin.Context, queries *db.Queries) {
	id := ctx.Param("id")
	idInt32, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	err = queries.DeleteFixedExpense(ctx, int32(idInt32))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete fixed expense"})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
