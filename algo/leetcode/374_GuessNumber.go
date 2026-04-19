package main

func guessNumber(n int) int {
	l, r := 1, n

	for l <= r {
		mid := l + (r-l)/2
		num := guess(mid)

		if num == -1 {
			r = mid - 1
		} else if num == 1 {
			l = mid + 1
		} else {
			return mid
		}
	}
	return 0
}

// Составить массив
