package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/isaiaspereira307/gowallet/internal/db"
)

// @BasePath /api/v1
// @Summary Update a Transaction
// @Description Update a Transaction
// @Tags transaction
// @Accept json
// @Produce json
// @Param request body UpdateTransactionRequest true "Update Transaction Request"
// @Success 200 {object} UpdateTransactionResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /transaction [put]
func UpdateTransaction(ctx *gin.Context) {
	var req UpdateTransactionRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	err := req.Validate()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newTransaction := db.UpdateTransactionParams{
		ID:          req.ID,
		Amount:      req.Amount,
		Description: req.Description,
		CreditDebit: req.CreditDebit,
		Timestamp:   time.Now(),
	}
	err = queries.UpdateTransaction(ctx, newTransaction)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update transaction"})
		return
	}

	ctx.JSON(http.StatusOK, req)
}
