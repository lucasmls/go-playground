package main

import "fmt"

func main() {
	numbers := [...]int{5, 4, 3, 2, 1}

	for i, number := range numbers {
		fmt.Printf("[%d] - %d\n", i, number)
	}

	for _, number := range numbers {
		fmt.Printf("%d\n", number)
	}
}
