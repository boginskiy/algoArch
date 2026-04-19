package main

import (
	"fmt"
	"strings"
)

// Реверс слов - переворот с сохранением порядка

// Программа принимает строку и возвращает новую строку, где каждое
// отдельное слово перевёрнуто (слова идут в том же порядке, но сами слова
// записаны задом наперёд)

// Стандартное решение
func ReverseWords(line string) string {
	words := strings.Split(line, " ")
	for i, word := range words {
		words[i] = reverse(word)
	}
	return strings.Join(words, " ")
}

func reverse(w string) string {
	runes := []rune(w)
	l, r := 0, len(runes)-1

	for l < r {
		runes[l], runes[r] = runes[r], runes[l]
		l++
		r--
	}
	return string(runes)
}

func main() {
	l := "Hello world!"
	res := ReverseWords(l)
	fmt.Println(res)
}
