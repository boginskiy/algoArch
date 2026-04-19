package main

import "fmt"

// Минимальная разница между k элементами - выбор подмножества
// Массив данных должен быть отсортирован

func main() {
	arr := []int{1, 4, 7, 8, 10}
	k := 3

	res := MinDifference(arr, k)
	fmt.Println(res)
}

func MinDifference(arr []int, k int) int {
	// Сортировка данных
	QSort(arr) //

	min, max := 0, k-1
	minDiff := arr[len(arr)-1]

	for max < len(arr) {
		if arr[max]-arr[min] < minDiff {
			minDiff = arr[max] - arr[min]
		}
		min++
		max++
	}

	return minDiff
}

func QSort(arr []int) {
	if len(arr) < 2 {
		return
	}
	l, r, m := 0, len(arr)-1, len(arr)/2

	arr[m], arr[r] = arr[r], arr[m]

	for i, v := range arr {
		if v < arr[r] {
			arr[l], arr[i] = arr[i], arr[l]
			l++
		}
	}

	arr[l], arr[r] = arr[r], arr[l]
	QSort(arr[:l])
	QSort(arr[l+1:])
}
