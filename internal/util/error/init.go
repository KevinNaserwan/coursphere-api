package error

import (
	"fmt"
	"net/http"
)

type ClientError struct {
	Code    int
	Message string
}

func (e ClientError) Error() string {
	return fmt.Sprintf("%d\t%s", e.Code, e.Message)
}

func NewBadRequest(msg string) *ClientError {
	return &ClientError{
		Code:    http.StatusBadRequest,
		Message: msg,
	}
}

func NewNotFound(msg string) *ClientError {
	return &ClientError{
		Code:    http.StatusNotFound,
		Message: msg,
	}
}

func NewForbidden(msg string) *ClientError {
	return &ClientError{
		Code:    http.StatusForbidden,
		Message: msg,
	}
}

func NewUnauthorized(msg string) *ClientError {
	return &ClientError{
		Code:    http.StatusUnauthorized,
		Message: msg,
	}
}
