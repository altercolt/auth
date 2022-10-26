package main

import (
	"auth/cmd/api/handlers"
	"auth/internal/sys"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var logo = "                                                                                                                                               \n    ,o888888o.        ,o888888o.     8 8888         8888888 8888888888          .8.          8 8888      88 8888888 8888888888 8 8888        8 \n   8888     `88.   . 8888     `88.   8 8888               8 8888               .888.         8 8888      88       8 8888       8 8888        8 \n,8 8888       `8. ,8 8888       `8b  8 8888               8 8888              :88888.        8 8888      88       8 8888       8 8888        8 \n88 8888           88 8888        `8b 8 8888               8 8888             . `88888.       8 8888      88       8 8888       8 8888        8 \n88 8888           88 8888         88 8 8888               8 8888            .8. `88888.      8 8888      88       8 8888       8 8888        8 \n88 8888           88 8888         88 8 8888               8 8888           .8`8. `88888.     8 8888      88       8 8888       8 8888        8 \n88 8888           88 8888        ,8P 8 8888               8 8888          .8' `8. `88888.    8 8888      88       8 8888       8 8888888888888 \n`8 8888       .8' `8 8888       ,8P  8 8888               8 8888         .8'   `8. `88888.   ` 8888     ,8P       8 8888       8 8888        8 \n   8888     ,88'   ` 8888     ,88'   8 8888               8 8888        .888888888. `88888.    8888   ,d8P        8 8888       8 8888        8 \n    `8888888P'        `8888888P'     8 888888888888       8 8888       .8'       `8. `88888.    `Y88888P'         8 8888       8 8888        8 \n\n"
var version = "SNAPSHOT 0.0.1"

func main() {
	fmt.Printf("\033[0;32m%s\033[0m", logo)

	log.Println("Starting...")
	log.Printf("VERSION : %s", version)
	ctx := context.Background()
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

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	db, err := sys.Postgres(ctx, conf.Postgres.DSN)
	if err != nil {
		log.Fatalf("Cannot connect to the database : %v", err)
	}

	//TODO:
	app, err := handlers.API(shutdown, logger, db)
	if err != nil {
		log.Fatalf("Cannot initialize application : %v", err)
	}

	log.Println("Starting server...")
	server := http.Server{
		Addr:    conf.Addr,
		Handler: app,
	}

	serverErrChan := make(chan error, 1)

	go func() {
		serverErrChan <- server.ListenAndServe()
	}()

	log.Printf("Started the server on %s", conf.Addr)

	select {
	case <-serverErrChan:
		log.Fatalf("server error: %v", err)
	case <-shutdown:
		if err := server.Shutdown(context.Background()); err != nil {
			log.Fatalf("Error when server.Shutdown() : %v\n", err)
			server.Close()
			return
		}
		log.Println("Server was gracefully stopped")
	}
}
