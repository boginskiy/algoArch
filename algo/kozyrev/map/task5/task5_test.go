package main

import "testing"

func TestSearchUniqueSymbol(t *testing.T) {
	tests := []struct {
		name   string
		line   string
		result int
	}{
		{"test 1", "leetcode", 0},
		{"test 2", "aaabbbcccdddzzz", -1},
		{"test 3", "abacabadaba", 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SearchUniqueSymbol(tt.line)

			if result != tt.result {
				t.Errorf("fact %v, need: %v\n\r", result, tt.result)
			}
		})
	}
}
