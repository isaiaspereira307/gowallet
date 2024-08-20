package handlers

import (
	"net/http"
	"strconv"

	"github.com/isaiaspereira307/gowallet/internal/db"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1
// @Summary Delete an user
// @Description Delete an user
// @Tags user
// @Accept json
// @Produce json
// @Param id query string true "Delete User Param"
// @Success 200 {object} DeleteUserResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /users [delete]
func DeleteUser(ctx *gin.Context, queries *db.Queries) {
	id := ctx.Param("id")
	idInt32, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	err = queries.DeleteUser(ctx, int32(idInt32))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete user"})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
