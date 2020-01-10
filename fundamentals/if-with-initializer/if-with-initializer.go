package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomNumber() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	return r.Intn(10)
}

func main() {
	// With this approach, the variable lucky number
	// is accessible only inside the if statement
	if luckyNumber := randomNumber(); luckyNumber > 5 {
		fmt.Println(luckyNumber)
		fmt.Println("You win!!")
	} else {
		fmt.Println(luckyNumber)
		fmt.Println("You lose :(")
	}
}
