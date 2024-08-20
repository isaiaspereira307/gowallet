package handlers

import (
	"fmt"
	"time"
)

func errParamIsRequired(param string, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", param, typ)
}

type CreateBitcoinRequest struct {
	BankAccountID int32     `json:"bank_account_id" binding:"required"`
	PurchasePrice string    `json:"purchase_price"`
	Quantity      string    `json:"quantity"`
	PurchaseDate  time.Time `json:"purchase_date"`
}

func (r *CreateBitcoinRequest) Validate() error {
	fields := map[string]interface{}{
		"bank_account_id": r.BankAccountID,
		"purchase_price":  r.PurchasePrice,
		"quantity":        r.Quantity,
		"purchase_date":   r.PurchaseDate,
	}

	types := map[string]string{
		"bank_account_id": "int32",
		"purchase_price":  "string",
		"quantity":        "string",
		"purchase_date":   "time.Time",
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

type UpdateBitcoinRequest struct {
	ID            int32     `json:"id" binding:"required"`
	BankAccountID int32     `json:"bank_account_id" binding:"required"`
	PurchasePrice string    `json:"purchase_price" binding:"required"`
	Quantity      string    `json:"quantity" binding:"required"`
	PurchaseDate  time.Time `json:"purchase_date" binding:"required"`
}

func (r *UpdateBitcoinRequest) Validate() error {
	if r.BankAccountID != 0 || r.PurchasePrice != "" || r.PurchaseDate != (time.Time{}) || r.Quantity != "" {
		return nil
	}
	return fmt.Errorf("at least one valid field must be provided")
}
