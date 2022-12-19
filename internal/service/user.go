package service

import (
	"auth/internal/core/auth"
	errors2 "auth/internal/core/errors"
	"auth/internal/core/user"
	"auth/internal/repository"
	"context"
	"errors"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepo user.Repository
}

func NewUserService(userRepo user.Repository) user.Service {
	return UserService{
		userRepo: userRepo,
	}
}

func (u UserService) Create(ctx context.Context, nu user.New) error {
	// validating newUser
	if err := nu.Validate(); err != nil {
		return err
	}

	// generating password hash
	passHashBytes, err := bcrypt.GenerateFromPassword([]byte(nu.Password), bcrypt.DefaultCost)
	passHash := string(passHashBytes)
	if err != nil {
		return err
	}

	model := user.Model{
		Email:    &nu.Email,
		Username: &nu.Username,
		PassHash: &passHash,
	}

	if err := u.userRepo.Create(ctx, model); err != nil {
		if errors.Is(err, repository.UniqueError{}) {
			//TODO : reconsider field
			err = NewDuplicateEntryError("user already exists", "", err)
		}
		return err
	}

	return nil
}

func (u UserService) Update(ctx context.Context, payload auth.Payload, upd user.Update) error {
	// validating newUser
	if err := upd.Validate(); err != nil {
		return err
	}

	usr, err := u.FetchOneByID(ctx, payload.UserID)
	if err != nil {
		return err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(usr.PassHash), []byte(*upd.Password)); err != nil {
		return ErrInvalidCredentials
	}

	model := user.Model{}
	// generating password hash
	if upd.NewPassword != nil {
		passHashBytes, err := bcrypt.GenerateFromPassword([]byte(*upd.NewPassword), bcrypt.DefaultCost)
		passHash := string(passHashBytes)
		if err != nil {
			// ???
			return err
		}
		model.PassHash = &passHash
	}

	model.Email = upd.Email
	model.Username = upd.Username

	if err := u.userRepo.Create(ctx, model); err != nil {
		if errors.Is(err, repository.UniqueError{}) {
			//TODO : reconsider field
			err = NewDuplicateEntryError("user already exists", "", err)
		}
		return err
	}

	return nil
}

// Delete
// IDK about this one
// TODO : ask for user's password
func (u UserService) Delete(ctx context.Context, id uuid.UUID) error {
	return nil
}

func (u UserService) FetchByUsernames(ctx context.Context, usernames []string) ([]user.User, error) {
	users, err := u.userRepo.Fetch(ctx, user.Filter{Usernames: usernames})
	if err != nil {
		var e *repository.NotFoundError
		if errors.Is(err, e) {
			return users, nil
		}
		return nil, err
	}

	return users, nil
}

func (u UserService) FetchByIDs(ctx context.Context, ids []string) ([]user.User, error) {
	users, err := u.userRepo.Fetch(ctx, user.Filter{IDs: parseIDs(ids)})
	if err != nil {
		var e *repository.NotFoundError
		if errors.Is(err, e) {
			return users, nil
		}
		return nil, err
	}

	return users, nil
}

func (u UserService) FetchOneByID(ctx context.Context, id string) (user.User, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		validationError := errors2.NewValidationError()
		validationError.Append("id", "invalid uuid")
		return user.User{}, validationError
	}

	usr, err := u.userRepo.FetchOne(ctx, user.SingleFilter{ID: uid})
	if err != nil {
		var e *repository.NotFoundError
		if errors.Is(err, e) {
			// might be a bad decision
			return usr, nil
		}
		return user.User{}, err
	}

	return usr, nil
}

func (u UserService) FetchOneByUsername(ctx context.Context, username string) (user.User, error) {
	usr, err := u.userRepo.FetchOne(ctx, user.SingleFilter{Username: username})
	if err != nil {
		var e *repository.NotFoundError
		if errors.Is(err, e) {
			// might be a bad decision
			return usr, nil
		}
		return user.User{}, err
	}

	return usr, nil
}

func (u UserService) FetchOneByEmail(ctx context.Context, email string) (user.User, error) {
	usr, err := u.userRepo.FetchOne(ctx, user.SingleFilter{Email: email})
	if err != nil {
		var e *repository.NotFoundError
		if errors.Is(err, e) {
			// might be a bad decision
			return usr, nil
		}
		return user.User{}, err
	}

	return usr, nil
}

func parseIDs(ids []string) []uuid.UUID {
	var res []uuid.UUID
	for i := range ids {
		id, err := uuid.Parse(ids[i])
		if err != nil {
			continue
		}
		res = append(res, id)
	}

	return res
}
