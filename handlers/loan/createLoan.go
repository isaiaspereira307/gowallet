package handlers

import (
	"net/http"
	"sync"
	"time"

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
// @Summary Create an Loan
// @Description Create an Loan
// @Tags loan
// @Accept json
// @Produce json
// @Param request body CreateLoanRequest true "Create Loan Params"
// @Success 200 {object} CreateLoanResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /loan [post]
func CreateLoan(ctx *gin.Context) {
	var req CreateLoanRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	newLoan := db.CreateLoanParams{
		ID:            generateUniqueID(),
		BankAccountID: req.BankAccountID,
		Amount:        req.Amount,
		InterestRate:  req.InterestRate,
		DueDate:       req.DueDate,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	err := queries.CreateLoan(ctx, newLoan)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create loan"})
		return
	}

	ctx.JSON(http.StatusCreated, req)
}
