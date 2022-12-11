package usergrp

import (
	"auth/internal/core/user"
	"auth/pkg/web"
	"context"
	"net/http"
)

type Handler struct {
	userService user.Service
}

func NewHandler(userService user.Service) Handler {
	return Handler{
		userService: userService,
	}
}

func (h Handler) Signup(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

	return web.Respond()
}

func (h Handler) Update(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

}

func (h Handler) Delete(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

}

func (h Handler) Get(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

}

func (h Handler) GetOne(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

}
