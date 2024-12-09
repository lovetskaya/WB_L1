// Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива,
// во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.

package main

import (
	"fmt"
	"sync"
)

func main() {
	// Исходный массив чисел
	numbers := []int{1, 2, 3, 4, 5}

	// Создаем два канала
	numberChan := make(chan int)
	resultChan := make(chan int)

	var wg sync.WaitGroup

	// Запускаем горутину для записи чисел в первый канал
	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, x := range numbers {
			numberChan <- x // Передаем число в первый канал
		}
		close(numberChan) // Закрываем канал после записи всех чисел
	}()

	// Запускаем горутину для обработки чисел во втором канале
	wg.Add(1)
	go func() {
		defer wg.Done()
		for x := range numberChan {
			resultChan <- x * 2 // Умножаем число на 2 и отправляем во второй канал
		}
		close(resultChan) // Закрываем канал результата
	}()

	// Основная горутина отвечает за вывод результатов в stdout
	wg.Add(1)
	go func() {
		defer wg.Done()
		for result := range resultChan {
			fmt.Println(result) // Выводим результат
		}
	}()

	// Ожидаем завершения всех горутин
	wg.Wait()
}