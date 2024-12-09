package main

import (
	"fmt"
	"math"
)

// Структура Point с инкапсулированными параметрами x и y
type Point struct {
	x, y float64
}

// Конструктор для создания новой точки
func NewPoint(x, y float64) Point {
	return Point{x: x, y: y}
}

// Метод для вычисления расстояния до другой точки
func (p Point) Distance(other Point) float64 {
	return math.Sqrt(math.Pow(other.x-p.x, 2) + math.Pow(other.y-p.y, 2))
}

func main() {
	// Создаем две точки
	point1 := NewPoint(1.0, 2.0)
	point2 := NewPoint(4.0, 6.0)

	// Вычисляем расстояние между ними
	distance := point1.Distance(point2)

	// Выводим результат
	fmt.Printf("Расстояние между точками (%.2f, %.2f) и (%.2f, %.2f) равно %.2f\n",
		point1.x, point1.y, point2.x, point2.y, distance)
}