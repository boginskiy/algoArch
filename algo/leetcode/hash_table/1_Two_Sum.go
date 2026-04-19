package main

import "fmt"

func twoSum(nums []int, target int) []int {
	set := make(map[int]int)

	for i, v := range nums {

		tmpV := target - v

		if idx, ok := set[tmpV]; ok {
			return []int{idx, i}
		}
		set[v] = i
	}
	return []int{}
}

func main() {
	arr := []int{3, 2, 4}
	target := 6

	res := twoSum(arr, target)
	fmt.Println(res)
}
