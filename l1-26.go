// Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc).
// Функция проверки должна быть регистронезависимой.
package main

import (
	"fmt"
	"strings"
)

// Функция проверки уникальности символов в строке
func hasUniqueCharacters(s string) bool {
	// Приводим строку к нижнему регистру для регистронезависимой проверки
	s = strings.ToLower(s)

	// Создаем мапу для отслеживания встреченных символов
	charMap := make(map[rune]bool)

	for _, char := range s {
		// Если символ уже встречался, возвращаем false
		if charMap[char] {
			return false
		}
		// В противном случае добавляем символ в мапу
		charMap[char] = true
	}

	// Если все символы уникальные, возвращаем true
	return true
}

func main() {
	examples := []string{"abcd", "abCdefAaf", "aabcd"}

	for _, example := range examples {
		result := hasUniqueCharacters(example)
		fmt.Printf("Уникальные символы в \"%s\": %t\n", example, result)
	}
}