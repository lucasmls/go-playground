package main

import "net/http"

import "log"

func main() {
	http.HandleFunc("/users/", UsersHandler)
	log.Println("Running...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
