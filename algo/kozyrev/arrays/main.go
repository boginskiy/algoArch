package main

import (
	"fmt"
)

// Поиск пропущенных чисел - найти все числа из диапазона [1, n]

func main() {
	// arr := []int{1, 0, 0, 0, 1, 0, 1}
	// arr := []int{1, 0, 0, 0}
	// arr := []int{0, 1, 0, 1, 0}
	// arr := []int{1, 0, 0, 1, 0, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0}
	arr := []int{1, 0, 0, 1}
	// arr := []int{0, 1, 0, 1, 0, 1, 0, 0, 1, 1, 0, 1, 1, 1}

	res := maxDistToClosest(arr)
	fmt.Println(res)

}

func maxDistToClosest(arr []int) int {
	result := 0

	// left
	left := 0
	for left < len(arr) && arr[left] == 0 {
		left++
	}
	if result < len(arr[:left]) {
		result = len(arr[:left])
	}

	// right
	right := len(arr) - 1
	for right >= 0 && arr[right] == 0 {
		right--
	}
	if result < len(arr[right+1:]) {
		result = len(arr[right+1:])
	}

	m := 0
	c := 0
	for _, v := range arr[left : right+1] {
		if v == 0 {
			c++
		} else {
			if m < c {
				m = c
			}
			c = 0
		}
	}

	//
	if m%2 == 0 && result < (m/2) {
		return (m / 2)
	}

	if m%2 != 0 && result < (m/2)+1 {
		return (m / 2) + 1
	}

	return result
}
