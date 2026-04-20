package main

// Произведение элементов кроме текущего - решение без деления

import "fmt"

// productElems. Основная идея заключается в расчёте префиксных и суффиксных произведений:
//
//	Префиксное произведение — это произведение всех предыдущих элементов относительно текущего.
//	Суффиксное произведение — произведение всех последующих элементов относительно текущего.
//  Тогда произведение всех элементов, кроме текущего, получается умножением
//  префиксного и суффиксного произведений.

// !!!
func ProductElems(arr []int) []int {
	result := make([]int, len(arr))

	// Префиксные произведения
	result[0] = 1
	for i := 1; i < len(arr); i++ {
		result[i] = result[i-1] * arr[i-1]
	}

	// Суффиксные произведения. Получаем результат
	suff := 1
	for i := len(arr) - 1; i >= 0; i-- {
		result[i] *= suff
		suff *= arr[i]
	}

	return result
}

func main() {
	arr := []int{1, 2, 3, 4}
	res := ProductElems(arr)
	fmt.Println(res)
}

// Использование префиксов и суффиксов
func productElems3(arr []int) []int {
	result := make([]int, len(arr))
	prodPref := make([]int, len(arr))
	prodSuff := make([]int, len(arr))

	// Этап 1: Расчёт префиксных произведений
	prodPref[0] = 1
	for i := 1; i < len(arr); i++ {
		prodPref[i] = prodPref[i-1] * arr[i-1]
	}

	// Этап 2: Расчёт суффиксных произведений
	prodSuff[len(arr)-1] = 1
	for i := len(arr) - 2; i >= 0; i-- {
		prodSuff[i] = prodSuff[i+1] * arr[i+1]
	}

	// Этап 3: Получение финального результата путём умножения префикса и суффикса
	for i := 0; i < len(arr); i++ {
		result[i] = prodPref[i] * prodSuff[i]
	}

	return result
}

// Базовое решение с делением.
func productElems2(arr []int) []int {
	mult := 1

	for _, num := range arr {
		mult *= num
	}

	for i, num := range arr {
		arr[i] = mult / num
	}
	return arr
}
