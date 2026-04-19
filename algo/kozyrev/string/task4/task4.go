package main

import (
	"fmt"
)

func DefinLongestLine(line string) string {
	store := make(map[rune]int)
	result := ""
	l := 0

	for r, val := range line {
		if prevIdx, ok := store[val]; ok && l <= prevIdx {
			l = prevIdx + 1
		}

		store[val] = r

		if len(result) < len(line[l:r+1]) {
			result = line[l : r+1]
		}
	}
	return result
}

func main() {
	st := "pwwkew"
	res := DefinLongestLine(st)
	fmt.Println(res)
}
