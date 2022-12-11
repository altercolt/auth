package user

import (
	"auth/internal/core/role"
	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `json:"id"`
	Email    string    `json:"email"`
	Username string    `json:"username"`
	Role     role.Role `json:"role"`
	PassHash string    `json:"-"`
}

type New struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Update struct {
	Email       *string `json:"email,omitempty"`
	Username    *string `json:"username,omitempty"`
	NewPassword *string `json:"new-password,omitempty"`
	Password    *string `json:"password,omitempty"`
}

type Model struct {
	ID       uuid.UUID
	Email    *string
	Username *string
	PassHash *string
}
