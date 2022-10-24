package service

import "github.com/pkg/errors"

var (
	ErrEmailAlreadyExists    = errors.New("user with such email already exists")
	ErrUsernameAlreadyExists = errors.New("user with such username already exists")
)
