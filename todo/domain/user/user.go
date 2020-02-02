package user

import (
	"fmt"
	"log"

	"github.com/lucasmls/todo/domain"
	"github.com/lucasmls/todo/infra"
)

// ServiceInput ..
type ServiceInput struct {
	Postgres infra.PostgresProvider
}

// Service ...
type Service struct {
	in ServiceInput
}

// NewService ...
func NewService(in ServiceInput) *Service {
	return &Service{in: in}
}

// List ...
func (s Service) List() ([]*domain.User, string) {
	fmt.Println("List was called!")
	users, err := s.in.Postgres.GetUsers()
	if err != "" {
		log.Panic(err)
	}

	return users, ""
}

// Register ...
func (s Service) Register(userDto domain.User) (string, string) {
	fmt.Println("Register was called!")
	fmt.Println(userDto)

	successMessage, err := s.in.Postgres.SaveUser(userDto)
	if err != "" {
		log.Panic(err)
	}

	return successMessage, ""
}
