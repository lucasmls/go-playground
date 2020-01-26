package main

import "github.com/lucasmls/todo/domain/server"

import "fmt"

func main() {
	srvr := server.NewService(server.ServiceInput{})

	errCh := srvr.Run()
	for err := range errCh {
		fmt.Println(err)
	}
}
