package main

import (
	"fmt"
	"sort"
)

// Анаграммы - определение через хеш-таблицу и сортировку

// Определить являются ли две строки анаграммами, используя
// два подхода — сортировку символов и подсчёт частот символов через хеш-таблицы

// Определение анаграмм с использованием хеш-таблицы
func DefinAnagramsHashTb(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	hashTb := make(map[rune]int, 26)

	for i, s := range a {
		hashTb[s]++
		hashTb[rune(b[i])]--
	}

	for _, v := range hashTb {
		if v != 0 {
			return false
		}
	}
	return true
}

func main() {
	a := "listen"
	b := "silent"
	res := DefinAnagramsHashTb(a, b)
	fmt.Println(res)
}

// Определение анаграмм с помощью сортировки
func DefinAnagramsSort(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	runes1 := []rune(a)
	runes2 := []rune(b)
	sort.Slice(runes1, func(i, j int) bool { return runes1[i] < runes1[j] })
	sort.Slice(runes2, func(i, j int) bool { return runes2[i] < runes2[j] })

	return string(runes1) == string(runes2)
}
