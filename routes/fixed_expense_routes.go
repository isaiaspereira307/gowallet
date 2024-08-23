package routes

import (
	"github.com/gin-gonic/gin"
	fixed_expense_handlers "github.com/isaiaspereira307/gowallet/handlers/fixed_expense"
)

func InitializeFixedExpenseRoutes(router *gin.RouterGroup) {
	router.GET("/fixed_expenses/:id", fixed_expense_handlers.GetFixedExpense)
	router.POST("/fixed_expenses", fixed_expense_handlers.CreateFixedExpense)
	router.PUT("/fixed_expenses/:id", fixed_expense_handlers.UpdateFixedExpense)
	router.DELETE("/fixed_expenses/:id", fixed_expense_handlers.DeleteFixedExpense)
}
