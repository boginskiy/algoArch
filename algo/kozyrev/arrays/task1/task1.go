package main

// #### Слияние отсортированных массивов - merge двух массивов in-palce

import "fmt"

// MergeTwoSortedArrays. Подход #1: использование свободного места в одном из массивов
// Функция объединяет два отсортированных массива arr1 и arr2 in-place,
// предполагая, что arr2 расширяется и вмещает весь arr1.

// ВНИМАНИЕ! Необходимо наличие избыточного пространства в одном из массивов.

func MergeTwoSortedArrays(arr1, arr2 []int) {
	M1 := len(arr1) - len(arr2) - 1
	M2 := len(arr2) - 1
	MR := len(arr1) - 1

	for M1 >= 0 && M2 >= 0 {
		if arr1[M1] > arr2[M2] {
			arr1[MR] = arr1[M1]
			M1--
		} else {
			arr1[MR] = arr2[M2]
			M2--
		}
		MR--
	}

	// Если остались элементы в arr2
	if M2 >= 0 {
		copy(arr1[:MR+1], arr2[:M2+1])
	}
}

func main() {

	arr1 := make([]int, 8) // Результирующий массив
	copy(arr1, []int{0, 3, 5, 7, 25})

	arr2 := []int{7, 10, 15}

	MergeTwoSortedArrays(arr1, arr2)

	fmt.Println(arr1)

}

// Слияние массивов с использованием дополнительной памяти
func mergeTwoSortedArrays2(arr1, arr2 []int) []int {
	result := make([]int, len(arr1)+len(arr2))
	l, r, k := 0, 0, 0

	for l < len(arr1) && r < len(arr2) {
		if arr1[l] < arr2[r] {
			result[k] = arr1[l]
			l++
		} else {
			result[k] = arr2[r]
			r++
		}
		k++
	}

	for l < len(arr1) {
		result[k] = arr1[l]
		l++
		k++
	}

	for r < len(arr2) {
		result[k] = arr2[r]
		r++
		k++
	}

	return result
}
