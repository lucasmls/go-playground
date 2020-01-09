package main

import "fmt"

func main() {
	i := 1

	// Creating a pointer, with nil associated to it
	var p *int = nil

	// Associating the memory local to my pointer (&)
	p = &i

	// In go, we can't use operators in a pointer directly,
	// first we need to access it's value to do so

	fmt.Println("variable i value:", i)
	fmt.Println("variable i memory reference:", &i)

	fmt.Println("pointer p value:", *p) // Accessing the pointer value, instead of its memory local (*)
	fmt.Println("pointer p memory reference:", p)

	fmt.Println("====")

	*p++
	i++

	fmt.Println("pointer p value:", *p)
	fmt.Println("pointer p memory reference:", p)

	fmt.Println("variable i value:", i)
	fmt.Println("variable i memory reference:", &i)
}
