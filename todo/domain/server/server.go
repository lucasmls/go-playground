package server

import (
	"github.com/gin-gonic/gin"
	"github.com/lucasmls/todo/domain"
)

// ServiceInput ...
type ServiceInput struct {
	UsersProvider domain.UsersProvider
}

// Service ...
type Service struct {
	in    ServiceInput
	errCh chan string
}

// NewService ...
func NewService(in ServiceInput) *Service {
	// here will come the dependencies of the service
	return &Service{
		in:    in,
		errCh: make(chan string),
	}
}

// Engine ...
func (s Service) Engine() *gin.Engine {
	router := gin.Default()

	router.GET(
		"/ping",
		s.pingEndpoint(),
	)

	router.POST("/login", s.login())

	usersGroup := router.Group(
		"/user",
	)
	{
		usersGroup.GET("", s.listUsersEndpoint())
		usersGroup.POST("", s.registerUsersEndpoint())
	}

	return router
}

// Run ...
func (s Service) Run() <-chan string {

	go func() {
		if err := s.Engine().Run(); err != nil {
			s.errCh <- "Unexpected error occurred"
		}

		close(s.errCh)
	}()

	return s.errCh
}
