// Реализовать пересечение двух неупорядоченных множеств.

package main

import (
	"fmt"
)

func intersection(set1, set2 []int) []int {
	// Создаем мапу для хранения уникальных элементов первого множества (set1).
	setMap := make(map[int]struct{})
	for _, v := range set1 {
		setMap[v] = struct{}{} // Добавляем каждый элемент первого множества в мапу.
	}

	// Создаем срез для хранения результата - пересечения множеств.
	var result []int
	for _, v := range set2 {
		// Проверяем, содержится ли элемент второго множества (set2) в мапу.
		if _, found := setMap[v]; found {
			result = append(result, v)
			delete(setMap, v)          // Удаляем элемент из карты, чтобы избежать дублирования.
		}
	}
	return result // Возвращаем срез с пересечением.
}

func main() {
	// Определяем два множества для пересечения.
	set1 := []int{1, 2, 3, 4, 5}
	set2 := []int{3, 4, 5, 6, 7}

	// Вызываем функцию для нахождения пересечения.
	result := intersection(set1, set2)
	fmt.Println("Intersection:", result)
}