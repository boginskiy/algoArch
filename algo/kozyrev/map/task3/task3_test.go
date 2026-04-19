package main

import "testing"

func TestSetIntersection(t *testing.T) {
	tests := []struct {
		name   string
		set1   map[int]struct{}
		set2   map[int]struct{}
		result int
	}{
		{"test 1", map[int]struct{}{1: struct{}{}, 2: struct{}{}}, map[int]struct{}{2: struct{}{}}, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SetIntersection(tt.set1, tt.set2)

			_, ok := result[tt.result]

			if len(result) != 1 && !ok {
				t.Errorf("fact %v, need: %v\n\r", result, tt.result)
			}
		})
	}
}
