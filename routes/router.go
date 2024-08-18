package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/isaiaspereira307/gowallet/handlers"
	"github.com/isaiaspereira307/gowallet/internal/db"
)

func SetupRouter(queries *db.Queries) *gin.Engine {
	r := gin.Default()

	r.GET("/users/:id", func(ctx *gin.Context) { handlers.GetUser(ctx, queries) })
	r.POST("/users", func(ctx *gin.Context) { handlers.CreateUser(ctx, queries) })
	r.PUT("/users/:id", func(ctx *gin.Context) { handlers.UpdateUser(ctx, queries) })
	r.DELETE("/users/:id", func(ctx *gin.Context) { handlers.DeleteUser(ctx, queries) })

	r.GET("/fixed_expenses/:id", func(ctx *gin.Context) { handlers.GetFixedExpense(ctx, queries) })
	r.POST("/fixed_expenses", func(ctx *gin.Context) { handlers.CreateFixedExpense(ctx, queries) })
	r.PUT("/fixed_expenses/:id", func(ctx *gin.Context) { handlers.UpdateFixedExpense(ctx, queries) })
	r.DELETE("/fixed_expenses/:id", func(ctx *gin.Context) { handlers.DeleteFixedExpense(ctx, queries) })

	r.GET("/loans/:id", func(ctx *gin.Context) { handlers.GetLoan(ctx, queries) })
	r.POST("/loans", func(ctx *gin.Context) { handlers.CreateLoan(ctx, queries) })
	r.PUT("/loans/:id", func(ctx *gin.Context) { handlers.UpdateLoan(ctx, queries) })
	r.DELETE("/loans/:id", func(ctx *gin.Context) { handlers.DeleteLoan(ctx, queries) })

	r.GET("/bitcoins/:id", func(ctx *gin.Context) { handlers.GetBitcoin(ctx, queries) })
	r.POST("/bitcoins", func(ctx *gin.Context) { handlers.CreateBitcoin(ctx, queries) })
	r.PUT("/bitcoins/:id", func(ctx *gin.Context) { handlers.UpdateBitcoin(ctx, queries) })
	r.DELETE("/bitcoins/:id", func(ctx *gin.Context) { handlers.DeleteBitcoin(ctx, queries) })
	r.GET("/bitcoins/price", func(ctx *gin.Context) { handlers.GetBitcoinPriceUSD() })

	r.GET("investments/:id", func(ctx *gin.Context) { handlers.GetInvestment(ctx, queries) })
	r.POST("investments", func(ctx *gin.Context) { handlers.CreateInvestment(ctx, queries) })
	r.PUT("investments/:id", func(ctx *gin.Context) { handlers.UpdateInvestment(ctx, queries) })
	r.DELETE("investments/:id", func(ctx *gin.Context) { handlers.DeleteInvestment(ctx, queries) })

	r.GET("bank_accounts/:id", func(ctx *gin.Context) { handlers.GetBankAccount(ctx, queries) })
	r.POST("bank_accounts", func(ctx *gin.Context) { handlers.CreateBankAccount(ctx, queries) })
	r.PUT("bank_accounts/:id", func(ctx *gin.Context) { handlers.UpdateBankAccount(ctx, queries) })
	r.DELETE("bank_accounts/:id", func(ctx *gin.Context) { handlers.DeleteBankAccount(ctx, queries) })

	r.GET("transactions/:id", func(ctx *gin.Context) { handlers.GetTransaction(ctx, queries) })
	r.POST("transactions", func(ctx *gin.Context) { handlers.CreateTransaction(ctx, queries) })
	r.PUT("transactions/:id", func(ctx *gin.Context) { handlers.UpdateTransaction(ctx, queries) })
	r.DELETE("transactions/:id", func(ctx *gin.Context) { handlers.DeleteTransaction(ctx, queries) })

	return r
}
