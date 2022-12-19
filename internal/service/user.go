package service

import (
	"auth/internal/core/user"
	"context"
	"github.com/google/uuid"
)

type UserService struct {
	userRepo user.Repository
}

func NewUserService() user.Service {
	return UserService{}
}

func (u UserService) Create(ctx context.Context, nu user.New) error {
	return nil
}

func (u UserService) Update(ctx context.Context, upd user.Update) error {
	return nil
}

func (u UserService) Delete(ctx context.Context, id uuid.UUID) error {
	return nil
}

func (u UserService) FetchByUsernames(ctx context.Context, usernames []string) ([]user.User, error) {
	users, err := u.userRepo.Fetch(ctx, user.Filter{Usernames: usernames})
	if err != nil {

	}

	return nil, nil
}

func (u UserService) FetchByIDs(ctx context.Context, ids []string) ([]user.User, error) {

	return nil, nil
}

func (u UserService) FetchOneByID(ctx context.Context, id string) (user.User, error) {
	return user.User{}, nil

}

func (u UserService) FetchOneByUsername(ctx context.Context, username string) (user.User, error) {
	return user.User{}, nil
}

func (u UserService) FetchOneByEmail(ctx context.Context, email string) (user.User, error) {
	return user.User{}, nil
}
