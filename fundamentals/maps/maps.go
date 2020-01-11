package main

import "fmt"

func main() {

	approved := make(map[int]string)

	approved[12312312312] = "Lucas"
	approved[34534534554] = "Laisla"

	for cpf, name := range approved {
		fmt.Printf("%s, CPF: %d\n", name, cpf)
	}

	salary := map[string]float64{
		"Gabriel Souza": 2350.00,
		"Luis Souza":    6650.00,
	}

	fmt.Println(salary)
	delete(salary, "Luis Souza")
	fmt.Println(salary)
	fmt.Println("-----")

	usersByLeter := map[string]map[string]float64{
		"L": {
			"Lucas Mendes": 123.12,
			"Laisla":       2550.00,
		},
		"G": {
			"Gabriel Lima": 7500.50,
		},
		"A": {
			"Ana Silveira": 3200.00,
		},
	}

	for letter, users := range usersByLeter {
		fmt.Println(letter, "-", users)
	}

}
