package user

import (
	"fmt"
)

// ServiceInput ..
type ServiceInput struct{}

// Service ...
type Service struct {
	in ServiceInput
}

// NewService ...
func NewService(in ServiceInput) *Service {
	return &Service{in: in}
}

// Register ...
func (s Service) Register() {
	fmt.Println("Register was called!")
}
