package user

import (
	"context"
)

type Repository interface {
	Create(ctx context.Context, m Model) error
	Delete(ctx context.Context, id int) error
	Update(ctx context.Context, m Model) error
	Fetch(ctx context.Context, filter Filter) ([]User, error)
	FetchOne(ctx context.Context, filter Filter) (User, error)
}

type Filter struct {
	IDs       []int
	Emails    []string
	Usernames []string
}
