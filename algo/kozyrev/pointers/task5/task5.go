package main

import "fmt"

// Удаление дубликатов in-place - два указателя на одном массиве

func main() {
	arr := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}

	res := DeleteDubl(arr)
	fmt.Println(res)
}

func DeleteDubl(arr []int) int {
	QSort(arr)

	l := 0

	for r := 1; r < len(arr); r++ {
		if arr[l] != arr[r] {
			l++
			arr[l] = arr[r]
		}
	}
	return len(arr[:l+1])
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
