package auth

import "github.com/google/uuid"

type TokenRepository interface {
}

type Filter struct {
	IDs    []uuid.UUID
	Users  []int
	Tokens []string
}
