package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucasmls/todo/domain"
)

func (s Service) pingEndpoint() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, MessageResponse{Message: "Pong!"})
	}
}

func (s Service) listUsersEndpoint() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {

		payload := domain.User{}
		if err := ctx.ShouldBindJSON(&payload); err != nil {
			fmt.Println("Error when getting the payload")
		}

		users, err := s.in.UsersProvider.Register()
		if err != "" {
			log.Panic(err)
		}

		ctx.JSON(http.StatusCreated, users)
	}
}
