package auth

import "context"

type Service interface {
	Login(ctx context.Context, login Login) (TokenPair, error)
	ValidateAccess(ctx context.Context, accessToken string) (Payload, error)
}

type TokenService interface {
	Create(ctx context.Context)
	Delete(ctx context.Context)
}
