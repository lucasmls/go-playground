package main

import "fmt"

// Color ...
type Color string

const (
	// Red ...
	Red Color = "red"
	// Green ...
	Green Color = "green"
	// Blue ...
	Blue Color = "blue"
)

// Size ...
type Size string

const (
	// Small ...
	Small Size = "small"
	// Medium ...
	Medium Size = "medium"
	// Large ...
	Large Size = "large"
)

// Product ...
type Product struct {
	name  string
	color Color
	size  Size
}

// Specification ...
type Specification interface {
	Satisfies(p *Product) bool
}

// ColorSpecification ...
type ColorSpecification struct {
	color Color
}

// Satisfies ...
func (spec ColorSpecification) Satisfies(p *Product) bool {
	return spec.color == p.color
}

// Filter ...
type Filter struct{}

// Filter ...
func (f *Filter) Filter(products []Product, spec Specification) []*Product {
	result := make([]*Product, 0)

	for i, p := range products {
		if spec.Satisfies(&p) {
			result = append(result, &products[i])
		}
	}

	return result
}

func main() {
	products := []Product{
		{"Apple", Green, Small},
		{"Tree", Green, Large},
		{"Car", Blue, Large},
	}

	f := Filter{}

	greenSpec := ColorSpecification{color: Green}
	for _, p := range f.Filter(products, greenSpec) {
		fmt.Printf(" - %s is green\n", p.name)
	}
}
