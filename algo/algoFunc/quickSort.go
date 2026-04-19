package main

func quickSort(arr []int) {
	// Базовый случай
	if len(arr) < 2 {
		return
	}

	l, r, m := 0, len(arr)-1, len(arr)/2

	// Делаем перестановку r <-> m
	arr[r], arr[m] = arr[m], arr[r]

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
