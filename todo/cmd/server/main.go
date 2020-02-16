package main

import (
	"fmt"
	"os"
	"time"

	"github.com/lucasmls/todo/domain/server"
	"github.com/lucasmls/todo/domain/user"
	userrepository "github.com/lucasmls/todo/domain/user_repository"
	"github.com/lucasmls/todo/infra/jwt"
	"github.com/lucasmls/todo/infra/postgres"
)

func main() {
	postgres := postgres.NewClient(postgres.ClientInput{
		Endpoint: "user=postgres port=1234 password=postgres dbname=go_todo sslmode=disable",
	})

	userRepository := userrepository.NewClient(userrepository.ClientInput{
		Pg: postgres,
	})

	jwt := jwt.NewClient(jwt.ClientInput{
		Secret: os.Getenv("JWT_SECRET"),
		TTL:    time.Now().Add(time.Minute * 30).Unix(),
	})

	user, _ := user.NewService(user.ServiceInput{
		Repository:  userRepository,
		JwtProvider: jwt,
	})

	srvr := server.NewService(server.ServiceInput{
		UsersProvider: user,
	})

	errCh := srvr.Run()
	for err := range errCh {
		fmt.Println(err)
	}
}
