package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/isaiaspereira307/gowallet/internal/db"
)

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

func CreateBankAccount(ctx *gin.Context, queries *db.Queries) {
	var req db.CreateBankAccountParams
	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Printf("error: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "invalid request"})
		return
	}

	err := queries.CreateBankAccount(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create bank account"})
		return
	}

	ctx.JSON(http.StatusCreated, req.ID)
}

func UpdateBankAccount(ctx *gin.Context, queries *db.Queries) {
	id := ctx.Param("id")
	idInt32, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var req db.UpdateBankAccountParams
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	req.ID = int32(idInt32)
	err = queries.UpdateBankAccount(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update bank account"})
		return
	}

	ctx.JSON(http.StatusOK, req)
}

func DeleteBankAccount(ctx *gin.Context, queries *db.Queries) {
	id := ctx.Param("id")
	idInt32, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	err = queries.DeleteBankAccount(ctx, int32(idInt32))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete bank account"})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
