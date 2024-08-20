package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isaiaspereira307/gowallet/internal/db"
)

// @BasePath /api/v1
// @Summary Create an Loan
// @Description Create an Loan
// @Tags loan
// @Accept json
// @Produce json
// @Param request body CreateLoanRequest true "Create Loan Params"
// @Success 200 {object} CreateLoanResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /loans [post]
func CreateLoan(ctx *gin.Context, queries *db.Queries) {
	var req db.CreateLoanParams
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	err := queries.CreateLoan(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create loan"})
		return
	}

	ctx.JSON(http.StatusCreated, req)
}
