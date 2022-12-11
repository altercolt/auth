package user

import (
	"auth/internal/core/errors"
	"net/mail"
	"regexp"
)

var (
	usernameMatcher = regexp.MustCompile("^[A-Za-z0-9_]*$")
)

func (n New) Validate() error {
	validationError := errors.NewValidationError()

	if _, err := mail.ParseAddress(n.Email); err != nil {
		validationError.Append("email", err.Error())
	}

	usernameLenPass := true
	if len(n.Username) < 4 {
		validationError.Append("username", "length of username cannot be less than 4")
		usernameLenPass = false
	}

	if usernameLenPass {
		if ok := usernameMatcher.Match([]byte(n.Username)); !ok {
			validationError.Append("username", "username can only contain characters [a-z] and numbers [0-9]")
		}
	}

	if len(n.Password) < 8 {
		validationError.Append("password", "length of your password must be more than 8")
	}

	if validationError.Empty() {
		return nil
	}

	return validationError
}

func (u Update) Validate() error {
	validationError := errors.NewValidationError()

	if u.Username != nil {
		username := *u.Username
		usernameLenPass := true

		if len(username) < 4 {
			validationError.Append("username", "length of username cannot be less than 4")
			usernameLenPass = false
		}

		if ok := usernameMatcher.Match([]byte(username)); !ok && usernameLenPass {
			validationError.Append("username", "username can only contain characters [a-z] and numbers [0-9]")
		}

	}

	if u.Email != nil {
		if _, err := mail.ParseAddress(*u.Email); err != nil {
			validationError.Append("email", err.Error())
		}
	}

	if u.NewPassword != nil {
		if len(*u.NewPassword) < 8 {
			validationError.Append("new_password", "length of your new password must be more than 8")
		}
	}

	if u.Password == nil {
		validationError.Append("password", "password cannot be empty")
	}

	if u.Password != nil && len(*u.Password) < 8 {
		validationError.Append("password", "password length cannot be less than 8")
	}

	if validationError.Empty() {
		return nil
	}

	return validationError
}
