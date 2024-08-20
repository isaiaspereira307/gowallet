package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/isaiaspereira307/gowallet/internal/db"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1
// @Summary Update an user
// @Description Update an user
// @Tags user
// @Accept json
// @Produce json
// @Param id query string true "User ID"
// @Param request body UpdateUserRequest true "Update User Request"
// @Success 200 {object} UpdateUserResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /users [put]
func UpdateUser(ctx *gin.Context, queries *db.Queries) {
	id := ctx.Param("id")
	idInt32, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse{Message: "invalid id"})
		return
	}

	var req UpdateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse{Message: "invalid request"})
		return
	}

	_, err = queries.GetUserById(ctx, int32(idInt32))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse{Message: "failed to get user"})
		return
	}

	newUser := db.UpdateUserParams{
		ID:        int32(idInt32),
		Name:      req.Name,
		Email:     req.Email,
		Password:  req.Password,
		UpdatedAt: time.Now(),
	}
	err = queries.UpdateUser(ctx, newUser)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse{Message: "failed to update user"})
		return
	}
	user := db.User{
		ID:       int32(idInt32),
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}
	resp := UpdateUserResponse{
		Message: "user updated successfully",
		Data:    user,
	}
	ctx.JSON(http.StatusOK, resp)
}
