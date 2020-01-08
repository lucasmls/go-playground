package main

import "fmt"

func sum(n1 int, n2 int) int {
	return n1 + n2
}

func printResult(result int) {
	fmt.Println(result)
}

func main() {
	result := sum(10, 10)
	printResult(result)
}
