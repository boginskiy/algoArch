package main

import "fmt"

// Вращение массива - сдвиг элементов на k позиций

func RotatingArray(arr []int, k int) {
	n := len(arr)
	k %= n

	reverse(arr, 0, n-1) // Шаг 1: разворачиваем весь массив
	reverse(arr, 0, k-1) // Шаг 2: разворачиваем первые k элементов
	reverse(arr, k, n-1) // Шаг 3: разворачиваем оставшуюся часть
}

func reverse(arr []int, left, right int) {
	for left < right {
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}
}

func main() {

	arr := []int{1, 2, 3, 4, 5, 6}
	k := 3

	RotatingArray(arr, k)
	fmt.Println(arr)

	// [6 5 4 3 2 1]
	// [4 5 6 3 2 1]
	// [4 5 6 1 2 3]
}

// Базовый способ с выделением памяти
func RotatingArray2(arr []int, k int) []int {
	res := make([]int, len(arr))

	for i := range arr {
		res[(i+k)%len(arr)] = arr[i]
	}
	return res
}
