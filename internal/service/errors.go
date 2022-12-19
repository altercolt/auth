package service

import "errors"

var (
	ErrInvalidCredentials = errors.New("err : invalid credentials")
)

type DuplicateEntryError struct {
	msg     string
	field   string
	repoErr error
}

func NewDuplicateEntryError(msg string, field string, repoErr error) error {
	return &DuplicateEntryError{
		msg:     msg,
		field:   field,
		repoErr: repoErr,
	}
}

func (e DuplicateEntryError) Error() string {
	return "duplicateEntryError : " + e.msg + " on field:  " + "[" + e.field + "]"
}

func (e DuplicateEntryError) Unwrap() error {
	return e.repoErr
}
