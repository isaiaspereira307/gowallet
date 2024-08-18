package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isaiaspereira307/gowallet/internal/db"
)

func CreateBankAccount(ctx *gin.Context, queries *db.Queries) {
	var req db.CreateBankAccountParams
	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Printf("error: %v", err)
		sendErr(ctx, err, http.StatusBadRequest)
		return
	}

	err := queries.CreateBankAccount(ctx, req)
	if err != nil {
		sendErr(ctx, err, http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusCreated, req.ID)
	sendSuccess(ctx, "createBankAccount", req)
}
