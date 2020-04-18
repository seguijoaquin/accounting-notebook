package apierror

import "net/http"

//Error represents an HTTP response error
type Error interface {
	Message() string
	Status() int
	JSON() interface{}
}

type err struct {
	status  int
	message string
}

//Message returns the message for the current error
func (e *err) Message() string {
	return e.message
}

//Status returns the status code for the current error
func (e *err) Status() int {
	return e.status
}

//JSON returns the Json codification for the current error
func (e *err) JSON() interface{} {
	return &struct {
		Status  int    `json:"status_code"`
		Message string `json:"message"`
	}{
		e.status,
		e.message,
	}
}

//NewBadRequest creates an error configured to be a BadRequest response
func NewBadRequest(msg string) Error {
	return &err{
		status:  http.StatusBadRequest,
		message: msg,
	}
}

//NewInternalServerError creates an error configured to be an InternalServerError response
func NewInternalServerError(msg string) Error {
	return &err{
		status:  http.StatusInternalServerError,
		message: msg,
	}
}

//NewNotFoundError creates an error configured to be a NotFound response
func NewNotFoundError(msg string) Error {
	return &err{
		status:  http.StatusNotFound,
		message: msg,
	}
}

//NewUnprocessableEntity creates an error configured to be a UnprocessableEntity response
func NewUnprocessableEntity(msg string) Error {
	return &err{
		status:  http.StatusUnprocessableEntity,
		message: msg,
	}
}
