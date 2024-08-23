package handlers

import (
	bankaccount_handler "github.com/isaiaspereira307/gowallet/handlers/bank_account"
	bitcoin_handler "github.com/isaiaspereira307/gowallet/handlers/bitcoin"
	fixedespense_handler "github.com/isaiaspereira307/gowallet/handlers/fixed_expense"
	investment_handlers "github.com/isaiaspereira307/gowallet/handlers/investment"
	loan_handlers "github.com/isaiaspereira307/gowallet/handlers/loan"
	transaction_handlers "github.com/isaiaspereira307/gowallet/handlers/transaction"
	users_handlers "github.com/isaiaspereira307/gowallet/handlers/user"
	"github.com/isaiaspereira307/gowallet/internal/db"
)

var queries *db.Queries

func InitializeHandlers(q *db.Queries) {
	queries = q
	bankaccount_handler.InitializeBankAccountHandlers(queries)
	bitcoin_handler.InitializeBitcoinHandlers(queries)
	fixedespense_handler.InitializeFixedExpenseHandlers(queries)
	investment_handlers.InitializeInvestmentHandlers(queries)
	loan_handlers.InitializeLoanHandlers(queries)
	transaction_handlers.InitializeTransactionHandlers(queries)
	users_handlers.InitializeUserHandlers(queries)
}
