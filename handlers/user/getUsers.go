package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1
// @Summary Show users
// @Description Show all users
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {object} ListUserResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /users [get]
func GetUsers(ctx *gin.Context) {
	users, err := queries.GetUsers(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get user"})
		return
	}
	if len(users) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "no users found"})
		return
	}

	ctx.JSON(http.StatusOK, users)
}
