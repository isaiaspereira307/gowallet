package routes

import (
	"github.com/gin-gonic/gin"
	bank_account_handlers "github.com/isaiaspereira307/gowallet/handlers/bank_account"
	bitcoin_handlers "github.com/isaiaspereira307/gowallet/handlers/bitcoin"
	fixed_expense_handlers "github.com/isaiaspereira307/gowallet/handlers/fixed_expense"
	investment_handlers "github.com/isaiaspereira307/gowallet/handlers/investment"
	loan_handlers "github.com/isaiaspereira307/gowallet/handlers/loan"
	transaction_handlers "github.com/isaiaspereira307/gowallet/handlers/transaction"
	user_handlers "github.com/isaiaspereira307/gowallet/handlers/user"
	"github.com/isaiaspereira307/gowallet/internal/db"
)

func InitializeRoutes(router *gin.Engine, queries *db.Queries) {
	user_handlers.InitializeUserHandlers(queries)
	bank_account_handlers.InitializeBankAccountHandlers(queries)
	bitcoin_handlers.InitializeBitcoinHandlers(queries)
	fixed_expense_handlers.InitializeFixedExpenseHandlers(queries)
	investment_handlers.InitializeInvestmentHandlers(queries)
	loan_handlers.InitializeLoanHandlers(queries)
	transaction_handlers.InitializeTransactionHandlers(queries)
	basePath := "/api/v1"
	v1 := router.Group(basePath)
	{
		v1.GET("/users/:id", func(ctx *gin.Context) {
			user_handlers.GetUser(ctx, queries)
		})
		v1.POST("/users", func(ctx *gin.Context) {
			user_handlers.CreateUser(ctx)
		})
		v1.PUT("/users/:id", func(ctx *gin.Context) {
			user_handlers.UpdateUser(ctx, queries)
		})
		v1.DELETE("/users/:id", func(ctx *gin.Context) {
			user_handlers.DeleteUser(ctx, queries)
		})

		v1.GET("/bank_accounts/:id", func(ctx *gin.Context) {
			bank_account_handlers.GetBankAccount(ctx, queries)
		})
		v1.POST("/bank_accounts", func(ctx *gin.Context) {
			bank_account_handlers.CreateBankAccount(ctx, queries)
		})
		v1.PUT("/bank_accounts/:id", func(ctx *gin.Context) {
			bank_account_handlers.UpdateBankAccount(ctx, queries)
		})
		v1.DELETE("/bank_accounts/:id", func(ctx *gin.Context) {
			bank_account_handlers.DeleteBankAccount(ctx, queries)
		})

		v1.GET("/bitcoins/:id", func(ctx *gin.Context) {
			bitcoin_handlers.GetBitcoin(ctx, queries)
		})
		v1.POST("/bitcoins", func(ctx *gin.Context) {
			bitcoin_handlers.CreateBitcoin(ctx, queries)
		})
		v1.PUT("/bitcoins/:id", func(ctx *gin.Context) {
			bitcoin_handlers.UpdateBitcoin(ctx, queries)
		})
		v1.DELETE("/bitcoins/:id", func(ctx *gin.Context) {
			bitcoin_handlers.DeleteBitcoin(ctx, queries)
		})

		v1.GET("/fixed_expenses/:id", func(ctx *gin.Context) {
			fixed_expense_handlers.GetFixedExpense(ctx, queries)
		})
		v1.POST("/fixed_expenses", func(ctx *gin.Context) {
			fixed_expense_handlers.CreateFixedExpense(ctx, queries)
		})
		v1.PUT("/fixed_expenses/:id", func(ctx *gin.Context) {
			fixed_expense_handlers.UpdateFixedExpense(ctx, queries)
		})
		v1.DELETE("/fixed_expenses/:id", func(ctx *gin.Context) {
			fixed_expense_handlers.DeleteFixedExpense(ctx, queries)
		})

		v1.GET("/investments/:id", func(ctx *gin.Context) {
			investment_handlers.GetInvestment(ctx, queries)
		})
		v1.POST("/investments", func(ctx *gin.Context) {
			investment_handlers.CreateInvestment(ctx, queries)
		})
		v1.PUT("/investments/:id", func(ctx *gin.Context) {
			investment_handlers.UpdateInvestment(ctx, queries)
		})
		v1.DELETE("/investments/:id", func(ctx *gin.Context) {
			investment_handlers.DeleteInvestment(ctx, queries)
		})

		v1.GET("/loans/:id", func(ctx *gin.Context) {
			loan_handlers.GetLoan(ctx, queries)
		})
		v1.POST("/loans", func(ctx *gin.Context) {
			loan_handlers.CreateLoan(ctx, queries)
		})
		v1.PUT("/loans/:id", func(ctx *gin.Context) {
			loan_handlers.UpdateLoan(ctx, queries)
		})
		v1.DELETE("/loans/:id", func(ctx *gin.Context) {
			loan_handlers.DeleteLoan(ctx, queries)
		})

		v1.GET("/transactions/:id", func(ctx *gin.Context) {
			transaction_handlers.GetTransaction(ctx, queries)
		})
		v1.POST("/transactions", func(ctx *gin.Context) {
			transaction_handlers.CreateTransaction(ctx, queries)
		})
		v1.PUT("/transactions/:id", func(ctx *gin.Context) {
			transaction_handlers.UpdateTransaction(ctx, queries)
		})
		v1.DELETE("/transactions/:id", func(ctx *gin.Context) {
			transaction_handlers.DeleteTransaction(ctx, queries)
		})
	}
}
