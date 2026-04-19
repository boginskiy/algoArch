package main

import "fmt"

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	buy := prices[0]
	profit := 0

	for i := 1; i < len(prices); i++ {
		if buy > prices[i] {
			buy = prices[i]
		} else {
			if profit < prices[i]-buy {
				profit = prices[i] - buy
			}
		}
	}
	return profit
}

func main() {
	arr := []int{7, 6, 4, 3, 1}
	res := maxProfit(arr)
	fmt.Println(res)
}

// ### Обычное решение
// func maxProfit(prices []int) int {
// 	if len(prices) < 2 {
// 		return 0
// 	}

// 	buy := prices[0]
// 	profit := 0

// 	for i := 1; i < len(prices); i++ {
// 		if buy > prices[i] {
// 			buy = prices[i]
// 		} else {
// 			if profit < prices[i]-buy {
// 				profit = prices[i] - buy
// 			}
// 		}
// 	}
// 	return profit
// }
