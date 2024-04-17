//Реализовать функцию, которая из заданного каталога выбирает имена файлов изображений.
package main

import "fmt"

// sumArray вычисляет сумму элементов числового массива.
func sumArray(arr []int) int {
    sum := 0
    for _, num := range arr {
        sum += num
    }
    return sum
}

func main() {
    numbers := []int{1, 2, 3, 4, 5}
    result := sumArray(numbers)
    fmt.Printf("Сумма элементов массива: %d\n", result)
}
