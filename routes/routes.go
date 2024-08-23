package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/isaiaspereira307/gowallet/docs"
	"github.com/isaiaspereira307/gowallet/handlers"
	"github.com/isaiaspereira307/gowallet/internal/db"
	"github.com/isaiaspereira307/gowallet/middleware"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitializeRoutes(router *gin.Engine, queries *db.Queries) {
	handlers.InitializeHandlers(queries)
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	InitializeAuthRoutes(router)
	protected := router.Group(basePath)
	protected.Use(middleware.AuthMiddleware())
	{
		InitializeUserRoutes(protected)
		InitializeBitcoinRoutes(protected)
		InitializeFixedExpenseRoutes(protected)
		InitializeInvestmentRoutes(protected)
		InitializeLoanRoutes(protected)
		InitializeBankAccountRoutes(protected)
		InitializeTransactionRoutes(protected)
	}
}
