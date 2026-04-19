package main

import "fmt"

func search3(nums []int, target int) int {
	// 1. Найти место сдвига и создать новый, но уже отсортированный массив
	shift := 0
	for i, v := range nums {
		if i > 0 && v < nums[i-1] {
			nums = append(nums[i:], nums[:i]...)
			shift = len(nums) - i
			break
		}
	}

	// Бинарный поиск
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2

		if nums[mid] > target {
			r = mid - 1
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			reverce := (mid - shift)

			if reverce < 0 {
				return len(nums) + reverce
			}
			return reverce
		}
	}
	return -1
}

func main() {
	arr := []int{5, 1, 3}
	target := 3

	res := search3(arr, target)
	fmt.Println(res)
}

// Прикольное решение !
// func search(nums []int, target int) int {
// 	left, right := 0, len(nums)-1

// 	for left <= right {
// 		mid := (left + right) / 2

// 		if nums[mid] == target {
// 			return mid
// 		}

// 		if nums[left] <= nums[mid] {
// 			if nums[left] <= target && target < nums[mid] {
// 				right = mid - 1
// 			} else {
// 				left = mid + 1
// 			}
// 		} else {
// 			if nums[mid] < target && target <= nums[right] {
// 				left = mid + 1
// 			} else {
// 				right = mid - 1
// 			}
// 		}
// 	}
// 	return -1
// }

// Найти место и создать новый отсортированный массив
// func search(nums []int, target int) int {
// 	// 1. Найти место сдвига и создать новый, но уже отсортированный массив
// 	shift := 0
// 	for i, v := range nums {
// 		if i > 0 && v < nums[i-1] {
// 			nums = append(nums[i:], nums[:i]...)
// 			shift = len(nums) - i
// 			break
// 		}
// 	}

// 	// Бинарный поиск
// 	l, r := 0, len(nums)-1
// 	for l <= r {
// 		mid := (l + r) / 2

// 		if nums[mid] > target {
// 			r = mid - 1
// 		} else if nums[mid] < target {
// 			l = mid + 1
// 		} else {
// 			reverce := (mid - shift)

// 			if reverce < 0 {
// 				return len(nums) + reverce
// 			}
// 			return reverce
// 		}
// 	}
// 	return -1
// }
