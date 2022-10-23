package sys

import (
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

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
