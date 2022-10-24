package handlers

import (
	"auth/cmd/api/handlers/v1/usergrp"
	auth2 "auth/internal/core/auth"
	"auth/internal/core/role"
	"auth/internal/repository/postgres"
	"auth/internal/service"
	"auth/internal/web/v1/middleware"
	"auth/pkg/web"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
	"net/http"
	"os"
)

func API(shutdown chan os.Signal, log *zap.SugaredLogger, db *pgxpool.Pool) (http.Handler, error) {
	app := web.NewApp(mux.NewRouter(), shutdown)
	{
		userRepo := postgres.NewUserRepository(db)
		userService := service.NewUserService(log, userRepo)
		userHandler := usergrp.NewHandler(userService)
		var auth auth2.Service

		app.Handle(http.MethodPost, "/v1/user", userHandler.Register)
		app.Handle(http.MethodGet, "/v1/user", userHandler.GetSelf, middleware.Authorize(auth, role.User))
	}
	{

	}

	return app, nil
}
