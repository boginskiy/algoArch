package main

import "fmt"

func search(nums []int, target int) bool {
	l, r := 0, len(nums)-1

	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return true
		}

		//
		if nums[l] == nums[mid] && nums[mid] == nums[r] {
			l++
			r++

		} else if nums[l] <= nums[mid] {
			if nums[l] <= target && target < nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}

		} else {
			if nums[mid] < target && target <= nums[r] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
}

func main() {
	nums := []int{2, 3, 1, 1}
	target := 0

	res := search(nums, target)
	fmt.Println(res)
}

// ### Мое решение
// func search(nums []int, target int) bool {
// 	// Базовый случай
// 	if len(nums) == 0 {
// 		return false
// 	}

// 	l, r := 0, len(nums)-1

// 	for l <= r {
// 		mid := (l + r) / 2

// 		if nums[mid] == target {
// 			return true
// 		}

// 		// Запуск на левой-правой половине массива
// 		if nums[l] == nums[mid] {
// 			left := search(nums[:mid], target)
// 			right := search(nums[mid+1:], target)
// 			return left || right
// 		} else if nums[mid] > target {
// 			r = mid - 1
// 		} else {
// 			l = mid + 1
// 		}
// 	}
// 	return false
// }
