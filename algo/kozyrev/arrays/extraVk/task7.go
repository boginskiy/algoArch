package main

// // Потенциальная задача в Vk
// 228. Summary Ranges

import (
	"fmt"
	"strconv"
)

func main() {
	arr := []int{0, 1, 2, 4, 5, 7}
	QSort(arr)
	res := summaryRanges(arr)
	fmt.Println(res)
}

func summaryRanges(nums []int) []string {
	var res []string
	n := len(nums)
	if n == 0 {
		return res
	}

	for i := 0; i < n; i++ {
		start := nums[i]
		for i+1 < n && nums[i+1] == nums[i]+1 {
			i++
		}
		if start == nums[i] {
			res = append(res, strconv.Itoa(start))
		} else {
			res = append(res, fmt.Sprintf("%d->%d", start, nums[i]))
		}
	}
	return res
}

func summaryRanges2(arr []int) []string {
	result := []string{}
	l := 0
	r := 1

	for r < len(arr) {
		if arr[r]-1 != arr[r-1] {
			result = append(result, fFunc(arr[l:r]))
			l = r
		}
		r++
	}

	if l < len(arr) {
		result = append(result, fFunc(arr[l:]))
	}

	return result

}

func fFunc(arr []int) string {
	switch len(arr) {
	case 0:
		return ""
	case 1:
		return fmt.Sprintf("%d", arr[0])
	default:
		return fmt.Sprintf("%d->%d", arr[0], arr[len(arr)-1])
	}
}

func QSort(arr []int) {
	if len(arr) < 2 {
		return
	}

	l, r, m := 0, len(arr)-1, len(arr)/2
	arr[m], arr[r] = arr[r], arr[m]

	for i := range arr {
		if arr[i] < arr[r] {
			arr[l], arr[i] = arr[i], arr[l]
			l++
		}
	}

	arr[l], arr[r] = arr[r], arr[l]
	QSort(arr[:l])
	QSort(arr[l+1:])
}
