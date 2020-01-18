package main

import "fmt"

// Item ...
type Item struct {
	id    int
	qtd   int
	price float64
}

// Order ...
type Order struct {
	id    int
	items []Item
}

func (order Order) calculateTotal() float64 {
	total := 0.0
	for _, item := range order.items {
		total += item.price * float64(item.qtd)
	}

	return total
}

func main() {
	order := Order{
		id: 1,
		items: []Item{
			Item{
				id:    2,
				price: 10,
				qtd:   3,
			},
		},
	}

	fmt.Println(order, order.calculateTotal())
}
