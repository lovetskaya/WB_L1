// Разработать программу, которая будет последовательно отправлять значения в канал, а с другой стороны канала — читать.
// По истечению N секунд программа должна завершаться.
package main

import (
	"fmt"
	"time"
)

func main() {
	const duration = 10 // Программа будет работать 10 секунд
	dataChannel := make(chan int) // Создаем канал для передачи данных

	// Запускаем горутину для отправки данных в канал
	go func() {
		for i := 1; ; i++ {
			dataChannel <- i // Отправляем значение в канал
			time.Sleep(1 * time.Second) // Пауза
		}
	}()

	timer := time.NewTimer(time.Duration(duration) * time.Second)// Запускаем таймер на завершение программы

	// Читаем данные из канала
	go func() {
		for {
			select {
			case data := <-dataChannel:
				fmt.Printf("Получено: %d\n", data)
			case <-timer.C:
				fmt.Println("Время вышло, завершаем программу.")
				close(dataChannel) // Закрываем канал перед завершением
				return
			}
		}
	}()

	<-timer.C // Ожидаем, пока таймер сработает
}
