package middleware

import (
	"auth/pkg/web"
	"context"
	"go.uber.org/zap"
	"net/http"
)

func Errors(log *zap.SugaredLogger) web.Middleware {

	m := func(handler web.Handler) web.Handler {
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
			err := handler(ctx, w, r)

		}
	}

	return m
}