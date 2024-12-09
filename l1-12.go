// Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.

package main

import (
	"fmt"
)

func main() {
	// Исходные строки
	animals := []string{"cat", "cat", "dog", "cat", "tree"}

	// Создаем множество с помощью мапы
	uniqueAnimals := make(map[string]struct{})

	// Заполняем множество уникальными значениями
	for _, animal := range animals {
		uniqueAnimals[animal] = struct{}{} // Используем пустую структуру как значение
	}

	// Выводим уникальные значения
	fmt.Println("Уникальные значения:", uniqueKeys(uniqueAnimals))
}

// Функция для получения уникальных ключей в виде среза
func uniqueKeys(set map[string]struct{}) []string {
	keys := make([]string, 0, len(set))
	for key := range set {
		keys = append(keys, key)
	}
	return keys
}