package main

import (
	"concurrency/html"
	"fmt"
)

func forward(originCh <-chan string, destinyCh chan string) {
	for {
		destinyCh <- <-originCh
	}
}

func join(input1, input2 <-chan string) <-chan string {
	ch := make(chan string)
	go forward(input1, ch)
	go forward(input2, ch)

	return ch
}

func main() {
	ch := join(
		html.GetTitles("https://www.google.com", "https://www.twitter.com"),
		html.GetTitles("https://www.github.com", "https://www.youtube.com"),
	)

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
