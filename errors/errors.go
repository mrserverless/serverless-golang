package errors

import (
	"net/http"
)

type StatusError interface {
	Status() int
	StatusText() string
	error
}

type NotFoundError interface {
	StatusError
}

type UnprocessableEntityError interface {
	StatusError
}

type BaseStatusError struct {
	status     int
	statusText string
	error      string
}

func (e *BaseStatusError) Status() int {
	return e.status
}

func (e *BaseStatusError) StatusText() string {
	return e.statusText
}

func (e *BaseStatusError) Error() string {
	return e.error
}

func NewNotFound(err error) NotFoundError {
	return &BaseStatusError{http.StatusNotFound, http.StatusText(http.StatusNotFound), err.Error()}
}

func NewUnprocessableEntityError(err error) UnprocessableEntityError {
	return &BaseStatusError{http.StatusUnprocessableEntity, http.StatusText(http.StatusUnprocessableEntity), err.Error()}
}
