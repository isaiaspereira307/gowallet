package handlers

import (
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1
// @Summary Show BankAccount
// @Description Show all BankAccount
// @Tags bank account
// @Accept json
// @Produce json
// @Param id path string true "Show BankAccount Request"
// @Success 200 {object} ListBankAccountResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /bank-accounts/{id} [get]
func GetBankAccountsByUserId(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt64, err := strconv.ParseInt(id, 10, 64)
	if err != nil || idInt64 > math.MaxInt32 || idInt64 < math.MinInt32 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	idInt32 := int32(idInt64)
	bank_accounts, err := queries.GetBankAccountsByUserId(ctx, idInt32)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get Category"})
		return
	}
	if len(bank_accounts) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "no Category found"})
		return
	}

	ctx.JSON(http.StatusOK, bank_accounts)
}
