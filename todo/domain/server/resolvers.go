package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucasmls/todo/domain"
)

func (s Service) pingEndpoint() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, MessageResponse{Message: "Pong!"})
	}
}

func (s Service) registerUserEndpoint() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {

		payload := domain.User{}
		if err := ctx.ShouldBindJSON(&payload); err != nil {
			fmt.Println("Error when getting the payload")
		}

		s.in.UsersProvider.Register()

		ctx.JSON(http.StatusCreated, MessageResponse{Message: "User registered"})
	}
}
