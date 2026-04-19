package main

import (
	"fmt"
)

func search2(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for l <= r {
		mid := (l + r) / 2

		if nums[mid] > target {
			r = mid - 1
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

func main() {
	arr := []int{-1, 0, 3, 5, 9, 12}
	tagret := 2

	res := search2(arr, tagret)
	fmt.Println(res)
}

// Мое решение
// func search(nums []int, target int) int {
// 	l, r := 0, len(nums)

// 	for l < r {
// 		mid := (l + r) / 2

// 		if nums[mid] == target {
// 			return mid

// 		} else if nums[mid] > target {
// 			r = mid

// 		} else {
// 			if l == mid {
// 				l++
// 			} else {
// 				l = mid
// 			}
// 		}
// 	}
// 	return -1
// }
