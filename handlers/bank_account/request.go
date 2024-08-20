package handlers

import (
	"fmt"
	"time"
)

func errParamIsRequired(param string, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", param, typ)
}

type CreateBankAccountRequest struct {
	UserID  int32  `json:"user_id"`
	Name    string `json:"name"`
	Balance string `json:"balance"`
}

func (r *CreateBankAccountRequest) Validate() error {
	fields := map[string]interface{}{
		"user_id": r.UserID,
		"name":    r.Name,
		"balance": r.Balance,
	}

	types := map[string]string{
		"user_id": "int32",
		"name":    "string",
		"balance": "string",
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

type UpdateBankAccountRequest struct {
	ID        int32     `json:"id" binding:"required"`
	UserID    int32     `json:"user_id"`
	Name      string    `json:"name"`
	Balance   string    `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (r *UpdateBankAccountRequest) Validate() error {
	if r.Name != "" || r.UserID != 0 || r.Balance != "" {
		return nil
	}
	return fmt.Errorf("at least one valid field must be provided")
}
