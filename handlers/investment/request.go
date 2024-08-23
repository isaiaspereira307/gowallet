package handlers

import "fmt"

func errParamIsRequired(param string, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", param, typ)
}

type CreateInvestmentRequest struct {
	BankAccountID int32  `json:"bank_account_id" binding:"required"`
	Amount        string `json:"amount" binding:"required"`
}

func (r *CreateInvestmentRequest) Validate() error {
	fields := map[string]interface{}{
		"bank_account_id": r.BankAccountID,
		"amount":          r.Amount,
	}

	types := map[string]string{
		"bank_account_id": "int32",
		"amount":          "string",
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

type UpdateInvestmentRequest struct {
	ID     int32  `json:"id" binding:"required"`
	Amount string `json:"amount" binding:"required"`
}

func (r *UpdateInvestmentRequest) Validate() error {
	if r.Amount != "" {
		return nil
	}
	return fmt.Errorf("at least one valid field must be provided")
}
