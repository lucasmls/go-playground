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

func calcAverage(values ...float64) float64 {
	total := 0.0
	for _, num := range values {
		total += num
	}

	return total / float64(len(values))
}

func printApproved(approved ...string) {
	fmt.Println("Approved:")
	for i, person := range approved {
		fmt.Printf("(%d) %s \n", i+1, person)
	}
}

var multiply = func(n1, n2 int) int {
	return n1 * n2
}

// we could use a "uint", so that we would not need to
// verifiy if the number is below 0
func factorial(n int) (int, error) {
	switch {
	case n < 0:
		return -1, fmt.Errorf("%d is a invalid number", n)
	case n == 0:
		return 1, nil
	default:
		lastFactorial, _ := factorial(n - 1)
		return n * lastFactorial, nil
	}
}

func main() {
	result := sum(10, 10)
	printResult(result)

	name, surname := getNameAndSurname("Lucas Mendes")
	fmt.Println(name, surname)

	fmt.Println(multiply(10, 23))

	executedMultiplication := exec(multiply, 2, 4)
	fmt.Println(executedMultiplication)

	average := calcAverage(2.34, 9.0, 8.3)
	fmt.Println(average)

	approved := []string{"Lucas", "Laisla"}
	printApproved(approved...)

	factorialOf5, err := factorial(5)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Factorial of 5 %d", factorialOf5)
}
