package user

import (
	"context"
	"github.com/google/uuid"
)

type Service interface {
	Create(ctx context.Context, nu New) error
	Update(ctx context.Context, upd Update) error
	Delete(ctx context.Context, id uuid.UUID) error

	FetchByUsernames(ctx context.Context, usernames []string)
	FetchByIDs(ctx context.Context, ids []string)
	FetchOneByID(ctx context.Context, id string)
	FetchOneByUsername(ctx context.Context, username string)
	FetchOneByEmail(ctx context.Context, email string)
}
