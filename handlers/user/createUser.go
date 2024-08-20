package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/isaiaspereira307/gowallet/internal/db"
)

// @BasePath /api/v1
// @Summary Create an user
// @Description Create an user
// @Tags user
// @Accept json
// @Produce json
// @Param request body CreateUserRequest true "Create User Params"
// @Success 200 {object} CreateUserResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /users [post]
func CreateUser(ctx *gin.Context) {
	var req CreateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		sendErr(ctx, err, http.StatusBadRequest)
		return
	}
	users, err := queries.GetUsers(ctx)
	if err != nil {
		sendErr(ctx, err, http.StatusInternalServerError)
		return
	}
	id := len(users) + 1
	user := db.CreateUserParams{
		ID:        int32(id),
		Name:      req.Name,
		Email:     req.Email,
		Password:  req.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = queries.CreateUser(ctx, user)
	if err != nil {
		sendErr(ctx, err, http.StatusInternalServerError)
		return
	}

	sendSuccess(ctx, "createUser", req)
}
