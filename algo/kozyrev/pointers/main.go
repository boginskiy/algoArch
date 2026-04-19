package main

// Удаление дубликатов in-place - два указателя на одном массиве

import "fmt"

func main() {
	arr := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	res := Conteinerre(arr)
	fmt.Println(res)
}

func Conteinerre(arr []int) int {
	l, r := 0, len(arr)-1
	maxRes := 0

	for l < r {
		tmpRes := min(arr[l], arr[r]) * (r - l)

		if tmpRes > maxRes {
			maxRes = tmpRes
		}

		if arr[l] < arr[r] {
			l++
		} else {
			r--
		}
	}
	return maxRes
}

func MinDifferenceArr2(arr1, arr2 []int) int {
	Ssort(arr1)
	Ssort(arr2)

	minN := arr1[len(arr1)-1]

	l, r := 0, 0

	for l < len(arr1) && r < len(arr2) {
		delta := ffunc(arr1[l] - arr2[r])

		if delta < minN {
			minN = delta
		}

		if arr1[l] < arr2[r] {
			l++
		} else {
			r++
		}
	}

	return minN
}

func ffunc(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func Ssort(arr []int) {
	if len(arr) < 2 {
		return
	}

	l, r, m := 0, len(arr)-1, len(arr)/2

	arr[m], arr[r] = arr[r], arr[m]

	for i := range arr {
		if arr[i] < arr[r] {
			arr[l] = arr[i]
			l++
		}
	}

	arr[l], arr[r] = arr[r], arr[l]
	Ssort(arr[:l])
	Ssort(arr[l+1:])
}
