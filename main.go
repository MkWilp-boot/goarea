package goarea

import "math"

// Pi é a constante do círculo
const Pi = 3.1416

// Circ calcula a area da circunferencia
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

// Rect calcula a area de um retangulo
func Rect(base, altura float64) float64 {
	return base * altura
}
