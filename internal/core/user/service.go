package user

import (
	"context"
)

type Service interface {
	GetOneByID(ctx context.Context, id int) (User, error)
	GetOneByUsername(ctx context.Context, username string) (User, error)
	GetOneByEmail(ctx context.Context, email string) (User, error)
	Update(ctx context.Context, id int, update Update) error
	Create(ctx context.Context, nu New) error
}
