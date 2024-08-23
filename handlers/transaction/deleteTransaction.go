package handlers

import (
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1
// @Summary Delete a transaction
// @Description Delete a transaction
// @Tags transaction
// @Accept json
// @Produce json
// @Param id path string true "Delete Transaction Param"
// @Success 200 {object} DeleteTransactionResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /transaction/{id} [delete]
func DeleteTransaction(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt64, err := strconv.ParseInt(id, 10, 64)
	if err != nil || idInt64 > math.MaxInt32 || idInt64 < math.MinInt32 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	idInt32 := int32(idInt64)

	err = queries.DeleteTransaction(ctx, idInt32)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete transaction"})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
