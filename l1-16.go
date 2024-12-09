// Реализовать быструю сортировку массива (quicksort) встроенными методами языка.

// Можно использовать пакет sort для сортировки срезов. Однако, если реализовать сам алгоритм быстрой сортировки,
// то я написала встроенную функцию для работы с массивами:

package main

import (
	"fmt"
)

// Функция быстрой сортировки
func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr // массив из одного элемента уже отсортирован
	}

	// Определяем опорный элемент и создаем срезы для элементов
	pi := arr[len(arr)/2]

	var left, right []int
	for _, num := range arr {
		if num < pi {
			left = append(left, num)
		} else if num > pi {
			right = append(right, num)
		}
	}

	// Рекурсивно сортируем левую и правую части
	return append(append(quickSort(left), pi), quickSort(right)...)
}

func main() {
	arr := []int{99, 45, 13, 24, 65, 1, 24, 54}

	fmt.Printf("Исходный массив: %v\n", arr)
	sortedArr := quickSort(arr)
	fmt.Printf("Отсортированный массив: %v\n", sortedArr)
}