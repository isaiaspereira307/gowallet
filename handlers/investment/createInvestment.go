package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isaiaspereira307/gowallet/internal/db"
)

// @BasePath /api/v1
// @Summary Create an Investment
// @Description Create an Investment
// @Tags investment
// @Accept json
// @Produce json
// @Param request body CreateInvestmentRequest true "Create Investment Params"
// @Success 200 {object} CreateInvestmentResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /investments [post]
func CreateInvestment(ctx *gin.Context, queries *db.Queries) {
	var req db.CreateInvestmentParams
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	err := queries.CreateInvestment(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create investment"})
		return
	}

	ctx.JSON(http.StatusCreated, req)
}
