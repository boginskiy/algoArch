package main

import "testing"

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name   string
		arr    []int
		target int
		result []int
	}{
		{"test 1", []int{1, 2, 3, 7}, 8, []int{0, 3}},
		{"test 2", []int{1, 2, 3, 4}, 7, []int{2, 3}},
		{"test 3", []int{1, 1, 1, 2}, 3, []int{2, 3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			result := TwoSum(tt.arr, tt.target)

			if result[0] != tt.result[0] && result[1] != tt.result[1] {
				t.Errorf("fact %v, need: %v\n\r", result, tt.result)
			}
		})
	}
}
