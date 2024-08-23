package handlers

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/isaiaspereira307/gowallet/internal/db"
)

var (
	lastID int32
	mu     sync.Mutex
)

func generateUniqueID() int32 {
	mu.Lock()
	defer mu.Unlock()
	lastID++
	return lastID
}

// @BasePath /api/v1
// @Summary Create a Transaction
// @Description Create a Transaction
// @Tags transaction
// @Accept json
// @Produce json
// @Param request body CreateTransactionRequest true "Create Transaction Params"
// @Success 200 {object} CreateTransactionResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /transaction [post]
func CreateTransaction(ctx *gin.Context) {
	var req CreateTransactionRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	err := req.Validate()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newTransaction := db.CreateTransactionParams{
		ID:            generateUniqueID(),
		BankAccountID: req.BankAccountID,
		Amount:        req.Amount,
		Description:   req.Description,
		CreditDebit:   req.CreditDebit,
	}

	err = queries.CreateTransaction(ctx, newTransaction)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create transaction"})
		return
	}

	ctx.JSON(http.StatusCreated, req)
}
