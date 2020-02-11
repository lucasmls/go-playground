package domain

// UsersProvider ...
type UsersProvider interface {
	List() ([]*User, error)
	Register(User) (*User, error)
}

// UsersRepository ...
type UsersRepository interface {
	Find() ([]*User, error)
	Save(User) (*User, error)
}
