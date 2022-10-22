package auth

import (
	"github.com/google/uuid"
	"time"
)

const (
	AccessTokenExpiration  = time.Minute * 15
	RefreshTokenExpiration = time.Hour * 24 * 7
)

// Payload
// is the claims part of jwt token
// stored in context
type Payload struct {
	UserID int    `json:"user_id"`
	Role   string `json:"role"`
	Exp    int64  `json:"exp"`
}

// Valid
// TODO("FINISH")
func (p Payload) Valid() error {
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
