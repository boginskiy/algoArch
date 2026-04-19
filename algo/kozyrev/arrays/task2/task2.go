package main

// Удаление нулей - фильтрация элементов с сохранением порядка

import "fmt"

func FilteredNums(arr []int) []int {
	// Индекс записи следующего непустого элемента
	j := 0

	for i := range arr {
		if arr[i] != 0 {
			arr[j] = arr[i]
			j++
		}
	}
	return arr[:j]
}

func main() {
	arr := []int{3, 10, 0, 5, 1, 100, 43, 0}

	res := FilteredNums(arr)
	fmt.Println(res)
}

// Базовое решение
func filteredNums1(arr []int) []int {
	res := arr[:0]

	for _, num := range arr {
		if num != 0 {
			res = append(res, num)
		}
	}
	return res
}
