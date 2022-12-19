package service

type DuplicateEntryError struct {
	msg   string
	field string
	dbErr error
}

func NewDuplicateEntryError(msg string, field string, dbErr error) error {
	return &DuplicateEntryError{
		msg:   msg,
		field: field,
		dbErr: dbErr,
	}
}

func (e DuplicateEntryError) Error() string {
	return "duplicateEntryError : " + e.msg + " on field:  " + "[" + e.field + "]"
}

func (e DuplicateEntryError) Unwrap() error {
	return e.dbErr
}
