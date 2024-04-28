package api

import (
	"net/http"
)

var (
	ErrorServer    = NewErrorWithStatus(http.StatusOK, 1, "sorry, an error occurred on the server")
	ErrorParameter = NewError(2, "invalid parameter")
)

type Error struct {
	Status int
	Code   int
	Text   string
}

func (e *Error) Error() string {
	return e.Text
}

func NewError(code int, text string) *Error {
	return NewErrorWithStatus(http.StatusBadRequest, code, text)
}

func NewErrorWithStatus(status, code int, text string) *Error {
	return &Error{status, code, text}
}
