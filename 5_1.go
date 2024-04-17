package main

import (
	"fmt"
	"math"
	"sync"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Создаем wait group для ожидания завершения горутин.
	var wg sync.WaitGroup
	wg.Add(2)

	// Инициализируем переменные для максимального и минимального элементов.
	var max, min int
	max = math.MinInt32
	min = math.MaxInt32

	// Горутина для вычисления максимального элемента.
	go func() {
		defer wg.Done()
		for _, num := range arr {
			if num > max {
				max = num
			}
		}
	}()

	// Горутина для вычисления минимального элемента.
	go func() {
		defer wg.Done()
		for _, num := range arr {
			if num < min {
				min = num
			}
		}
	}()

	// Ожидаем завершения обеих горутин.
	wg.Wait()

	fmt.Printf("Максимальный элемент: %d\n", max)
	fmt.Printf("Минимальный элемент: %d\n", min)
}
