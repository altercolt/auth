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

func (a *AuthService) Login(ctx context.Context, login auth.Login) (auth.TokenPair, error) {
	return auth.TokenPair{}, nil
}

func (a *AuthService) Logout(ctx context.Context, refreshToken string) error {
	return nil
}

func (a *AuthService) RefreshAccess(ctx context.Context, refreshToken string) (auth.TokenPair, error) {
	return auth.TokenPair{}, nil
}

func (a *AuthService) ValidateAccess(ctx context.Context, accessToken string) (auth.Payload, error) {
	return auth.Payload{}, nil
}
