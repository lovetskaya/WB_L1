// Удалить i-ый элемент из слайса.

package main

import (
	"fmt"
)

func removeElement(arr []int, index int) []int {
	if index < 0 || index >= len(arr) {
		fmt.Printf("Индекс %d выходит за пределы среза.\n", index)
		return arr // возвращаем оригинальный срез, если индекс некорректен
	}
	return append(arr[:index], arr[index+1:]...) // Объединяем левую и правую части среза
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	indexToRemove := 2 // Индекс элемента для удаления (в этом случае элемент 3)

	fmt.Printf("Исходный срез: %v\n", arr)
	updatedArr := removeElement(arr, indexToRemove)
	fmt.Printf("Срез после удаления элемента с индексом %d: %v\n", indexToRemove, updatedArr)
}