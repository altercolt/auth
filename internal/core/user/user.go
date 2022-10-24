package user

import "auth/internal/core/role"

type User struct {
	ID       int       `json:"id"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Role     role.Role `json:"role"`
	PassHash string    `json:"-"`
}

// New
// is used for creating new user record in the database
type New struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Update
// not sure about internals yet, but it is used for updating user data
type Update struct {
	Username    *string `json:"username"`
	Email       *string `json:"email"`
	NewPassword *string `json:"new-password"`
	Password    *string `json:"password"`
}

// Model
// database model
type Model struct {
	ID       *int
	Email    *string
	Username *string
	Role     *role.Role
	PassHash *string
}
