package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/isaiaspereira307/gowallet/internal/db"
)

// @BasePath /api/v1
// @Summary Show FixedExpense
// @Description Show an FixedExpense
// @Tags fixed expense
// @Accept json
// @Produce json
// @Param id query string true "Show FixedExpense Request"
// @Success 200 {object} ShowFixedExpenseResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /fixed_expenses [get]
func GetFixedExpense(ctx *gin.Context, queries *db.Queries) {
	id := ctx.Param("id")
	idInt32, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	fixedExpense, err := queries.GetFixedExpenseById(ctx, int32(idInt32))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get fixed expense"})
		return
	}

	ctx.JSON(http.StatusOK, fixedExpense)
}
