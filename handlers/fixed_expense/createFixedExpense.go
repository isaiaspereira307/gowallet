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
// @Summary Create an FixedExpense
// @Description Create an FixedExpense
// @Tags fixed expense
// @Accept json
// @Produce json
// @Param request body CreateFixedExpenseRequest true "Create FixedExpense Params"
// @Success 200 {object} CreateFixedExpenseResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /fixed-expense [post]
func CreateFixedExpense(ctx *gin.Context) {
	var req CreateFixedExpenseRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	newFixedExpense := db.CreateFixedExpenseParams{
		ID:            generateUniqueID(),
		BankAccountID: req.BankAccountID,
		Amount:        req.Amount,
		Description:   req.Description,
	}

	err := queries.CreateFixedExpense(ctx, newFixedExpense)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create fixed expense"})
		return
	}

	ctx.JSON(http.StatusCreated, req)
}
