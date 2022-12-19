package repository

import "github.com/pkg/errors"

var (
	ErrNotFound = errors.New("not found")
)

type UniqueError struct {
	message string
	column  string
}

func NewUniqueError(message, column string) *UniqueError {
	return &UniqueError{
		message: message,
		column:  column,
	}
}

func (e UniqueError) Error() string {
	return e.column + ": " + e.message
}
