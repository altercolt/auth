package repository

type UniqueError struct {
	message string
	// can't find a better name for it atm
	column string
	dbErr  error
}

func NewUniqueError(message, column string, dbErr error) *UniqueError {
	return &UniqueError{
		message: message,
		column:  column,
		dbErr:   dbErr,
	}
}

func (e UniqueError) Error() string {
	return e.column + ": " + e.message
}

func (e UniqueError) Unwrap() error {
	return e.dbErr
}

type NotFoundError struct {
	message   string
	entity    string
	clientErr error
}

func NewNotFoundError(message, entity string, clientErr error) *NotFoundError {
	return &NotFoundError{
		message:   message,
		entity:    entity,
		clientErr: clientErr,
	}
}

func (e NotFoundError) Error() string {
	return e.message + ": " + e.entity
}

func (e NotFoundError) Unwrap() error {
	return e.clientErr
}
