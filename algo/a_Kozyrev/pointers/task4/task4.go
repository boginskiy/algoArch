package main

import "fmt"

// Минимальная разница между массивами - поиск ближайших пар

func main() {

	arr1 := []int{1, 3, 15, 11, 2}
	arr2 := []int{23, 127, 235, 19, 8}

	res := MinDifferenceArr(arr1, arr2)

	fmt.Println(res)

}

// 	[1 2  3  11  15]
//  [8 19 23 127 235]

func MinDifferenceArr(arr1, arr2 []int) int {
	QSort(arr1)
	QSort(arr2)

	f, s := 0, 0
	min := 1000

	for f < len(arr1) && s < len(arr2) {

		diff := fFunc(arr1[f] - arr2[s])
		if diff < min {
			min = diff
		}

		if arr1[f] < arr2[s] {
			f++
		} else {
			s++
		}
	}
	return min
}

func fFunc(n int) int {
	if n < 0 {
		return -n
	}
	return n
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
