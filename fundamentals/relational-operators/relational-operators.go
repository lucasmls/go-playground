package main

import (
	"fmt"
	"time"
)

// Person ...
type Person struct {
	Name    string
	Surname string
}

func main() {
	fmt.Println("Strings:", "Banana" == "Banana")
	fmt.Println("!=", 3 != 2)
	fmt.Println("<", 3 < 2)
	fmt.Println(">", 3 > 2)
	fmt.Println("<=", 3 <= 2)
	fmt.Println(">=", 3 >= 2)

	d1 := time.Unix(0, 0)
	d2 := time.Unix(0, 0)

	fmt.Println("Dates:", d1 == d2)
	fmt.Println("Dates:", d1.Equal(d2))

	p1 := Person{"João", "Silva"}
	p2 := Person{"João", "Silva"}
	fmt.Println("Person comparisson:", p1 == p2)
}
