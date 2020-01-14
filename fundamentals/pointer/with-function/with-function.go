package main

import "fmt"

func fn1(n int) {
	n++
}

// receives a int pointer
func fn2(n *int) {
	// acessing the value itself so that we can increment
	*n++
}

func main() {
	n := 1

	fn1(n) // passing the variable by its value
	fmt.Println(n)

	fn2(&n) // passing the variable by its reference
	fmt.Println(n)
}
