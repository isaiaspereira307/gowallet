package routes

import (
	"github.com/gin-gonic/gin"
	user_handlers "github.com/isaiaspereira307/gowallet/handlers/user"
)

func InitializeUserRoutes(router *gin.RouterGroup) {
	router.GET("/user/:id", user_handlers.GetUser)
	router.POST("/user", user_handlers.CreateUser)
	router.PUT("/user/:id", user_handlers.UpdateUser)
	router.DELETE("/user/:id", user_handlers.DeleteUser)
}
