// Реализовать бинарный поиск встроенными методами языка.

package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{10, 7, 8, 9, 1, 5}
	sort.Ints(arr) // Сортируем массив перед бинарным поиском

	fmt.Println("Отсортированный массив:", arr)

	a := 8 // Значение, которое мы ищем

	// Выполняем бинарный поиск
	index := sort.Search(len(arr), func(i int) bool {
		return arr[i] >= a
	})

	// Проверяем, найдено ли значение
	if index < len(arr) && arr[index] == a {
		fmt.Printf("Элемент %d найден на индексе %d.\n", a, index)
	} else {
		fmt.Printf("Элемент %d не найден.\n", a)
	}
}
