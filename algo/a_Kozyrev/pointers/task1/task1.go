package main

// K ближайших элементов - поиск в отсортированном массиве

import "fmt"

func main() {

	arr := []int{-1, 0, 2, 3, 4, 6}
	k := 3
	x := 0

	res := SearhNeighbors(arr, x, k)
	fmt.Println(res)

}

func SearhNeighbors(arr []int, num, k int) []int {
	idx := BbSearch(arr, num)
	result := []int{}

	l, r := idx, idx+1
	dL, dR := 0, 0

	for k > 0 {

		if l >= 0 && r < len(arr) {
			dL = fFunc(arr[l] - arr[idx])
			dR = fFunc(arr[r] - arr[idx])

			if dL < dR {
				result = append(result, arr[l])
				l--
			} else {
				result = append(result, arr[r])
				r++
			}
		} else if l >= 0 {
			result = append(result, arr[l])
			l--
		} else {
			result = append(result, arr[r])
			r++
		}

		k--
	}

	return result
}

func fFunc(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func BbSearch(arr []int, num int) int {
	l, r := 0, len(arr)

	for l <= r {
		m := (l + r) / 2

		if arr[m] < num {
			l = m + 1
		} else if arr[m] > num {
			r = m - 1
		} else {
			return m
		}
	}
	return -1
}
