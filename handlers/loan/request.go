package handlers

import (
	"fmt"
	"time"
)

func errParamIsRequired(param string, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", param, typ)
}

type CreateLoanRequest struct {
	BankAccountID int32     `json:"bank_account_id" binding:"required"`
	Amount        string    `json:"amount" binding:"required"`
	InterestRate  string    `json:"interest_rate" binding:"required"`
	DueDate       time.Time `json:"due_date" binding:"required"`
}

func (r *CreateLoanRequest) Validate() error {
	fields := map[string]interface{}{
		"bank_account_id": r.BankAccountID,
		"amount":          r.Amount,
		"interest_rate":   r.InterestRate,
		"due_date":        r.DueDate,
	}

	types := map[string]string{
		"bank_account_id": "int32",
		"amount":          "string",
		"interest_rate":   "string",
		"due_date":        "time.Time",
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
		case time.Time:
			if v.IsZero() {
				return errParamIsRequired(field, types[field])
			}
		}
	}

	return nil
}

type UpdateLoanRequest struct {
	ID            int32     `json:"id" binding:"required"`
	BankAccountID int32     `json:"bank_account_id" binding:"required"`
	Amount        string    `json:"amount" binding:"required"`
	InterestRate  string    `json:"interest_rate" binding:"required"`
	DueDate       time.Time `json:"due_date" binding:"required"`
}

func (r *UpdateLoanRequest) Validate() error {
	if r.BankAccountID != 0 || r.Amount != "" || r.InterestRate != "" || !r.DueDate.IsZero() {
		return nil
	}
	return fmt.Errorf("at least one valid field must be provided")
}
