// routes/routes.go
package router

import (
	"github.com/gin-gonic/gin"
	"github.com/isaiaspereira307/gowallet/handler"
	"github.com/isaiaspereira307/gowallet/internal/db"
)

func SetupRouter(queries *db.Queries) *gin.Engine {
	r := gin.Default()

	r.GET("/users/:id", func(ctx *gin.Context) { handler.GetUser(ctx, queries) })
	r.POST("/users", func(ctx *gin.Context) { handler.CreateUser(ctx, queries) })
	r.PUT("/users/:id", func(ctx *gin.Context) { handler.UpdateUser(ctx, queries) })
	r.DELETE("/users/:id", func(ctx *gin.Context) { handler.DeleteUser(ctx, queries) })

	return r
}
