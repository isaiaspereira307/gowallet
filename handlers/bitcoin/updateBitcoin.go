package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/isaiaspereira307/gowallet/internal/db"
)

// @BasePath /api/v1
// @Summary Update an Bitcoin
// @Description Update an Bitcoin
// @Tags bitcoin
// @Accept json
// @Produce json
// @Param id query string true "Bitcoin ID"
// @Param request body UpdateBitcoinRequest true "Update Bitcoin Request"
// @Success 200 {object} UpdateBitcoinResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /bitcoins [put]
func UpdateBitcoin(ctx *gin.Context, queries *db.Queries) {
	id := ctx.Param("id")
	idInt32, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var req db.UpdateBitcoinParams
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	req.ID = int32(idInt32)
	err = queries.UpdateBitcoin(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update bitcoin"})
		return
	}

	ctx.JSON(http.StatusOK, req)
}
