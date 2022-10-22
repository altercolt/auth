package user

type Repository interface {
}

type Filter struct {
	IDs       []int
	Emails    []string
	Usernames []string
}
