package main

import "testing"

func TestSearhNeighbors(t *testing.T) {
	tests := []struct {
		name   string
		arr    []int
		num    int
		k      int
		result []int
	}{
		{"test 1", []int{-1, 0, 2, 3, 4, 6}, 0, 3, []int{0, -1, 2}},
		{"test 2", []int{0}, 0, 1, []int{0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SearhNeighbors(tt.arr, tt.num, tt.k)

			if !CompareArrs(result, tt.result) {
				t.Errorf("fact: %v, need: %v\n\r", result, tt.result)
			}

		})
	}
}

func CompareArrs(arr1, arr2 []int) bool {
	if len(arr1) != len(arr2) {
		return false
	}

	for i := range arr1 {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}
