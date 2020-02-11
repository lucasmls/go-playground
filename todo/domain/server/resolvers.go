package server

import (
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
		users, err := s.in.UsersProvider.List()
		if err != nil {
			log.Panic(err)
		}

		ctx.JSON(http.StatusOK, users)
	}
}

func (s Service) registerUsersEndpoint() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		registerPayload := domain.User{}
		if err := ctx.ShouldBindJSON(&registerPayload); err != nil {
			log.Panic("Failed to parse register payload.")
		}

		registerSuccessMessage, err := s.in.UsersProvider.Register(registerPayload)
		if err != nil {
			log.Panic(err)
		}

		ctx.JSON(http.StatusCreated, registerSuccessMessage)

	}
}
