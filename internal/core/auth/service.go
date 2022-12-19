package auth

import (
	"context"
)

type Service interface {
	Login(ctx context.Context, login Login) (TokenPair, error)
	Logout(ctx context.Context, refreshToken string) error
	RefreshAccess(ctx context.Context, refreshToken string) (TokenPair, error)

	// ValidateAccess
	// used in auth middleware
	// for access validation
	ValidateAccess(ctx context.Context, accessToken string) (Payload, error)
}
