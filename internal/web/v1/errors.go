package v1

import "errors"

type RequestError struct {
	Err    error
	Status int
}

func NewRequestError(err error, status int) error {
	return RequestError{
		Err:    err,
		Status: status,
	}
}

func (r RequestError) Error() string {
	return r.Err.Error()
}

// IsRequestError checks if an error of type RequestError exists.
func IsRequestError(err error) bool {
	var re *RequestError
	return errors.As(err, &re)
}

// GetRequestError returns a copy of the RequestError pointer.
func GetRequestError(err error) *RequestError {
	var re *RequestError
	if !errors.As(err, &re) {
		return nil
	}
	return re
}
