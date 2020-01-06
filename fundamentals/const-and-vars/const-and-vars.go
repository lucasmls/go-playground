package main

import (
	"fmt"
	"math"
)

func main() {
	const PI float64 = 3.14
	var radius float64 = 3.2

	// Short syntax
	area := PI * math.Pow(radius, 2)
	fmt.Println("The circumference area is:", area)

	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)

	fmt.Println(a, b, c, d)

	var e, f bool = true, false
	fmt.Println(e, f)

	g, h, i := 2, false, "hey"

	fmt.Println(g, h, i)
}
