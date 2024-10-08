package handlers

import (
	"fmt"
)

func errParamIsRequired(param string, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", param, typ)
}

type CreateFixedExpenseRequest struct {
	BankAccountID int32  `json:"bank_account_id" binding:"required"`
	Amount        string `json:"amount" binding:"required"`
	Description   string `json:"description" binding:"required"`
}

func (r *CreateFixedExpenseRequest) Validate() error {
	fields := map[string]interface{}{
		"bank_account_id": r.BankAccountID,
		"amount":          r.Amount,
		"description":     r.Description,
	}

	types := map[string]string{
		"bank_account_id": "int32",
		"amount":          "string",
		"description":     "string",
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

type UpdateFixedExpenseRequest struct {
	ID            int32  `json:"id" binding:"required"`
	BankAccountID int32  `json:"bank_account_id" binding:"required"`
	Amount        string `json:"amount" binding:"required"`
	Description   string `json:"description" binding:"required"`
}

func (r *UpdateFixedExpenseRequest) Validate() error {
	if r.BankAccountID != 0 || r.Amount != "" || r.Description != "" {
		return nil
	}
	return fmt.Errorf("at least one valid field must be provided")
}
