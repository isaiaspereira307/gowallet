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
// @Summary Create an Investment
// @Description Create an Investment
// @Tags investment
// @Accept json
// @Produce json
// @Param request body CreateInvestmentRequest true "Create Investment Params"
// @Success 200 {object} CreateInvestmentResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /investment [post]
func CreateInvestment(ctx *gin.Context) {
	var req CreateInvestmentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	newInvestment := db.CreateInvestmentParams{
		ID:            generateUniqueID(),
		BankAccountID: req.BankAccountID,
		Amount:        req.Amount,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	err := queries.CreateInvestment(ctx, newInvestment)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create investment"})
		return
	}

	ctx.JSON(http.StatusCreated, req)
}
