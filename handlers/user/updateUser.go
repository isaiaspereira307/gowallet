package handlers

import (
	"net/http"
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
// @Param request body UpdateUserRequest true "Update User Request"
// @Success 200 {object} UpdateUserResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /user [put]
func UpdateUser(ctx *gin.Context) {
	var req UpdateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse{Message: "invalid request"})
		return
	}

	err := req.Validate()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
		return
	}

	newUser := db.UpdateUserParams{
		ID:        req.ID,
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
	ctx.JSON(http.StatusOK, req)
}
