package handler

import (
	"auth/internal/core/auth"
	"context"
	"net/http"
)

type Auth struct {
	authService auth.Service
}

func NewAuthHandler(authService auth.Service) Auth {
	return Auth{
		authService: authService,
	}
}

func (h Auth) Login(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (h Auth) Logout(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	return nil
}
