package user

import (
	"context"
	"github.com/google/uuid"
)

type Repository interface {
	Create(ctx context.Context, m Model) error
	Update(ctx context.Context, m Model) error
	Delete(ctx context.Context, id uuid.UUID) error

	FetchOne(ctx context.Context, f SingleFilter) (User, error)
	Fetch(ctx context.Context, f Filter) ([]User, error):
}

type Filter struct {
	IDs   []uuid.UUID
	Email []string
}

type SingleFilter struct {
	ID       uuid.UUID
	Username string
	Email    string
}
