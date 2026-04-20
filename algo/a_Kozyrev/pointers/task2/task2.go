package main

import "fmt"

// Контейнер с максимальной водой - оптимизация площади

func main() {
	arr := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	res := Conteiner(arr)
	fmt.Println(res)
}

func Conteiner(arr []int) int {
	l, r := 0, len(arr)-1
	max := 0

	for l < r {
		delta := min(arr[l], arr[r]) * (r - l)

		if max < delta {
			max = delta
		}

		// Сдвигаем тот элемент что меньше.
		if arr[r] < arr[l] {
			r--
		} else {
			l++
		}
	}
	return max
}
