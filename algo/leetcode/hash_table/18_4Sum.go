package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	result := make([][]int, 0, 10)
	sort.Ints(nums)
	Long := len(nums)

	for i := 0; i < Long-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		minI := nums[i] + nums[i+1] + nums[i+2] + nums[i+3]
		if minI > target {
			break
		}

		maxI := nums[i] + nums[Long-1] + nums[Long-2] + nums[Long-3]
		if maxI < target {
			continue
		}

		for j := i + 1; j < Long-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			minJ := nums[i] + nums[j] + nums[j+1] + nums[j+2]
			if minJ > target {
				break
			}

			maxJ := nums[i] + nums[j] + nums[Long-1] + nums[Long-2]
			if maxJ < target {
				continue
			}

			l, r := j+1, Long-1

			for l < r {
				R := nums[i] + nums[j] + nums[l] + nums[r]

				if target == R {
					result = append(result, []int{nums[i], nums[j], nums[l], nums[r]})
					l++
					for l < r && nums[l] == nums[l-1] {
						l++
					}
				} else if target < R {
					r--
				} else {
					l++
				}
			}
		}
	}
	return result
}

func main() {

	arr := []int{-2, -1, 0, 0, 1, 2} // 1, 0, -1, 0, -2, 2
	target := 0

	// sort.Ints(arr)
	// fmt.Println(arr)

	res := fourSum(arr, target)
	fmt.Println(res)

}

// // Лучшее решение на leetcode
// func fourSum(nums []int, target int) [][]int {
// 	result := make([][]int, 0, 16)
// 	sort.Ints(nums)
// 	N := len(nums)

// 	for i := 0; i < N-3; i++ {
// 		if i > 0 && nums[i] == nums[i-1] {
// 			continue
// 		}

// 		// Part 1
// 		MinI := nums[i] + nums[i+1] + nums[i+2] + nums[i+3]
// 		if MinI > target {
// 			// Если собранный минимум (потому что список сортирован!)
// 			// уже больше target, то все остальные значения будут
// 			// лишь увеличивать 'MinI', значит далее смотреть нет смысла
// 			break
// 		}
// 		MaxI := nums[i] + nums[N-1] + nums[N-2] + nums[N-3]
// 		if MaxI < target {
// 			continue
// 		}

// 		// Part 2
// 		for j := i + 1; j < N-2; j++ {
// 			if j > i+1 && nums[j] == nums[j-1] {
// 				continue
// 			}

// 			MinJ := nums[i] + nums[j] + nums[j+1] + nums[j+2]
// 			if MinJ > target {
// 				break
// 			}
// 			MaxJ := nums[i] + nums[j] + nums[N-1] + nums[N-2]
// 			if MaxJ < target {
// 				continue
// 			}

// 			// Part 3
// 			l, r := j+1, N-1
// 			for l < r {
// 				sum := nums[i] + nums[j] + nums[l] + nums[r]
// 				if sum == target {
// 					result = append(result, []int{nums[i], nums[j], nums[l], nums[r]})

// 					// Убираем потенциальные дубли
// 					l++
// 					for l < r && nums[l] == nums[l-1] {
// 						l++
// 					}
// 				} else if sum < target {
// 					l++
// 				} else {
// 					r--
// 				}
// 			}
// 		}
// 	}
// 	return result
// }

// // Мой средний вариант
// func fourSum(nums []int, target int) [][]int {
// 	sort.Ints(nums)
// 	res := [][]int{}
// 	tmpA := 111

// 	for a := 0; a < len(nums)-3; a++ {
// 		// Jump
// 		if nums[a] == tmpA {
// 			continue
// 		}
// 		tmpB := 222

// 		for b := a + 1; b < len(nums)-2; b++ {
// 			// Jump
// 			if nums[b] == tmpB {
// 				continue
// 			}

// 			// Указатели
// 			c := b + 1
// 			d := len(nums) - 1

// 			// target2
// 			target2 := target - (nums[a] + nums[b])

// 			// Tmp
// 			tmpC := 333
// 			tmpD := 444

// 			for c < d {
// 				sum := nums[c] + nums[d]

// 				// Указатели сдвигаем друг к другу
// 				if target2 == sum {
// 					//
// 					if !(tmpC == nums[c] && tmpD == nums[d]) {
// 						res = append(res, []int{nums[a], nums[b], nums[c], nums[d]})
// 						tmpC = nums[c]
// 						tmpD = nums[d]
// 					}
// 					c++
// 					d--
// 				} else if sum < target2 {
// 					c++
// 				} else {
// 					d--
// 				}
// 			}
// 			//
// 			tmpB = nums[b]
// 		}
// 		//
// 		tmpA = nums[a]
// 	}
// 	return res
// }

// // Мой худший вариант
// func fourSum(nums []int, target int) [][]int {
// 	sort.Ints(nums)
// 	res := [][]int{}
// 	Numbers := map[string]struct{}{}

// 	for i := 0; i < len(nums); i++ {
// 		target2 := target - nums[i] // nums[i] - v1

// 		for j := i + 1; j < len(nums); j++ {
// 			target3 := target2 - nums[j] // nums[j] - v2

// 			// Хранилище
// 			set := make(map[int]struct{})

// 			for k := j + 1; k < len(nums); k++ {
// 				v3 := target3 - nums[k] //  nums[k] - v4

// 				if _, ok := set[v3]; ok {

// 					strNum := fmt.Sprintf("%d%d%d%d", nums[i], nums[j], v3, nums[k])

// 					if _, ok := Numbers[strNum]; !ok {
// 						res = append(res, []int{nums[i], nums[j], v3, nums[k]})
// 						Numbers[strNum] = struct{}{}
// 					}

// 				} else {
// 					set[nums[k]] = struct{}{}
// 				}
// 			}
// 		}
// 	}
// 	return res
// }

// // Подсказка 1. Пробуем решить задачу на основе 3Sum
// func ThreeSum(nums []int, target int) [][]int {
// 	set := make(map[int]struct{})
// 	res := [][]int{}

// 	// v1
// 	for i, v1 := range nums {

// 		for j := i + 1; j < len(nums); j++ {

// 			target2 := target - v1 // target2

// 			v2 := target2 - nums[j] // nums[j] - v3

// 			if _, ok := set[v2]; ok {
// 				res = append(res, []int{v1, v2, nums[j]})
// 			} else {
// 				set[nums[j]] = struct{}{}
// 			}
// 		}
// 	}
// 	return res
// }
