package goarea

import (
	"math"
)

//Pi é uma proporção númerica definida pela relação entre o perímetro de uma circunferência e seu diâmetro

const Pi = 3.1316

// é responsável por calcular a área da circuferência.
func Circle(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

// Rect é responsável por calcular a área de um retangulo
func Rect(base, altura float64) float64 {
	return base * altura
}

// não é uma função visível pq colocquei o _ na frente.
// func _TrianguloEq(base, altura, float64) float64 {
// 	return (base * altura) / 2
