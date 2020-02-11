package user

import (
	"fmt"
	"log"

	"github.com/lucasmls/todo/domain"
)

// ServiceInput ..
type ServiceInput struct {
	Repository domain.UsersRepository
}

// Service ...
type Service struct {
	in ServiceInput
}

// NewService ...
func NewService(in ServiceInput) (*Service, error) {
	if in.Repository == nil {
		return nil, fmt.Errorf("The users repository is required")
	}

	return &Service{in: in}, nil
}

// List ...
func (s Service) List() ([]*domain.User, error) {
	users, err := s.in.Repository.Find()
	if err != nil {
		log.Panic(err)
	}

	return users, nil
}

// Register ...
func (s Service) Register(userDto domain.User) (*domain.User, error) {
	user, err := s.in.Repository.Save(userDto)
	if err != nil {
		log.Panic(err)
		return nil, fmt.Errorf("Failed to register the user")
	}

	return user, nil
}
