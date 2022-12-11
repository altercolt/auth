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
	if err := nu.Validate(); err != nil {
		return err
	}

	u.userRepo.Count(ctx, user.SingleFilter{Email: nu.Email})

}

func (u UserService) Update(ctx context.Context, upd user.Update) error {
	//TODO implement me
	panic("implement me")
}

func (u UserService) Delete(ctx context.Context, id uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (u UserService) FetchByUsernames(ctx context.Context, usernames []string) {
	//TODO implement me
	panic("implement me")
}

func (u UserService) FetchByIDs(ctx context.Context, ids []string) {
	//TODO implement me
	panic("implement me")
}

func (u UserService) FetchOneByID(ctx context.Context, id string) {
	//TODO implement me
	panic("implement me")
}

func (u UserService) FetchOneByUsername(ctx context.Context, username string) {
	//TODO implement me
	panic("implement me")
}

func (u UserService) FetchOneByEmail(ctx context.Context, email string) {
	//TODO implement me
	panic("implement me")
}
