//Разработать программу, которая переворачивает слова в строке.
//Пример: «snow dog sun — sun dog snow».

package main

import (
	"fmt"
	"strings"
)

func reverseWords(input string) string {
	// Разбиваем строку на слова
	words := strings.Fields(input)

	// Переворачиваем слова
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i] // Меняем местами слова
	}

	// Объединяем слова обратно в строку
	return strings.Join(words, " ")
}

func main() {
	input := "snow dog sun"
	fmt.Printf("Исходная строка: \"%s\"\n", input)
	reversed := reverseWords(input)
	fmt.Printf("Перевернутая строка: \"%s\"\n", reversed)
}