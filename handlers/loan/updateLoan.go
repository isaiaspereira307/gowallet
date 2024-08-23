package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/isaiaspereira307/gowallet/internal/db"
)

// @BasePath /api/v1
// @Summary Update an Loan
// @Description Update an Loan
// @Tags loan
// @Accept json
// @Produce json
// @Param id query string true "Loan ID"
// @Param request body UpdateLoanRequest true "Update Loan Request"
// @Success 200 {object} UpdateLoanResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /loan [put]
func UpdateLoan(ctx *gin.Context) {
	var req UpdateLoanRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	err := req.Validate()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newLoan := db.UpdateLoanParams{
		ID:           req.ID,
		Amount:       req.Amount,
		InterestRate: req.InterestRate,
		DueDate:      req.DueDate,
		UpdatedAt:    time.Now(),
	}
	err = queries.UpdateLoan(ctx, newLoan)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update loan"})
		return
	}

	ctx.JSON(http.StatusOK, req)
}
