package main

import "fmt"

func bsearch(arr []int, t int) bool {
	if len(arr) == 0 {
		return false
	}

	l, r := 0, len(arr)-1
	m := l + r/2

	if arr[m] > t {
		return bsearch(arr[:m], t)
	} else if arr[m] < t {
		return bsearch(arr[m+1:], t)
	} else {
		return true
	}
}

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	firstL, firstR := 0, len(matrix)-1

	// Сначала бинарно ищем подмассив
	for firstL <= firstR {
		firstMID := (firstL + firstR) / 2
		deltaArr := matrix[firstMID]

		if target < deltaArr[0] {
			firstR = firstMID - 1
		} else if target > deltaArr[len(deltaArr)-1] {
			firstL = firstMID + 1
		} else {
			return bsearch(deltaArr, target)
		}
	}
	return false
}

func main() {
	matrix := [][]int{{1}, {3}, {5}}
	target := 4

	res := searchMatrix(matrix, target)
	fmt.Println(res)
}
