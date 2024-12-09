// К каким негативным последствиям может привести данный фрагмент кода, и как это исправить?
// Приведите корректный пример реализации.

//Последствия:

//1) Проблема с управлением памятью:
// Функция createHugeString создает строку размером 1024 символа, но хранит её в памяти до тех пор, пока переменная v
// не будет удалена (что произойдет, когда функция someFunc завершит выполнение).
// В результате это может привести к значительному потреблению памяти, особенно если someFunc вызывается многократно.
//
//2) Потенциальные проблемы с обработкой Unicode:
//В Go строки представлены в формате UTF-8, и простое обрезание строки может привести к нарушению кодировок символов,
//если обрезка осуществляется в середине многобайтового символа.

// Исправления:
package main

import (
	"fmt"
	"strings"
)

var justString string

func createHugeString(size int) string {
	return strings.Repeat("A", size)
}

func someFunc() {
	v := createHugeString(1 << 10)
	justString = v[:100]
}

func main() {
	someFunc()
	fmt.Println(justString)
}
