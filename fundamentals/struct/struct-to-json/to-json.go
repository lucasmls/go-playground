package main

import (
	"encoding/json"
	"fmt"
)

// Product ...
type Product struct {
	ID    int      `json:"id"`
	Name  string   `json:"name"`
	Price float64  `json:"price"`
	Tags  []string `json:"tags"`
}

func main() {
	// Struct to JSON
	notebook := Product{
		ID:    1,
		Name:  "Macbook Pro",
		Price: 8999,
		Tags:  []string{"i5", "Apple", "16gb"},
	}

	fmt.Println(notebook)

	jsonNotebook, _ := json.Marshal(notebook)
	fmt.Println(string(jsonNotebook))

	// JSON to Struct
	var tablet Product
	tabletJSON := `{"id": 2, "name": "iPad", "price": 4.500, "tags": ["Apple", "Thin", "Apple Pencil"]}`
	json.Unmarshal([]byte(tabletJSON), &tablet)
	fmt.Println(tablet)
}
