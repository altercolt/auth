package service

import (
	"auth/internal/core/user"
	"context"
	"fmt"
	"go.uber.org/zap"
)

type UserService struct {
	log      *zap.SugaredLogger
	userRepo user.Repository
}

func NewUserService(log *zap.SugaredLogger, userRepo user.Repository) user.Service {
	return UserService{
		userRepo: userRepo,
	}
}

func (u UserService) GetOneByID(ctx context.Context, id int) (user.User, error) {
	usr, err := u.userRepo.FetchOne(ctx, user.Filter{
		IDs: []int{id},
	})

	if err != nil {
		u.log.Errorw("error when fetching user by id",
			"err", err, "id", id)
		return user.User{}, fmt.Errorf("userService.GetOneByID(), err : %w", err)
	}

	return usr, nil
}

func (u UserService) GetOneByUsername(ctx context.Context, username string) (user.User, error) {
	usr, err := u.userRepo.FetchOne(ctx, user.Filter{
		Usernames: []string{username},
	})

	if err != nil {
		u.log.Errorw("error when fetching user by username", "err", err, "username", username)
		return user.User{}, fmt.Errorf("userService.GetOneByUsername(), err : %w", err)
	}

	return usr, nil
}

func (u UserService) GetOneByEmail(ctx context.Context, email string) (user.User, error) {
	usr, err := u.userRepo.FetchOne(ctx, user.Filter{
		Emails: []string{email},
	})

	if err != nil {
		u.log.Errorw("error when fetching user by email", "err", err, "email", email)
		return user.User{}, fmt.Errorf("userService.GetOneByEmail(), err : %w", err)
	}

	return usr, nil
}

func (u UserService) Update(ctx context.Context, update user.Update) error {

}