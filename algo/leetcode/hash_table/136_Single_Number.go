package main

import "fmt"

func singleNumber(nums []int) int {
	res := 0
	for _, v := range nums {

		// fmt.Printf("%08b\n", v)
		res ^= v
		// fmt.Printf("%08b\n", res)
	}
	return res
}

func main() {

	arr := []int{4, 1, 2, 1, 2}
	res := singleNumber(arr)

	fmt.Println(res)
}

// Extra veiw
// v 00000100 / res ^= v 00000100
// v 00000001 / res ^= v 00000101
// v 00000010 / res ^= v 00000111
// v 00000001 / res ^= v 00000110
// v 00000010 / res ^= v 00000100
