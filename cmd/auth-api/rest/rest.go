package rest

import (
	"auth/pkg/web"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
	"net/http"
	"os"
)

func API(shutdown chan os.Signal, log *zap.SugaredLogger, db *pgxpool.Pool) http.Handler {
	app := web.NewApp(mux.NewRouter(), shutdown)

	return app
}
