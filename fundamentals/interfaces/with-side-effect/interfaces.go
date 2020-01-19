package main

import "fmt"

// Sporting ...
type Sporting interface {
	toggleTurbo()
}

// Ferrari ...
type Ferrari struct {
	model       string
	turbo       bool
	actualSpeed int
}

func (f *Ferrari) toggleTurbo() {
	f.turbo = !f.turbo
}

func main() {
	f40 := Ferrari{
		model:       "F40",
		turbo:       false,
		actualSpeed: 0,
	}

	fmt.Println(f40)
	f40.toggleTurbo()
	fmt.Println(f40)

	var laFerrari Sporting = &Ferrari{
		model:       "La Ferrari",
		turbo:       false,
		actualSpeed: 0,
	}

	fmt.Println(laFerrari)
	laFerrari.toggleTurbo()
	fmt.Println(laFerrari)
}
