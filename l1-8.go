// Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.

package main

import (
	"fmt"
)

// Устанавливает i-й бит в 1
func setBit(n *int64, i int) {
	*n |= (1 << i) // Устанавливаем i-й бит в 1
}

// Устанавливает i-й бит в 0
func clearBit(n *int64, i int) {
	*n &= ^(1 << i) // Устанавливаем i-й бит в 0
}

// Получает значение i-го бита
func getBit(n int64, i int) int64 {
	return (n >> i) & 1 // Возвращает значение i-го бита
}

func main() {
	var num int64 = 0 // Изначально 0
	var i int

	// Установка i-го бита в 1
	i = 3
	setBit(&num, i)
	fmt.Printf("После установки %d-го бита в 1: %b (значение: %d)\n", i, num, num)

	// Установка i-го бита в 0
	i = 3
	clearBit(&num, i)
	fmt.Printf("После установки %d-го бита в 0: %b (значение: %d)\n", i, num, num)

	// Получение значения i-го бита
	i = 3
	bitValue := getBit(num, i)
	fmt.Printf("Значение %d-го бита: %d\n", i, bitValue)
}
