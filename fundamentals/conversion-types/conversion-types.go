package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2

	fmt.Println(x / float64(y))

	result := 6.9
	finalResult := int(result) // will become 6
	fmt.Println(finalResult)

	// Take care...
	// In this string conversion, we'll be printing the correspodent unicode character
	fmt.Println("Printing test: " + string(97))

	// Int to string
	fmt.Println("Conversion test: " + strconv.Itoa(97))
	fmt.Println("Conversion test: " + fmt.Sprint(97))

	// String to int
	num, _ := strconv.Atoi("123")
	fmt.Println(num - 23)

	// String to boolean
	b, _ := strconv.ParseBool("true")
	if b {
		fmt.Printf("It's true!")
	}
}
