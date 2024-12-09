// Поменять местами два числа без создания временной переменной.

package main

import (
	"fmt"
)

func main() {
	a := 4
	b := 11

	// В Go реализована поддержка множественного присваивания
	a, b = b, a

	fmt.Println("a - ", a)
	fmt.Println("b - ", b)
}