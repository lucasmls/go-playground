package main

import (
	"fmt"
	"time"
)

func speak(person string) <-chan string {
	ch := make(chan string)

	go func() {
		for i := 0; i < 3; i++ {
			ch <- fmt.Sprintf("%s says: %d", person, i)
			time.Sleep(time.Second)
		}
	}()

	return ch
}

func join(input1, input2 <-chan string) <-chan string {
	ch := make(chan string)

	go func() {
		for {
			select {
			case value := <-input1:
				ch <- value
			case value := <-input2:
				ch <- value
			}
		}
	}()

	return ch
}

func main() {
	ch := join(speak("Lucas"), speak("Laisla"))

	fmt.Println(<-ch, <-ch)
	fmt.Println(<-ch, <-ch)
	fmt.Println(<-ch, <-ch)
}
