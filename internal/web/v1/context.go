package v1

import (
	"auth/internal/core/auth"
	"context"
	"github.com/pkg/errors"
)

func Payload(ctx context.Context) (auth.Payload, error) {
	p := ctx.Value("payload")
	if p == nil {
		return auth.Payload{}, errors.New("no data in payload")
	}

	payload, ok := p.(auth.Payload)
	if !ok {
		return auth.Payload{}, errors.New("invalid payload")
	}

	return payload, nil
}
