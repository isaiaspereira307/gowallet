package handlers

import (
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1
// @Summary Show Bitcoin
// @Description Show an Bitcoin
// @Tags bitcoin
// @Accept json
// @Produce json
// @Param id path string true "Show Bitcoin Request"
// @Success 200 {object} ShowBitcoinResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /bitcoin/{id} [get]
func GetBitcoin(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt64, err := strconv.ParseInt(id, 10, 64)
	if err != nil || idInt64 > math.MaxInt32 || idInt64 < math.MinInt32 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	idInt32 := int32(idInt64)
	bitcoin, err := queries.GetBitcoinById(ctx, idInt32)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get bitcoin"})
		return
	}

	ctx.JSON(http.StatusOK, bitcoin)
}
