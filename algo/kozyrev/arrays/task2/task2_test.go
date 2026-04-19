package main

import "testing"

func TestFilteredNums(t *testing.T) {
	tests := []struct {
		name   string
		arr    []int
		arrRes []int
	}{
		{"test 1", []int{1, 0, 3, 0, 5}, []int{1, 3, 5}},
		{"test 2", []int{1, 2, 3}, []int{1, 2, 3}},
		{"test 3", []int{0, 0, 1}, []int{1}},
		{"test 4", []int{0, 0, 0}, []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FilteredNums(tt.arr)

			if !checkArr(result, tt.arrRes) {
				t.Errorf("fact %v, need: %v\n\r", result, tt.arrRes)
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
