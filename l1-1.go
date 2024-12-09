package main

import (
	"fmt"
)
// Решила реализовать через указатель чтобы легче управлять памятью и избегать копирования структур

// Определяем структуру Human
type Human struct {
	height int
	weight int
	hair   string
}

// Метод для структуры Human
func (h *Human) Introduce() {
	fmt.Printf("Мой рост: %d, мой вес: %d, цвет волос: %s.\n", h.height, h.weight, h.hair)
}

// Определяем структуру Action, содержащую указатель на Human
type Action struct {
	act *Human
}

// Метод для структуры Action
func (a Action) PerformAction() {
	fmt.Println("Я выполняю действие.")
}

func main() {
	// Создаем экземпляр Human
	human := &Human{
		height: 175,
		weight: 65,
		hair:   "черный",
	}

	// Создаем экземпляр Action, передавая указатель на Human
	action := Action{
		act: human,
	}

	action.act.Introduce() // Мой рост: 175, мой вес: 65, цвет волос: черный.
	action.PerformAction() // Я выполняю действие.
}