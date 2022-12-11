package service

import (
	"auth/internal/core/auth"
	"auth/internal/core/user"
	"context"
)

type AuthService struct {
	tokenRepo   auth.TokenRepository
	userService user.Service
}

func NewAuthService(tokenRepo auth.TokenRepository, userService user.Service) auth.Service {
	return &AuthService{
		tokenRepo:   tokenRepo,
		userService: userService,
	}
}

// Might be good, idk
/*func AuthServiceMiddleware() web.Middleware {
	m := func(handler web.Handler) web.Handler {
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
			err := handler(ctx, w, r)

		}

		return h
	}

	return m
}*/

func (a *AuthService) Login(ctx context.Context, login auth.Login) (auth.TokenPair, error) {
	a.userService
}

func (a *AuthService) Logout(ctx context.Context, refreshToken string) (auth.Payload, error) {

}

func (a *AuthService) RefreshAccess(ctx context.Context, refreshToken string) (auth.TokenPair, error) {
	//TODO implement me
	panic("implement me")
}

func (a *AuthService) ValidateAccess(ctx context.Context, accessToken string) (auth.Payload, error) {
	//TODO implement me
	panic("implement me")
}
