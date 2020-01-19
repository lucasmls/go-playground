package main

import "fmt"

type grade float64

func (g grade) between(minGrade float64, maxGrade float64) bool {
	return float64(g) >= minGrade && float64(g) <= maxGrade
}

func getConceptByGrade(g grade) string {
	if g.between(9.0, 10.0) {
		return "A"
	}

	if g.between(7.0, 8.99) {
		return "B"
	}

	if g.between(5.0, 7.99) {
		return "C"
	}

	if g.between(3.0, 4.99) {
		return "D"
	}

	return "E"
}

func main() {
	fmt.Println(getConceptByGrade(8.55))
	fmt.Println(getConceptByGrade(9.55))
}
