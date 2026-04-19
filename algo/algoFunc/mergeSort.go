package main

import "fmt"

func MergeSort(arr []int) []int {
	//Базовый случай. Когда массив len < 2
	if len(arr) < 2 {
		return arr
	}

	// Разделяем массив на левую и правую части
	left := MergeSort(arr[:len(arr)/2])
	right := MergeSort(arr[len(arr)/2:])

	// Делаем сборку нового массива
	l, r, k := 0, 0, 0

	result := make([]int, len(arr))

	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			result[k] = left[l]
			l++
		} else {
			result[k] = right[r]
			r++
		}
		k++
	}

	// Если остались данные в подассиве left, перекидываем и в result
	for l < len(left) {
		result[k] = left[l]
		l++
		k++
	}

	// Если остались данные в подмассиве right, перекидываем и в result
	for r < len(right) {
		result[k] = right[r]
		r++
		k++
	}
	return result
}

func main() {
	arr := []int{56, 19, 100, 3, 6, 7, 0}
	res := MergeSort(arr)

	fmt.Println(res)
}
