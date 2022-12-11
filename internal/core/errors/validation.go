package errors

import "encoding/json"

type ValidationError struct {
	err map[string]string
}

func NewValidationError() *ValidationError {
	return &ValidationError{
		err: map[string]string{},
	}
}

func (e *ValidationError) Error() string {
	msg, err := json.Marshal(e.err)
	if err != nil {
		return err.Error()
	}

	return string(msg)
}

func (e *ValidationError) Append(key, value string) {
	e.err[key] = value
}

func (e *ValidationError) Empty() bool {
	if e.err == nil || len(e.err) == 0 {
		return true
	}

	return false
}
