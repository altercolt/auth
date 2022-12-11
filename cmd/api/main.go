package main

import (
	"auth/cmd/api/handlers"
	"auth/internal/core/user"
	"auth/internal/sys"
	"context"
	"fmt"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var logo = "                                                                                                                                               \n    ,o888888o.        ,o888888o.     8 8888         8888888 8888888888          .8.          8 8888      88 8888888 8888888888 8 8888        8 \n   8888     `88.   . 8888     `88.   8 8888               8 8888               .888.         8 8888      88       8 8888       8 8888        8 \n,8 8888       `8. ,8 8888       `8b  8 8888               8 8888              :88888.        8 8888      88       8 8888       8 8888        8 \n88 8888           88 8888        `8b 8 8888               8 8888             . `88888.       8 8888      88       8 8888       8 8888        8 \n88 8888           88 8888         88 8 8888               8 8888            .8. `88888.      8 8888      88       8 8888       8 8888        8 \n88 8888           88 8888         88 8 8888               8 8888           .8`8. `88888.     8 8888      88       8 8888       8 8888        8 \n88 8888           88 8888        ,8P 8 8888               8 8888          .8' `8. `88888.    8 8888      88       8 8888       8 8888888888888 \n`8 8888       .8' `8 8888       ,8P  8 8888               8 8888         .8'   `8. `88888.   ` 8888     ,8P       8 8888       8 8888        8 \n   8888     ,88'   ` 8888     ,88'   8 8888               8 8888        .888888888. `88888.    8888   ,d8P        8 8888       8 8888        8 \n    `8888888P'        `8888888P'     8 888888888888       8 8888       .8'       `8. `88888.    `Y88888P'         8 8888       8 8888        8 \n\n"

func main() {
	fmt.Printf("\033[0;32m%s\033[0m", logo)
	log.Println("Starting...")

	ctx := context.Background()

	log.Printf("VERSION : %s", version)

	// Parsing configs from env file
	log.Println("Parsing configs...")

	conf, err := sys.NewConfigWithEnv()
	if err != nil {
		log.Fatalf("Failed to parse config : %v\n", err)
	}

	log.Printf("Starting %s environment\n", conf.Env)

	log.Println("Initializing logger...")
	logger, err := sys.Logger(conf.Env)
	if err != nil {
		log.Fatalf("Failed to create logger : %v", err)
	}

	log.Println("Initializing application...")
	if err := run(ctx, logger, conf); err != nil {
		logger.Fatalf("err : %v", err)
	}

}

func run(ctx context.Context, log *zap.SugaredLogger, conf *sys.Config) error {
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	db, err := sys.Postgres(ctx, conf.Postgres.DSN)
	if err != nil {
		return err
	}

	app, err := handlers.API(shutdown, log, db)
	if err != nil {
		return err
	}

	log.Info("Starting server...")
	server := http.Server{
		Addr:    conf.Addr,
		Handler: app,
	}

	serverErrChan := make(chan error, 1)

	go func() {
		serverErrChan <- server.ListenAndServe()
	}()

	log.Infof("Started the server on %s", conf.Addr)

	select {
	case <-serverErrChan:
		return err
	case <-shutdown:
		if err := server.Shutdown(context.Background()); err != nil {
			server.Close()
			return err
		}
		log.Info("server was gracefully stopped")
	}

	return nil
}

func hello() {
	repo := user.Repository
	service := user.Service()
}
