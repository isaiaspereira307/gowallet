package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1
// @Summary Show all transaction by account id
// @Description Show all transaction by account id
// @Tags transaction
// @Accept json
// @Produce json
// @Param id path string true "Show List Transaction Request"
// @Success 200 {object} ListTransactionResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /transactions/{id} [get]
func GetTransactionByAccountId(ctx *gin.Context) {
	accountID := ctx.Param("account_id")
	accountIDInt32, err := strconv.ParseInt(accountID, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid account_id"})
		return
	}
	transactions, err := queries.GetTransactionsByBankAccountId(ctx, int32(accountIDInt32))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get transactions"})
		return
	}

	ctx.JSON(http.StatusOK, transactions)
}
