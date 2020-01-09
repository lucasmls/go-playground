package main

import (
	"fmt"
	"math"
)

func main() {
	a := 3
	b := 2

	fmt.Println("Sum:", a+b)
	fmt.Println("Subtraction:", a-b)
	fmt.Println("Division:", a/b)
	fmt.Println("Multiply:", a*b)
	fmt.Println("Mod:", a%b)

	// Bitwise
	fmt.Println("And:", a&b) // 11 & 10 = 10
	fmt.Println("Or:", a|b)
	fmt.Println("Xor:", a^b)

	c := 3.0
	d := 2.0

	fmt.Println("Bigger:", math.Max(float64(c), float64(d)))
	fmt.Println("Smaller:", math.Min(c, d))
	fmt.Println("Pow:", math.Pow(c, d))
}
