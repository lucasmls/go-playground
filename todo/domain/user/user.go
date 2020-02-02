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
