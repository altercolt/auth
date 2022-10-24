package user

import (
	"auth/internal/core/errors"
	"net/mail"
)

func (n New) Validate() error {
	var errMap map[string]string

	if len(n.Username) < 4 {
		errMap["username"] = "length of username cannot be less than 4"
	}

	if _, err := mail.ParseAddress(n.Email); err != nil {
		errMap["email"] = err.Error()
	}

	if len(n.Password) < 8 {
		errMap["password"] = "length of the password cannot be less than 8"
	}

	if len(errMap) > 0 {
		return errors.NewValidationError("NewUser validation error", errMap)
	}

	return nil
}

func (u Update) Validate() error {
	var errMap map[string]string

	if u.Username != nil {
		if len(*u.Username) < 4 {
			errMap["username"] = "length of username cannot be less than 4"
		}
	}

	if u.Email != nil {
		if _, err := mail.ParseAddress(*u.Email); err != nil {
			errMap["email"] = err.Error()
		}
	}

	if u.NewPassword != nil {
		if len(*u.NewPassword) < 8 {
			errMap["password"] = "length of the password cannot be less than 8"
		}
	}

	if u.Password == nil {
		errMap["password"] = "password cannot be empty"
	} else {
		if len(*u.Password) < 8 {
			errMap["password"] = "length of the password is less than 8"
		}
	}

	if len(errMap) > 0 {
		return errors.NewValidationError("UpdateUser validation error", errMap)
	}

	return nil
}
