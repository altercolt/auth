package service

import (
	"auth/internal/core/user"
	"context"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	log      *zap.SugaredLogger
	userRepo user.Repository
}

func NewUserService(log *zap.SugaredLogger, userRepo user.Repository) user.Service {
	return UserService{
		log:      log,
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

func (u UserService) Update(ctx context.Context, id int, update user.Update) error {
	if err := update.Validate(); err != nil {
		u.log.Errorw("error when updating user", "err", err, "id", id)
		return fmt.Errorf("userService.Update(), err : %w", err)
	}

	//TODO: KAFKA SHIT

	usr, err := u.GetOneByID(ctx, id)
	if err != nil {
		u.log.Errorw("error when updating user", "err", err, "id", id)
		return fmt.Errorf("userService.Update, err : %w", err)
	}

	if err = bcrypt.CompareHashAndPassword([]byte(usr.PassHash), []byte(*update.Password)); err != nil {
		u.log.Errorw("error when updating user", "err", err, "id", id)
		return fmt.Errorf("userService.Update, err : %w", err)
	}

	if update.Username != nil {
		usr, err := u.GetOneByUsername(ctx, *update.Username)
		if err != nil {
			u.log.Errorw("error when updating user", "err", err, "id", id)
			return fmt.Errorf("userService.Update, err : %w", err)
		}

		if usr != (user.User{}) {
			return errors.New("user with such username already exists")
		}
	}

	if update.Email != nil {
		usr, err := u.GetOneByEmail(ctx, *update.Email)
		if err != nil {
			u.log.Errorw("error when updating user", "err", err, "id", id)
			return fmt.Errorf("userService.Update, err : %w", err)
		}

		if usr != (user.User{}) {
			return errors.New("user with such email already exists")
		}
	}

	if update.NewPassword != nil {
		passHash, err := bcrypt.GenerateFromPassword([]byte(*update.NewPassword), bcrypt.DefaultCost)
		if err != nil {
			u.log.Errorw("error when updating user", "err", err, "id", id)
			return fmt.Errorf("userService.Update, err : %w", err)
		}
		update.NewPassword = toPtr(string(passHash))
	}

	if err = u.userRepo.Update(ctx, user.Model{
		ID:       toPtr(id),
		Email:    update.Email,
		Username: update.Username,
		PassHash: update.NewPassword,
	}); err != nil {
		u.log.Errorw("error when updating user", "err", err, "id", id)
		return fmt.Errorf("userService.Update, err : %w", err)
	}

	return nil
}

func (u UserService) Create(ctx context.Context, nu user.New) error {
	if err := nu.Validate(); err != nil {
		u.log.Errorw("error when creating new user", "err", err, "NewUser", nu)
		return fmt.Errorf("userService.Create, err : %w", err)
	}

	if nu.Username != "" {
		usr, err := u.GetOneByUsername(ctx, nu.Username)
		if err != nil {
			u.log.Errorw("error when creating new user", "err", err, "NewUser", nu)
			return fmt.Errorf("userService.Create, err : %w", err)
		}

		if usr != (user.User{}) {
			return errors.New("user with such username already exists")
		}
	}

	if nu.Email != "" {
		usr, err := u.GetOneByEmail(ctx, nu.Email)
		if err != nil {
			u.log.Errorw("error when creating new user", "err", err, "NewUser", nu)
			return fmt.Errorf("userService.Create, err : %w", err)
		}

		if usr != (user.User{}) {
			return errors.New("user with such email already exists")
		}
	}

	//TODO : KAFKA SHIT

	bcryptedPass, err := bcrypt.GenerateFromPassword([]byte(nu.Password), bcrypt.DefaultCost)
	if err != nil {
		u.log.Errorw("error when creating new user", "err", err, "NewUser", nu)
		return fmt.Errorf("userService.Create, err : %w", err)
	}

	passHash := string(bcryptedPass)

	if err := u.userRepo.Create(ctx, user.Model{
		Email:    &nu.Email,
		Username: &nu.Username,
		PassHash: &passHash,
	}); err != nil {
		u.log.Errorw("error when creating new user", "err", err, "NewUser", nu)
		return fmt.Errorf("userService.Create, err : %w", err)
	}

	return nil
}
