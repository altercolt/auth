package errors

type ValidationError struct {
	message  string
	errorMap map[string]string
}

func NewValidationError(message string, errorMap map[string]string) error {
	return ValidationError{
		message:  message,
		errorMap: errorMap,
	}
}

func (e ValidationError) Error() string {
	return e.message
}
