package domain

// UsersProvider ...
type UsersProvider interface {
	List() ([]*User, error)
	Register(user User) (*User, error)
	Login(loginInput LoginInput) (string, error)
}

// UsersRepository ...
type UsersRepository interface {
	Find() ([]*User, error)
	FindOne(userID int64) (*User, error)
	FindByEmail(email string) (*User, error)
	Save(User) (*User, error)
}
