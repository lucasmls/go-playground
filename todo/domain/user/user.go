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

// Register ...
func (s Service) Register() ([]*domain.User, string) {
	fmt.Println("Register was called!")
	users, err := s.in.Postgres.GetUsers()
	if err != "" {
		log.Panic(err)
	}

	fmt.Println(users)
	return users, ""
}
