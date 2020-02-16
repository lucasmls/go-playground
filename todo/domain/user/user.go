package user

import (
	"fmt"
	"log"

	"github.com/lucasmls/todo/domain"
	"github.com/lucasmls/todo/infra"
	"golang.org/x/crypto/bcrypt"
)

// ServiceInput ..
type ServiceInput struct {
	Repository  domain.UsersRepository
	JwtProvider infra.JwtProvider
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

// Login ...
func (s Service) Login(loginDto domain.LoginInput) (string, error) {
	user, usrErr := s.in.Repository.FindByEmail(loginDto.Email)
	if usrErr != nil {
		log.Panic(usrErr)
		return "", fmt.Errorf("Failed to find the user by its email")
	}

	compareErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginDto.Password))
	if compareErr != nil {
		return "", fmt.Errorf("Wrong password")
	}

	jwtToken, tknErr := s.in.JwtProvider.GenerateJWT(user.ID)
	if tknErr != nil {
		fmt.Println(tknErr)
		return "", fmt.Errorf("Failed to generate the JWT")
	}

	return jwtToken, nil
}
