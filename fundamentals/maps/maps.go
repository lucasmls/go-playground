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
}
