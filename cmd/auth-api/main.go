package main

import (
	"auth/cmd/auth-api/rest"
	"auth/internal/sys"
	"context"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	ctx := context.Background()

	// Parsing configs from env file
	conf, err := sys.NewConfigWithEnv()
	if err != nil {
		log.Fatalf("Failed to parse config : %v\n", err)
	}

	logger, err := sys.Logger(conf.Env)
	if err != nil {
		log.Fatalf("Failed to create logger : %v", err)
	}

	if err := run(ctx, logger, conf); err != nil {
		logger.Fatalf("err : %v", err)
	}

}

func run(ctx context.Context, log *zap.SugaredLogger, conf *sys.Config) error {
	log.Infof("Starting %s environment", conf.Env)
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	db, err := sys.Postgres(ctx, conf.Postgres.DSN)
	if err != nil {
		return err
	}

	app := rest.API(shutdown, log, db)

	log.Infof("Starting server on %s", conf.Addr)
	server := http.Server{
		Addr:    conf.Addr,
		Handler: app,
	}

	serverErrChan := make(chan error, 1)

	go func() {
		serverErrChan <- server.ListenAndServe()
	}()

	log.Infof("Server was successfully started => http://%s", conf.Addr)

	select {
	case <-serverErrChan:
		return err
	case <-shutdown:
		if err := server.Shutdown(context.Background()); err != nil {
			server.Close()
			return err
		}
		log.Info("Server was gracefully stopped")
	}

	return nil
}
