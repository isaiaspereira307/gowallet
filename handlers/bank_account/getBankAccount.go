package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/isaiaspereira307/gowallet/internal/db"
)

// @BasePath /api/v1
// @Summary Show a BankAccount
// @Description Show an BankAccount
// @Tags bank account
// @Accept json
// @Produce json
// @Param id query string true "Show BankAccount Request"
// @Success 200 {object} ShowBankAccountResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /bank_accounts [get]
func GetBankAccount(ctx *gin.Context, queries *db.Queries) {
	id := ctx.Param("id")
	idInt32, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	bankAccount, err := queries.GetBankAccountById(ctx, int32(idInt32))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get bank account"})
		return
	}

	ctx.JSON(http.StatusOK, bankAccount)
}
