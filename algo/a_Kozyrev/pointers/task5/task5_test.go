package main

import "testing"

func TestDeleteDubl(t *testing.T) {
	tests := []struct {
		name   string
		arr    []int
		result int
	}{
		{"test 1", []int{1, 1, 2}, 2},
		{"test 2", []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 5},
		{"test 3", []int{1, 2, 3, 4, 5}, 5},
		{"test 4", []int{1, 1, 1, 1, 1, 2}, 2},
		{"test 4", []int{1, 2, 3, 4, 4, 4}, 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := DeleteDubl(tt.arr)

			if result != tt.result {
				t.Errorf("fact %v, need: %v\n\r", result, tt.result)
			}
		})
	}
}
