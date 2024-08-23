package handlers

import (
	"log"
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
// @Summary Create a BankAccount
// @Description Create a BankAccount
// @Tags bank account
// @Accept json
// @Produce json
// @Param request body CreateBankAccountRequest true "Create BankAccount Params"
// @Success 200 {object} CreateBankAccountResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /bank-account [post]
func CreateBankAccount(ctx *gin.Context) {
	var req CreateBankAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Printf("error: %v", err)
		sendErr(ctx, err, http.StatusBadRequest)
		return
	}

	err := req.Validate()
	if err != nil {
		sendErr(ctx, err, http.StatusBadRequest)
		return
	}

	newBankAccount := db.CreateBankAccountParams{
		ID:        generateUniqueID(),
		UserID:    req.UserID,
		Name:      req.Name,
		Balance:   req.Balance,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = queries.CreateBankAccount(ctx, newBankAccount)
	if err != nil {
		sendErr(ctx, err, http.StatusInternalServerError)
		return
	}

	sendSuccess(ctx, "createBankAccount", req)
}
