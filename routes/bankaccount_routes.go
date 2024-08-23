package routes

import (
	"github.com/gin-gonic/gin"
	bankaccount_handlers "github.com/isaiaspereira307/gowallet/handlers/bank_account"
)

func InitializeBankAccountRoutes(router *gin.RouterGroup) {
	router.GET("/bank-account/:id", bankaccount_handlers.GetBankAccount)
	router.GET("/bank-accounts/:id", bankaccount_handlers.GetBankAccountsByUserId)
	router.POST("/bank-account", bankaccount_handlers.CreateBankAccount)
	router.PUT("/bank-account/:id", bankaccount_handlers.UpdateBankAccount)
	router.DELETE("/bank-account/:id", bankaccount_handlers.DeleteBankAccount)
}
