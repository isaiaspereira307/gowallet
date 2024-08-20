package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/isaiaspereira307/gowallet/internal/db"
)

// @BasePath /api/v1
// @Summary Show Investment
// @Description Show an Investment
// @Tags investment
// @Accept json
// @Produce json
// @Param id query string true "Show Investment Request"
// @Success 200 {object} ShowInvestmentResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /investments [get]
func GetInvestment(ctx *gin.Context, queries *db.Queries) {
	id := ctx.Param("id")
	idInt32, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	investment, err := queries.GetInvestmentById(ctx, int32(idInt32))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get investment"})
		return
	}

	ctx.JSON(http.StatusOK, investment)
}
