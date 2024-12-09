// Реализовать конкурентную запись данных в map.

package main

import (
	"fmt"
	"sync"
)

func main() {
	var safeMap sync.Map
	var wg sync.WaitGroup

	// Запись в map из нескольких горутин
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			safeMap.Store(i, i*i)
			fmt.Printf("Записано: %d -> %d\n", i, i*i)
		}(i)
	}

	wg.Wait()

	// Чтение из map
	safeMap.Range(func(key, value interface{}) bool {
		fmt.Printf("Получено: %d -> %d\n", key, value)
		return true
	})
}
