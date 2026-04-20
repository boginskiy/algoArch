package main

import "testing"

func TestMinDifferenceArr(t *testing.T) {
	tests := []struct {
		name   string
		arr1   []int
		arr2   []int
		result int
	}{
		{"test 1", []int{1, 3, 15, 11, 2}, []int{23, 127, 235, 19, 8}, 3},
		{"test 2", []int{1, 1, 1}, []int{1, 1, 1}, 0},
		{"test 3", []int{1, 3, 5, 7, 9}, []int{4, 6, 8, 10, 12}, 1},
		{"test 4", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, []int{11, 12}, 1},
		{"test 5", []int{-100, -50, 0, 50, 100}, []int{-99, -49, 1, 51, 101}, 1},
		{"test 6", []int{1}, []int{10}, 9},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			result := MinDifferenceArr(tt.arr1, tt.arr2)

			if result != tt.result {
				t.Errorf("fact %v, need: %v\n\r", result, tt.result)
			}
		})
	}
}
