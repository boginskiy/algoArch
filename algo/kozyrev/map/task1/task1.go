package main

import "fmt"

// Two Sum - поиск двух чисел с заданной суммой

func main() {
	arr := []int{2, 7, 11, 15}
	k := 9

	res := TwoSum(arr, k)
	fmt.Println(res)
}

func TwoSum(arr []int, k int) []int {
	store := map[int]int{}

	for j, v := range arr {
		i, ok := store[k-v]

		if !ok {
			store[v] = j

		} else {
			return []int{i, j}
		}
	}
	return []int{}
}
