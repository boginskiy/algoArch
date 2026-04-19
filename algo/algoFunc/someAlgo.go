package main

import "fmt"

// Определение является ли число простым
func SearchPrimeNumbers(n int) bool {
	if n == 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// Решето Эратосфена
func Eratosthenes(n int) []int {
	arr := make([]int, n)
	for i := range arr {
		arr[i] = i
	}

	arr[0], arr[1] = -1, -1
	for i := 2; i < n; i++ {
		for j := i * i; j < n; {
			arr[j] = -1
			j += i
		}
	}
	return arr
}

func main() {
	res := Eratosthenes(20)
	fmt.Println(res) // [-1 -1 2 3 -1 5 -1 7 -1 -1 -1 11 -1 13 -1 -1 -1 17 -1 19]
}
