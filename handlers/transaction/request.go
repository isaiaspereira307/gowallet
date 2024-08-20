package handlers

import (
	"fmt"
)

func errParamIsRequired(param string, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", param, typ)
}

type CreateTransactionRequest struct {
	BankAccountID int32  `json:"bank_account_id"`
	Amount        string `json:"amount"`
	Description   string `json:"description"`
	CreditDebit   bool   `json:"credit_debit"`
}

func (r *CreateTransactionRequest) Validate() error {
	fields := map[string]interface{}{
		"bank_account_id": r.BankAccountID,
		"amount":          r.Amount,
		"description":     r.Description,
		"credit_debit":    r.CreditDebit,
	}

	types := map[string]string{
		"bank_account_id": "int32",
		"amount":          "string",
		"description":     "string",
		"credit_debit":    "bool",
	}

	for field, value := range fields {
		switch v := value.(type) {
		case string:
			if v == "" {
				return errParamIsRequired(field, types[field])
			}
		case int32:
			if v == 0 {
				return errParamIsRequired(field, types[field])
			}
		}
	}

	return nil
}

type UpdateTransactionRequest struct {
	ID          int32  `json:"id" binding:"required"`
	BankAccount int32  `json:"bank_account_id"`
	Amount      string `json:"amount"`
	Description string `json:"description"`
	CreditDebit bool   `json:"credit_debit"`
}

func (r *UpdateTransactionRequest) Validate() error {
	if r.Amount != "" || r.Description != "" || r.BankAccount != 0 {
		return nil
	}
	return fmt.Errorf("at least one valid field must be provided")
}
