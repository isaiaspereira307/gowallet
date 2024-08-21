package handlers

import (
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/isaiaspereira307/gowallet/internal/db"
)

// @BasePath /api/v1
// @Summary Show Loan
// @Description Show an Loan
// @Tags loan
// @Accept json
// @Produce json
// @Param id query string true "Show Loan Request"
// @Success 200 {object} ShowLoanResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /loans [get]
func GetLoan(ctx *gin.Context, queries *db.Queries) {
	id := ctx.Param("id")
	idInt64, err := strconv.ParseInt(id, 10, 64)
	if err != nil || idInt64 > math.MaxInt32 || idInt64 < math.MinInt32 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	idInt32 := int32(idInt64)
	loan, err := queries.GetLoanById(ctx, idInt32)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get loan"})
		return
	}

	ctx.JSON(http.StatusOK, loan)
}
