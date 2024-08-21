package handlers

type ErrorResponse struct {
	Message   string `json:"message"`
	ErrorCode string `json:"error_code"`
}

type LoginResponse struct {
	Message string `json:"message"`
	Data    string `json:"data"`
}
