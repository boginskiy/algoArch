package main

import (
	"testing"
)

func TestMinDifference(t *testing.T) {
	tests := []struct {
		name   string
		arr    []int
		k      int
		result int
	}{
		{"test 1", []int{1, 4, 7, 8, 10}, 3, 3},
		{"test 2", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5, 4},
		{"test 3", []int{-5, -3, -1, 0, 2, 4, 6}, 4, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := MinDifference(tt.arr, tt.k)
			if res != tt.result {
				t.Errorf("fact: %v, need: %v\n\r", res, tt.result)
			}
		})
	}
}
