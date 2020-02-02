package domain

// UsersProvider ...
type UsersProvider interface {
	List() ([]*User, string)
	Register(User) (string, string)
}
