package handlers

import "fmt"

func errParamIsRequired(param string, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", param, typ)
}

type CreateUserRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (r *CreateUserRequest) Validate() error {
	fields := map[string]interface{}{
		"name":     r.Name,
		"email":    r.Email,
		"password": r.Password,
	}

	types := map[string]string{
		"name":     "string",
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

type UpdateUserRequest struct {
	ID       int32  `json:"id" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (r *UpdateUserRequest) Validate() error {
	if r.Name != "" || r.Email != "" || r.Password != "" {
		return nil
	}
	return fmt.Errorf("at least one valid field must be provided")
}
