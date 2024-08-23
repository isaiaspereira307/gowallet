package routes

import (
	"github.com/gin-gonic/gin"
	loan_handlers "github.com/isaiaspereira307/gowallet/handlers/loan"
)

func InitializeLoanRoutes(router *gin.RouterGroup) {
	router.GET("/loans/:id", loan_handlers.GetLoan)
	router.POST("/loans", loan_handlers.CreateLoan)
	router.PUT("/loans/:id", loan_handlers.UpdateLoan)
	router.DELETE("/loans/:id", loan_handlers.DeleteLoan)
}
