package auth

import (
	"context"
	"github.com/google/uuid"
)

type TokenRepository interface {
	Create(ctx context.Context, token RefreshToken) error
	Delete(ctx context.Context, id uuid.UUID) error
	Fetch(ctx context.Context, filter Filter) ([]RefreshToken, error)
	FetchOne(ctx context.Context, filter Filter) (RefreshToken, error)
}

type Filter struct {
	IDs    []uuid.UUID
	Users  []int
	Tokens []string
}
