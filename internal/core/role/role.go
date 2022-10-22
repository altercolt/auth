package role

type Role string

const (
	Admin Role = "ROLE_ADMIN"
	User  Role = "ROLE_USER"
)

type Model struct {
	ID int
}
