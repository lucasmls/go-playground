package main

import "fmt"

// Printable ...
type Printable interface {
	toString() string
}

// Person ...
type Person struct {
	name    string
	surname string
}

func (person Person) toString() string {
	return person.name + " " + person.surname
}

// Product ...
type Product struct {
	name  string
	price float64
}

func (product Product) toString() string {
	return fmt.Sprintf("%s - R$ %.2f", product.name, product.price)
}

func print(item Printable) {
	fmt.Println(item.toString())
}

func main() {
	var person Printable = Person{
		name:    "Lucas",
		surname: "Mendes",
	}

	fmt.Println(person.toString())
	print(person)

	p2 := Product{"Calça Jeans", 179.90}
	print(p2)
}
