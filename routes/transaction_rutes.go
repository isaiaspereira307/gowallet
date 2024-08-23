package routes

import (
	"github.com/gin-gonic/gin"
	transaction_handlers "github.com/isaiaspereira307/gowallet/handlers/transaction"
)

func InitializeTransactionRoutes(router *gin.RouterGroup) {
	router.GET("/transaction/:id", transaction_handlers.GetTransaction)
	router.GET("/transactions/:id", transaction_handlers.GetTransactionByAccountId)
	router.POST("/transaction", transaction_handlers.CreateTransaction)
	router.PUT("/transaction/:id", transaction_handlers.UpdateTransaction)
	router.DELETE("/transaction/:id", transaction_handlers.DeleteTransaction)
}
