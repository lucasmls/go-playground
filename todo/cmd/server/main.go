package main

import (
	"fmt"
	"github.com/lucasmls/todo/domain/server"
	"github.com/lucasmls/todo/domain/user"
	"github.com/lucasmls/todo/infra/postgres"
)

func main() {

	postgres := postgres.NewClient(postgres.ClientInput{
		Endpoint: "user=postgres port=1234 password=postgres dbname=go_todo sslmode=disable",
	})

	user := user.NewService(user.ServiceInput{
		Postgres: postgres,
	})

	srvr := server.NewService(server.ServiceInput{
		UsersProvider: user,
	})

	errCh := srvr.Run()
	for err := range errCh {
		fmt.Println(err)
	}
}
