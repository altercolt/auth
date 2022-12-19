package middleware

import (
	"auth/pkg/web"
	"context"
	"net/http"
)

func Errors() web.Middleware {
	m := func(handler web.Handler) web.Handler {

		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

			return handler(ctx, w, r)
		}

		return h
	}

	return m
}
