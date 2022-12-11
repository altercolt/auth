package authgrp

import (
	"auth/internal/core/auth"
	"context"
	"net/http"
)

type Handler struct {
	authService auth.Service
}

func NewHandler(authService auth.Service) Handler {
	return Handler{
		authService: authService,
	}
}

func (h Handler) Login(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

}

func (h Handler) Logout(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

}
