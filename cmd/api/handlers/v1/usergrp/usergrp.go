package usergrp

import (
	"auth/internal/core/user"
	"auth/internal/service"
	v1 "auth/internal/web/v1"
	"auth/pkg/web"
	"context"
	"fmt"
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

func (h Handler) Register(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var nu user.New
	if err := web.Decode(r, &nu); err != nil {
		return fmt.Errorf("cannot read : %w", err)
	}

	if err := h.userService.Create(ctx, nu); err != nil {
		switch err {
		case service.ErrEmailAlreadyExists:
			return v1.NewRequestError(err, http.StatusConflict)
		case service.ErrUsernameAlreadyExists:
			return v1.NewRequestError(err, http.StatusConflict)
		default:
			return err
		}
	}

	return web.Respond(ctx, w, nil, http.StatusCreated)
}
