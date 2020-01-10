package main

import (
	"fmt"
	"time"
)

func discoverParameterType(typeToBeDiscovered interface{}) string {
	switch typeToBeDiscovered.(type) {
	case int:
		return "It's a integer"

	case float32, float64:
		return "It's a float"

	case string:
		return "It's a string"

	case func():
		return "It's a function"

	default:
		return "I don't know this type, sorry"
	}
}

func main() {
	fmt.Println(discoverParameterType(2.3))
	fmt.Println(discoverParameterType(2))
	fmt.Println(discoverParameterType("Lucas"))
	fmt.Println(discoverParameterType(func() {}))
	fmt.Println(discoverParameterType(time.Now()))
}
