// Разработать программу, которая в рантайме способна определить тип переменной:
// int, string, bool, channel из переменной типа interface{}.

package main

import (
	"fmt"
)

func determineType(v interface{}) {
	switch v := v.(type) {
	case int:
		fmt.Printf("Переменная типа int со значением: %d\n", v)
	case string:
		fmt.Printf("Переменная типа string со значением: %s\n", v)
	case bool:
		fmt.Printf("Переменная типа bool со значением: %t\n", v)
	case chan struct{}:
		fmt.Printf("Переменная типа channel\n")
	default:
		fmt.Println("Тип переменной не поддерживается.")
	}
}

func main() {
	var intVar interface{} = 42
	var stringVar interface{} = "Привет, мир!"
	var boolVar interface{} = true
	var chanVar interface{} = make(chan struct{})

	determineType(intVar)
	determineType(stringVar)
	determineType(boolVar)
	determineType(chanVar)
	determineType(3.14) // Пример с неподдерживаемым типом
}