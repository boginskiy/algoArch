package main

import (
	"fmt"
)

func main() {

	s := "cbaebabacd"
	p := "abc"

	res := SearchAllAnagrams(s, p)
	fmt.Println(res)

}

func SearchAllAnagrams(s, p string) []int {
	result := []int{}

	hashP := [26]int{}
	for _, ch := range p {
		hashP[ch-'a']++
	}

	hashA := [26]int{}
	for _, ch := range s[:len(p)] {
		hashA[ch-'a']++
	}

	if hashP == hashA {
		result = append(result, 0)
	}

	l := 0
	for _, ch := range s[len(p):] {
		hashA[ch-'a']++
		hashA[s[l]-'a']--

		l++

		if hashP == hashA {
			result = append(result, l)
		}
	}

	return result
}

// func DefinLongestLine(s string) string {
// 	store := map[rune]int{}
// 	result := ""

// 	l := 0

// 	for r, ch := range s {
// 		if preIDX, ok := store[ch]; ok && l <= preIDX {
// 			l = preIDX + 1
// 		}

// 		store[ch] = r

// 		if len(result) < len(s[l:r+1]) {
// 			result = s[l : r+1]
// 		}
// 	}
// 	return result
// }

// var Store = map[rune]rune{
// 	'(': ')',
// 	'[': ']',
// 	'{': '}',
// }

// func ValidBrackets(l string) bool {
// 	stack := make([]rune, 0, len(l)/2)

// 	for _, ch := range l {
// 		_, ok := Store[ch]

// 		if ok {
// 			stack = append(stack, ch)
// 			continue
// 		}

// 		if len(stack) == 0 {
// 			return false
// 		}

// 		lastB := stack[len(stack)-1]
// 		if Store[lastB] != ch {
// 			return false
// 		}

// 		stack = stack[:len(stack)-1]

// 	}
// 	return len(stack) == 0
// }

// // func ReverseWords777(l string) string {
// 	arrWords := strings.Split(l, " ")

// 	for i, word := range arrWords {
// 		arrWords[i] = rev(word)
// 	}

// 	return strings.Join(arrWords, " ")
// }

// func rev(w string) string {
// 	arrChs := strings.Split(w, "")
// 	l, r := 0, len(arrChs)-1

// 	for l < r {
// 		arrChs[l], arrChs[r] = arrChs[r], arrChs[l]
// 		l++
// 		r--
// 	}
// 	return strings.Join(arrChs, "")
// }

// func DefinAnagramsHashTb777(a, b string) bool {
// 	if len(a) != len(b) {
// 		return false
// 	}

// 	slA := []rune(a)
// 	slB := []rune(b)

// 	sort.Slice(slA, func(i, j int) bool { return slA[i] < slA[j] })
// 	sort.Slice(slB, func(i, j int) bool { return slB[i] < slB[j] })

// 	return string(slA) == string(slB)

// }

// func DefinLongestLine2(line string) string {
// 	store := make(map[rune]int)
// 	result := ""
// 	l := 0

// 	for r, val := range line {
// 		if preI, ok := store[val]; ok && l <= preI {
// 			l = preI + 1
// 		}

// 		store[val] = r

// 		if len(result) < len(line[l:r+1]) {
// 			result = line[l : r+1]
// 		}
// 	}
// 	return result
// }

// func fFunc(input string) string {
// 	store := make(map[rune]int)
// 	result := ""
// 	l := 0

// 	for i, v := range input {

// 		if preNum, ok := store[v]; ok && l <= preNum {
// 			l = preNum + 1
// 		}

// 		store[v] = i

// 		if len(result) < len(input[l:i+1]) {
// 			result = input[l : i+1]
// 		}

// 	}
// 	return result
// }

// func fFunc(a string) string {
// 	arr := strings.Split(a, " ")

// 	for i, word := range arr {
// 		arr[i] = revo(word)
// 	}
// 	return strings.Join(arr, " ")
// }

// func revo(a string) string {
// 	arrR := []rune(a)
// 	l := 0
// 	r := len(arrR) - 1

// 	for l < r {
// 		arrR[l], arrR[r] = arrR[r], arrR[l]
// 		l++
// 		r--
// 	}
// 	return string(arrR)
// }

// func QuickSort(line string) {
// 	if len(line) < 2 {
// 		return
// 	}

// 	l := 0
// 	r := len(line) - 1
// 	p := len(line) / 2

// 	line[p], line[r] = line[r], line[p]

// 	for i := range arr {
// 		if arr[i] < arr[r] {
// 			arr[l], arr[i] = arr[i], arr[l]
// 			l++
// 		}
// 	}
// 	arr[l], arr[r] = arr[r], arr[l]

// 	QuickSort(arr[:l])
// 	QuickSort(arr[l+1:])
// }

// Через #-таблицу
// func fFunc(a, b string) bool {
// 	store := make(map[rune]int)

// 	if len(a) != len(b) {
// 		return false
// 	}

// 	for i, ch := range a {
// 		// Добавляем из a
// 		store[ch]++
// 		// Убираем
// 		store[rune(b[i])]--
// 	}

// 	for _, v := range store {
// 		if v != 0 {
// 			return false
// 		}
// 	}
// 	return true
// }
