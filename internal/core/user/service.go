package user

import (
	"auth/internal/core/auth"
	"context"
	"github.com/google/uuid"
)

type Service interface {
	Create(ctx context.Context, nu New) error
	Update(ctx context.Context, payload auth.Payload, upd Update) error
	Delete(ctx context.Context, id uuid.UUID) error

	FetchByUsernames(ctx context.Context, usernames []string) ([]User, error)
	FetchByIDs(ctx context.Context, ids []string) ([]User, error)
	FetchOneByID(ctx context.Context, id string) (User, error)
	FetchOneByUsername(ctx context.Context, username string) (User, error)
	FetchOneByEmail(ctx context.Context, email string) (User, error)
}
