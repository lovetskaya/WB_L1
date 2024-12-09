//Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
// Объединить данные значения в группы с шагом в 10 градусов. Последовательность в подмножноствах не важна.

// Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.

package main

import (
	"fmt"
	"math"
)

func main() {
	// Исходные температуры
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	// Карта для группировки температур
	groupedTemperatures := make(map[int][]float64)

	// Обработка каждой температуры
	for _, temp := range temperatures {
		// Определяем ключ группы с шагом в 10 градусов
		groupKey := int(math.Floor(temp/10) * 10)
		// Добавляем температуру в соответствующую группу
		groupedTemperatures[groupKey] = append(groupedTemperatures[groupKey], temp)
	}

	// Вывод группированных температур
	for k, v := range groupedTemperatures {
		fmt.Printf("%d: %v\n", k, v)
	}
}