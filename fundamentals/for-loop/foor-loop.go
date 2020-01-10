package main

import (
	"fmt"
	"time"
)

func main() {
	i := 1
	for i <= 10 {
		fmt.Print(i)
		i++
	}

	fmt.Print("//")

	for x := 0; x <= 20; x += 2 {
		fmt.Print(x)
	}

	for {
		fmt.Println("Forever...")
		time.Sleep(time.Second)
	}
}
