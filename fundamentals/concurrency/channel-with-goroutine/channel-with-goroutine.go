package main

import (
	"fmt"
	"time"
)

func twoThreeFourFive(baseValue int, ch chan int) {
	time.Sleep(time.Second)
	ch <- 2 * baseValue

	time.Sleep(time.Second)
	ch <- 3 * baseValue

	time.Sleep(time.Second)
	ch <- 4 * baseValue
}

func main() {
	ch := make(chan int)

	go twoThreeFourFive(2, ch)

	v1, v2 := <-ch, <-ch

	fmt.Println(v1, v2)

	v3 := <-ch
	fmt.Println(v3)
}
