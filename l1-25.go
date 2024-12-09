// Реализовать собственную функцию sleep.

package main

import (
	"fmt"
	"time"
)


func Sleep(seconds int) {
	start := time.Now()

	// Выполняем пустой цикл, пока не пройдет указанное количество секунд
	for time.Since(start) < time.Duration(seconds)*time.Second {
	}
}

func main() {
	fmt.Println("Начало ожидания...")
	Sleep(4) // Ожидание 4 секунды
	fmt.Println("Ожидание завершено.")
}