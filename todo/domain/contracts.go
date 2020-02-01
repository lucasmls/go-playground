package domain

// UsersProvider ...
type UsersProvider interface {
	Register() ([]*User, string)
}
