package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/isaiaspereira307/gowallet/internal/db"
)

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
// @Router /bank_accounts [post]
func CreateBankAccount(ctx *gin.Context, queries *db.Queries) {
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

	newBankAccount := db.CreateBankAccountParams{}
	newBankAccount.UserID = req.UserID
	newBankAccount.Name = req.Name
	newBankAccount.Balance = req.Balance
	newBankAccount.CreatedAt = time.Now()
	newBankAccount.UpdatedAt = time.Now()

	err = queries.CreateBankAccount(ctx, newBankAccount)
	if err != nil {
		sendErr(ctx, err, http.StatusInternalServerError)
		return
	}

	sendSuccess(ctx, "createBankAccount", req)
}
