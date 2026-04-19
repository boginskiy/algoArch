package main

import "fmt"

// Потенциальная задача в Vk
// 849. Maximize Distance to Closest Person

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

func maxDistToClosest(seats []int) int {
	result := 0
	l := 0

	for i := range seats {
		if seats[i] == 1 {
			if l == 0 {
				result = i
			}

			result = max(result, (i-l)/2)
			l = i
		}
	}

	return max(result, len(seats)-1-l)
}

// Черновое решение.
// func maxDistToClosest(seats []int) int {
// 	mWay := 0

// 	// Сначала пробуем получить delta с лева
// 	l := 0
// 	for l < len(seats) && seats[l] == 0 {
// 		l++
// 	}
// 	if mWay < len(seats[:l]) {
// 		mWay = len(seats[:l])
// 	}

// 	// Сначала пробуем получить delta с права
// 	r := len(seats) - 1
// 	for r >= 0 && seats[r] == 0 {
// 		r--
// 	}
// 	if mWay < len(seats[r+1:]) {
// 		mWay = len(seats[r+1:])
// 	}

// 	// Случаи те, что по середине
// 	m := 0
// 	cnt := 0

// 	for i := l; i < len(seats); i++ {
// 		if seats[i] == 0 {
// 			cnt++
// 		} else {
// 			if m < cnt {
// 				m = cnt
// 			}
// 			cnt = 0
// 		}
// 	}

// 	// Нечетное m
// 	if mWay < (m/2+1) && m%2 != 0 {
// 		mWay = (m/2 + 1)
// 	}

// 	// Четное m
// 	if mWay < (m / 2) {
// 		mWay = (m / 2)
// 	}

// 	return mWay
// }
