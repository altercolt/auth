package main

import (
	"auth/cmd/api/handlers"
	"auth/internal/sys"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
)

var logo = "                                                                                                                                               \n    ,o888888o.        ,o888888o.     8 8888         8888888 8888888888          .8.          8 8888      88 8888888 8888888888 8 8888        8 \n   8888     `88.   . 8888     `88.   8 8888               8 8888               .888.         8 8888      88       8 8888       8 8888        8 \n,8 8888       `8. ,8 8888       `8b  8 8888               8 8888              :88888.        8 8888      88       8 8888       8 8888        8 \n88 8888           88 8888        `8b 8 8888               8 8888             . `88888.       8 8888      88       8 8888       8 8888        8 \n88 8888           88 8888         88 8 8888               8 8888            .8. `88888.      8 8888      88       8 8888       8 8888        8 \n88 8888           88 8888         88 8 8888               8 8888           .8`8. `88888.     8 8888      88       8 8888       8 8888        8 \n88 8888           88 8888        ,8P 8 8888               8 8888          .8' `8. `88888.    8 8888      88       8 8888       8 8888888888888 \n`8 8888       .8' `8 8888       ,8P  8 8888               8 8888         .8'   `8. `88888.   ` 8888     ,8P       8 8888       8 8888        8 \n   8888     ,88'   ` 8888     ,88'   8 8888               8 8888        .888888888. `88888.    8888   ,d8P        8 8888       8 8888        8 \n    `8888888P'        `8888888P'     8 888888888888       8 8888       .8'       `8. `88888.    `Y88888P'         8 8888       8 8888        8 \n\n"
var version = "SNAPSHOT 0.0.1"

type Config struct {
	Env         string
	Version     string
	Addr        string
	PostgresDSN string
}

func main() {
	fmt.Fprintf(os.Stdout, "\033[0;32m%s\033[0m", logo)

	log.Println("Starting...")
	log.Printf("VERSION : %s", version)

	// Parsing configs from env file
	log.Println("Parsing configs...")
	conf, err := sys.NewConfigWithEnv()
	if err != nil {
		log.Fatalf("Failed to parse config : %v\n", err)
	}

	log.Printf("Starting %s environment\n", conf.Env)

	log.Println("Initializing logger...")
	logger, err := Logger(conf.Env)
	if err != nil {
		log.Fatalf("Failed to create logger : %v", err)
	}

	log.Println("Initializing application...")
	shutdown := make(chan os.Signal)

	app, err := handlers.API(shutdown, logger)
	if err != nil {
		log.Fatalf("Cannot initialize application : %v", err)
	}

	log.Println("Starting server...")
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: app,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("err : %v", err)
	}

}

func Logger(env string) (*zap.SugaredLogger, error) {
	var log *zap.Logger
	var err error

	if env == "PROD" {
		log, err = zap.NewProduction()
	} else if env == "DEV" {
		log, err = zap.NewDevelopment()
	} else {
		return nil, errors.New("cannot determine env type")
	}

	if err != nil {
		return nil, nil
	}

	return log.Sugar(), nil
}
