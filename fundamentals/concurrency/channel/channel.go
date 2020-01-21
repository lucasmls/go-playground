package main

import "fmt"

func main() {
	ch := make(chan int, 1)

	ch <- 1 // Sending data to the channel
	<-ch    // Reading data from channel

	ch <- 2
	fmt.Println(<-ch)
}
