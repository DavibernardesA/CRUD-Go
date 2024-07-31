package rest_err

import "net/http"

// ApiErr representa o objeto de erro.
// @Summary Informações sobre o erro
// @Description Estrutura para descrever por que o erro ocorreu
type ApiErr struct {
	Message string   `json:"message" example:"error trying to process request"`
	Err     string   `json:"error" example:"internal_server_error"`
	Code    int      `json:"code" example:"500"`
	Causes  []Causes `json:"causes"`
}

// Causes representa as causas do erro.
// @Summary Causas do erro
// @Description Estrutura representando as causas de um erro.
type Causes struct {
	Field   string `json:"field" example:"name"`
	Message string `json:"message" example:"name is required"`
}

func (r *ApiErr) Error() string {
	return r.Message
}

func NewBadRequestError(message string) *ApiErr {
	return &ApiErr{
		Message: message,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
	}
}

func NewUnauthorizedRequestError(message string) *ApiErr {
	return &ApiErr{
		Message: message,
		Err:     "unauthorized",
		Code:    http.StatusUnauthorized,
	}
}

func NewBadRequestValidationError(message string, causes []Causes) *ApiErr {
	return &ApiErr{
		Message: message,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
		Causes:  causes,
	}
}

func NewInternalServerError(message string) *ApiErr {
	return &ApiErr{
		Message: message,
		Err:     "internal_server_error",
		Code:    http.StatusInternalServerError,
	}
}

func NewNotFoundError(message string) *ApiErr {
	return &ApiErr{
		Message: message,
		Err:     "not_found",
		Code:    http.StatusNotFound,
	}
}

func NewForbiddenError(message string) *ApiErr {
	return &ApiErr{
		Message: message,
		Err:     "forbidden",
		Code:    http.StatusForbidden,
	}
}
