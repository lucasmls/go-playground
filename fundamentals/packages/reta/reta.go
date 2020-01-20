package main

import "math"

/*
 * Every func, struct and etc that starts with a capitalized letter it's public,
 * and when it's not capitalized, its private.

 * Public and private in package level, not in file level
 */

// Ponto representa uma coordenada no plano cartesiano
type Ponto struct {
	x float64
	y float64
}

func catetos(a Ponto, b Ponto) (catX float64, catY float64) {
	catX = math.Abs(a.x - b.x)
	catY = math.Abs(a.y - b.y)
	return
}

// Distancia é responsável por calcular a distância entre os dois pontos
func Distancia(a, b Ponto) float64 {
	catX, catY := catetos(a, b)
	return math.Sqrt(math.Pow(catX, 2) + math.Pow(catY, 2))
}
