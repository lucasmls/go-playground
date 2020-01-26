package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s Service) pingEndpoint() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, MessageResponse{Message: "Pong!"})
	}
}
