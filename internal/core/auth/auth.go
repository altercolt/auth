package auth

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"time"
)

const (
	AccessTokenExpiration  = time.Minute * 15
	RefreshTokenExpiration = time.Hour * 24 * 7
)

var (
	ErrInvalidAccessToken  = errors.New("invalid access token")
	ErrInvalidRefreshToken = errors.New("invalid refresh token")
)

type Login struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

// Payload
// is the claims part of jwt token
// stored in context
type Payload struct {
	UserID int    `json:"sub"`
	Role   string `json:"role"`
	Exp    int64  `json:"exp"`
}

func NewPayload(userID int, role string) Payload {
	return Payload{
		UserID: userID,
		Role:   role,
		Exp:    time.Now().Add(AccessTokenExpiration).Unix(),
	}
}

// Valid
// TODO("FINISH")
func (p Payload) Valid() error {
	now := time.Now()
	if now.After(time.Unix(p.Exp, 0)) {
		return ErrInvalidAccessToken
	}

	return nil
}

type RefreshPayload struct {
	UserID int   `json:"refresh_sub"`
	Exp    int64 `json:"refresh_exp"`
}

func NewRefreshPayload(userID int) RefreshPayload {
	return RefreshPayload{
		UserID: userID,
		Exp:    time.Now().Add(RefreshTokenExpiration).Unix(),
	}
}

func (p RefreshPayload) Valid() error {
	now := time.Now()
	if now.After(time.Unix(p.Exp, 0)) {
		return ErrInvalidAccessToken
	}

	return nil
}

// RefreshToken
// is model stored in the database
type RefreshToken struct {
	ID             uuid.UUID `json:"id"`
	UserID         int       `json:"user_id"`
	RefreshToken   string    `json:"refresh_token"`
	ExpirationTime time.Time `json:"expiration_time"`
}

// TokenPair
// is the structure that is returned to the user as JSON
type TokenPair struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	AccessExp    int64  `json:"access_token_exp"`
	RefreshExp   int64  `json:"refresh_token_exp"`
}
