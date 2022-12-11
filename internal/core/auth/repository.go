package auth

import (
	"context"
	"github.com/google/uuid"
)

type TokenRepository interface {
	Create(ctx context.Context, token RefreshToken) error
	DeleteByID(ctx context.Context, id uuid.UUID) error
	DeleteByRefreshTokenString(ctx context.Context, refreshToken string) error

	FetchOne(ctx context.Context, filter SingleFilter) (RefreshToken, error)
	Fetch(ctx context.Context, filter Filter) ([]RefreshToken, error)
}

type Filter struct {
	IDs    []uuid.UUID
	Users  []int
	Tokens []string
}

type SingleFilter struct {
	ID    uuid.UUID
	Token string
}
