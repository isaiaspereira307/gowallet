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
// @Summary Create a Bitcoin
// @Description Create a Bitcoin
// @Tags bitcoin
// @Accept json
// @Produce json
// @Param request body CreateBitcoinRequest true "Create Bitcoin Params"
// @Success 200 {object} db.CreateBitcoinParams
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /bitcoin [post]
func CreateBitcoin(ctx *gin.Context) {
	var req CreateBitcoinRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	newBitcoin := db.CreateBitcoinParams{
		ID:            generateUniqueID(),
		BankAccountID: req.BankAccountID,
		PurchasePrice: req.PurchasePrice,
		Quantity:      req.Quantity,
		PurchaseDate:  req.PurchaseDate,
	}

	err := queries.CreateBitcoin(ctx, newBitcoin)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create bitcoin"})
		return
	}

	ctx.JSON(http.StatusCreated, newBitcoin)
}
