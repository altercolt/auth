package sys

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	Env      string
	Addr     string
	Postgres PostgresConfig
}

type PostgresConfig struct {
	DSN string
}

func NewConfigWithEnv() (*Config, error) {
	if err := godotenv.Load("local.env"); err != nil {
		return nil, err
	}

	conf := Config{
		Env:  getEnv("ENV"),
		Addr: getEnv("ADDR"),
		Postgres: PostgresConfig{
			DSN: getEnv("POSTGRES_DSN"),
		},
	}

	return &conf, nil
}

func getEnv(key string) string {
	val := os.Getenv(key)
	if val == "" {
		panic(fmt.Sprintf("[KEY : %s] IS EMPTY", key))
	}
	return val
}
