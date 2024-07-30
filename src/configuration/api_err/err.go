package err

import "net/http"

type restErr struct {
	Message string   `json: "message"`
	Err     string   `json: "error"`
	Code    int64    `json: "code"`
	Causes  []Causes `json: "causes"`
}

type Causes struct {
	Field   string `json: "field"`
	Message string `json: "message"`
}

func (r *restErr) Error() string {
	return r.Message
}

func apiErr(message, err string, code int64, causes []Causes) *restErr {
	return &restErr{
		Message: message,
		Err:     err,
		Code:    code,
		Causes:  causes,
	}
}

func newBadRequestError(message string) *restErr {
	return &restErr{
		Message: message,
		Err:     "Bad Request Error",
		Code:    http.StatusBadRequest,
	}
}

func newBadRequestValidationError(message string, causes []Causes) *restErr {
	return &restErr{
		Message: message,
		Err:     "Bad Request Error",
		Code:    http.StatusBadRequest,
		Causes:  causes,
	}
}

func newInternalServererror(message string) *restErr {
	return &restErr{
		Message: message,
		Err:     "Internal server Error",
		Code:    http.StatusInternalServerError,
	}
}

func newNotFoundError(message string) *restErr {
	return &restErr{
		Message: message,
		Err:     "Not found Error",
		Code:    http.StatusNotFound,
	}
}

func nweForbiddenError(message string) *restErr {
	return &restErr{
		Message: message,
		Err:     "Forbidden Error",
		Code:    http.StatusForbidden,
	}
}
