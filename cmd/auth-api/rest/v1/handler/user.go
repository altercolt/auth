package handler

import (
	"auth/internal/core/user"
	"context"
	"net/http"
)

type User struct {
	userService user.Service
}

func NewHandler(userService user.Service) User {
	return User{
		userService: userService,
	}
}

func (h User) Signup(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

	return nil
}

func (h User) Update(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (h User) Delete(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (h User) Get(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (h User) GetOne(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	return nil
}
