package main

import "testing"

func TestMergeTwoSortedArrays(t *testing.T) {
	tests := []struct {
		name   string
		arr1   []int
		arr2   []int
		arrRes []int
	}{
		{"test 1", []int{1, 2, 0, 0}, []int{3, 5}, []int{1, 2, 3, 5}},
		{"test 2", []int{10, 0, 0, 0}, []int{3, 5, 7}, []int{3, 5, 7, 10}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MergeTwoSortedArrays(tt.arr1, tt.arr2)

			if !checkArr(tt.arr1, tt.arrRes) {
				t.Errorf("fact %v, need: %v\n\r", tt.arr1, tt.arrRes)
			}
		})
	}
}

func checkArr(arr1, arr2 []int) bool {
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
