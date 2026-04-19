package main

import "testing"

func TestProductElems(t *testing.T) {
	tests := []struct {
		name   string
		arrIN  []int
		arrOUT []int
	}{
		{"test 1", []int{1, 2, 3}, []int{6, 3, 2}},
		{"test 2", []int{0, 0}, []int{0, 0}},
		{"test 3", []int{1}, []int{1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ProductElems(tt.arrIN)

			if !checkArr(result, tt.arrOUT) {
				t.Errorf("fact %v, need: %v\n\r", result, tt.arrOUT)
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
