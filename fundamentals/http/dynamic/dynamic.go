package main

import "net/http"

import "time"

import "fmt"

import "log"

func getHour(w http.ResponseWriter, r *http.Request) {
	hour := time.Now().Format("02/01/2006 03:04:05")
	fmt.Fprintf(w, "<h1>%s</h1>", hour)
}

func main() {
	http.HandleFunc("/hours", getHour)
	log.Println("Running on port 3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
