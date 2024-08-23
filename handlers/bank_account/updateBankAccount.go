package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/isaiaspereira307/gowallet/internal/db"
)

// @BasePath /api/v1
// @Summary Update a BankAccount
// @Description Update a BankAccount
// @Tags bank account
// @Accept json
// @Produce json
// @Param request body UpdateUserRequest true "Update BankAccount Request"
// @Success 200 {object} UpdateBankAccountResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /bank_account [put]
func UpdateBankAccount(ctx *gin.Context) {
	var req UpdateBankAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	newBankAccount := db.UpdateBankAccountParams{
		ID:        req.ID,
		Name:      req.Name,
		Balance:   req.Balance,
		UpdatedAt: time.Now(),
	}

	err := queries.UpdateBankAccount(ctx, newBankAccount)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update bank account"})
		return
	}

	ctx.JSON(http.StatusOK, req)
}
