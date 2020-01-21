package main

import (
	"fmt"
	"time"
)

func speak(personName, phrase string, amount int) {
	for i := 0; i < amount; i++ {
		time.Sleep(time.Second)
		fmt.Printf("[%d] - %s: %s\n", i+1, personName, phrase)
	}
}

func main() {
	go speak("Lucas", "Hey!", 10)
	go speak("Laisla", "Ho!", 10)
}
