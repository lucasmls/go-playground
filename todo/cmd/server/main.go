package main

import (
	"fmt"
	"github.com/lucasmls/todo/domain/server"
	"github.com/lucasmls/todo/domain/user"
)

func main() {

	user := user.NewService(user.ServiceInput{})

	srvr := server.NewService(server.ServiceInput{
		UsersProvider: user,
	})

	errCh := srvr.Run()
	for err := range errCh {
		fmt.Println(err)
	}
}
