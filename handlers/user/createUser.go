package handlers

import (
	"net/http"

	"github.com/isaiaspereira307/gowallet/internal/db"

	"github.com/gin-gonic/gin"
)

func CreateUser(ctx *gin.Context) {
	var req db.CreateUserParams
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	err := queries.CreateUser(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
		return
	}

	ctx.JSON(http.StatusCreated, req)
}
