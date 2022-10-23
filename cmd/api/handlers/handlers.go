package handlers

import (
	"auth/pkg/web"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
	"os"
)

func API(shutdown chan os.Signal, log *zap.SugaredLogger) (http.Handler, error) {
	app := web.NewApp(mux.NewRouter(), shutdown)
	return app, nil
}
