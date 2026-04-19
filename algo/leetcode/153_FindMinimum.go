package main

import (
	"fmt"
)

func findMin(nums []int) int {
	l, r := 0, len(nums)-1

	for l < r {
		mid := (l + r) / 2

		if nums[mid] > nums[r] {
			l = mid + 1
		} else {
			r = mid
		}
	}

	return nums[l]
}

func main() {
	arr := []int{13, 14, 2, 6, 10}
	res := findMin(arr)
	fmt.Println(res)

}

// ### Полурабочий вариант
// func findMin(nums []int) int {
// 	if len(nums) == 0 {
// 		return 0
// 	}

// 	l, r := 0, len(nums)-1

// 	for l <= r {
// 		mid := (l + r) / 2

// 		//Мы нашли этот подмассив
// 		if nums[l] <= nums[mid] && nums[l] > nums[r] {
// 			l = mid + 1

// 		} else if nums[l] < nums[mid] && nums[mid] < nums[r] {
// 			r = mid - 1

// 		} else {
// 			return nums[mid]
// 		}

// 	}
// 	return 0
// }

// ### Самый простой вариант ;))))))
// func findMin(nums []int) int {
// 	if len(nums) == 0 {
// 		return 0
// 	}
// 	sort.Ints(nums)
// 	return nums[0]

// }

// ### Решение с O(N)
// func findMin(nums []int) int {
// 	if len(nums) == 0 {
// 		return 0
// 	}
// 	for i, num := range nums {
// 		if i > 0 {
// 			if num < nums[i-1] {
// 				nums = append(nums[i:], nums[:i]...)
// 				break
// 			}
// 		}
// 	}
// 	return nums[0]
// }
