package main

import "fmt"

// Product ...
type Product struct {
	name     string
	price    float64
	discount float64
}

func (p Product) priceWithDiscount() float64 {
	return p.price * (1 - p.discount)
}

func main() {
	var macbook Product
	macbook = Product{
		name:     "Macbook Pro",
		price:    8800,
		discount: 0.05,
	}

	fmt.Println(macbook, macbook.priceWithDiscount())

	iphone := Product{"Iphone 11", 5500, 0.10}
	fmt.Println(iphone, iphone.priceWithDiscount())
}
