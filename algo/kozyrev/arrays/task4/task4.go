package main

// Поиск пропущенных чисел - найти все числа из диапазона [1, n]

import "fmt"

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

// Способ, который меняет исходный массив.
func FindMissNums(arr []int) []int {
	// Перебираем каждый элемент и помечаем соответствующим числом позицию как посещённую

	for _, num := range arr {
		idx := abs(num) - 1
		if arr[idx] > 0 {
			arr[idx] = -arr[idx]
		}
	}

	// Собираем индексы положительных элементов, они соответствуют отсутствующим числам
	var result []int
	for i, v := range arr {
		if v > 0 {
			result = append(result, i+1)
		}
	}

	return result
}

func main() {
	arr := []int{3, 3, 2, 7, 1, 3, 4, 2}
	res := FindMissNums(arr)

	arr2 := []int{3, 3, 2, 7, 1, 3, 4, 2}
	res2 := FindMissNums3(arr2)
	fmt.Println(res, res2)
}

// Дополнительная память
func FindMissNums3(arr []int) []int {
	extraArr := make([]bool, len(arr)+1)
	result := make([]int, 0, 10)

	for _, v := range arr {
		extraArr[v] = true
	}

	for i := 1; i < len(extraArr); i++ {
		if extraArr[i] != true {
			result = append(result, i)
		}
	}
	return result
}

// Поиск пропущенных чисел - найти все числа из диапазона [1, n]
func FindMissNums2(arr []int) []int {
	quickSort(arr)

	prepar := make([]int, len(arr)+1)
	for i := range arr {
		prepar[arr[i]] += 1
	}

	result := make([]int, 0, 5)
	for i := 1; i < len(prepar); i++ {
		if prepar[i] == 0 {
			result = append(result, i)
		}
	}
	return result
}

// Быстрая сортировкаю
func quickSort(arr []int) {
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
	quickSort(arr[:l])
	quickSort(arr[l+1:])
}
