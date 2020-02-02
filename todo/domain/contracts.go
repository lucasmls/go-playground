package domain

// UsersProvider ...
type UsersProvider interface {
	List() ([]*User, string)
}
