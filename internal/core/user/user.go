package user

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	PassHash string `json:"-"`
}

// NewUser
// is used for creating new user record in the database
type NewUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// UpdateUser
// not sure about internals yet, but it is used for updating user data
type UpdateUser struct {
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
	Role     *string
	PassHash *string
}
