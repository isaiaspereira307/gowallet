package routes

import (
	"github.com/gin-gonic/gin"
	investment_handlers "github.com/isaiaspereira307/gowallet/handlers/investment"
)

func InitializeInvestmentRoutes(router *gin.RouterGroup) {
	router.GET("/investments/:id", investment_handlers.GetInvestment)
	router.POST("/investments", investment_handlers.CreateInvestment)
	router.PUT("/investments/:id", investment_handlers.UpdateInvestment)
	router.DELETE("/investments/:id", investment_handlers.DeleteInvestment)
}
