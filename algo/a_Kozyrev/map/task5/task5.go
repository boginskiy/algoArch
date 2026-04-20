package main

import "fmt"

// Первый уникальный символ - поиск через частотную таблицу

func main() {
	line := "aabbccdd"

	res := SearchUniqueSymbol(line)
	fmt.Println(res)

}

func SearchUniqueSymbol(line string) int {
	store := map[rune]int{}

	for _, ch := range line {
		store[ch]++
	}

	for i, ch := range line {
		if v := store[ch]; v == 1 {
			return i
		}
	}
	return -1
}
