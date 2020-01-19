package main

import "fmt"

// Car ...
type Car struct {
	brand        string
	maximumSpeed float64
}

// Ferrari ...
type Ferrari struct {
	Car
	name string
}

func main() {
	// car := Car{
	// 	brand:        "Ferrari",
	// 	maximumSpeed: 500.0,
	// }

	// ferrari := Ferrari{
	// 	Car:  car,
	// 	name: "F40",
	// }

	ferrari := Ferrari{}
	ferrari.brand = "Ferrari"
	ferrari.name = "F50"
	ferrari.maximumSpeed = 500.00

	fmt.Println(ferrari)
}
