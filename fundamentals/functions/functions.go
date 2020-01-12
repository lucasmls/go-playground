package main

import (
	"fmt"
	"strings"
)

func sum(n1 int, n2 int) int {
	return n1 + n2
}

// Function with named return
// we don't need to add the variables in return because
// it was already specified in the function declaration
func getNameAndSurname(name string) (firstName, surname string) {
	splittedName := strings.Split(name, " ")
	firstName = splittedName[0]
	surname = splittedName[1]
	return
}

func printResult(result int) {
	fmt.Println(result)
}

func exec(fn func(int, int) int, n1 int, n2 int) int {
	return fn(n1, n2)
}

func main() {
	result := sum(10, 10)
	printResult(result)

	name, surname := getNameAndSurname("Lucas Mendes")
	fmt.Println(name, surname)

	sum := func(n1, n2 int) int {
		return n1 + n2
	}

	fmt.Println(sum(10, 23))

	executedSum := exec(sum, 20, 213)
	fmt.Println(executedSum)
}
