package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/isaiaspereira307/gowallet/internal/db"
)

// @BasePath /api/v1
// @Summary Update an Investment
// @Description Update an Investment
// @Tags investment
// @Accept json
// @Produce json
// @Param id query string true "Investment ID"
// @Param request body UpdateInvestmentRequest true "Update Investment Request"
// @Success 200 {object} UpdateInvestmentResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /investments [put]
func UpdateInvestment(ctx *gin.Context, queries *db.Queries) {
	id := ctx.Param("id")
	idInt32, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var req db.UpdateInvestmentParams
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	req.ID = int32(idInt32)
	err = queries.UpdateInvestment(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update investment"})
		return
	}

	ctx.JSON(http.StatusOK, req)
}
