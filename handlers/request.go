package handlers

import "fmt"

func errParamIsRequired(param string, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", param, typ)
}

type CredentialsRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (r *CredentialsRequest) Validate() error {
	fields := map[string]interface{}{
		"email":    r.Email,
		"password": r.Password,
	}

	types := map[string]string{
		"email":    "string",
		"passwrod": "string",
	}

	for field, value := range fields {
		switch v := value.(type) {
		case string:
			if v == "" {
				return errParamIsRequired(field, types[field])
			}
		}
	}

	return nil
}
