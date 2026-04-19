package main

import (
	"fmt"
)

func findAnagrams(s string, p string) []int {
	count := [26]int{}
	for i := range p {
		count[p[i]-'a']++
	}

	res := []int{}

	for start, end := 0, 0; end < len(s); end++ {
		count[s[end]-'a']--

		if count == [26]int{} {
			res = append(res, start)
		}
		if end+1 >= len(p) {
			count[s[start]-'a']++
			start++
		}
	}
	return res
}

func main() {
	s := "cbaebabacd"
	p := "abc"

	res := findAnagrams(s, p)
	fmt.Println(res)
}

// // Можно поломать голову
// func findAnagrams(s string, p string) []int {
// 	count := [26]int{}
// 	for i := range p {
// 		count[p[i]-'a']++
// 	}
// 	res := []int{}

// 	for start, end := 0, 0; end < len(s); end++ {
// 		count[s[end]-'a']--
// 		if count == [26]int{} {
// 			res = append(res, start)
// 		}
// 		if end+1 >= len(p) {
// 			count[s[start]-'a']++
// 			start++
// 		}
// 	}
// 	return res
// }

// // Неплохое решение
// func findAnagrams(s string, p string) []int {
// 	if len(s) < len(p) {
// 		return nil
// 	}

// 	var res []int
// 	var freqS, freqP [26]uint8

// 	for i := 0; i < len(p); i++ {
// 		freqS[s[i]-'a']++
// 		freqP[p[i]-'a']++
// 	}

// 	for i := 0; i < len(s)-len(p)+1; i++ {
// 		if freqS == freqP {
// 			res = append(res, i)
// 		}
// 		if i+len(p) < len(s) {
// 			freqS[s[i]-'a']--
// 			freqS[s[i+len(p)-'a']]++
// 		}
// 	}
// 	return res
// }

// // Легкая оптимизация
// func findAnagrams(s string, p string) []int {
// 	if len(s) < len(p) || len(p) == 0 {
// 		return []int{}
// 	}

// 	result := make([]int, 0, len(s)/len(p)+1)

// 	maskP := [26]uint8{} // Маска для шаблона
// 	maskS := [26]uint8{} // Маска для скользящего окна

// 	for i := 0; i < len(p); i++ {
// 		maskP[p[i]-'a']++
// 		maskS[s[i]-'a']++
// 	}

// 	if maskP == maskS {
// 		result = append(result, 0)
// 	}

// 	for i := 1; i <= len(s)-len(p); i++ {
// 		prevChar := s[i-1]
// 		nextChar := s[i+len(p)-1]

// 		maskS[prevChar-'a']--
// 		maskS[nextChar-'a']++

// 		if maskP == maskS {
// 			result = append(result, i)
// 		}
// 	}
// 	return result
// }

// Мое решение
// func CalcMask(p string) [26]uint8 {
// 	m := [26]uint8{}
// 	for _, v := range p {
// 		m[v-'a']++
// 	}
// 	return m
// }

// func findAnagrams(s string, p string) []int {
// 	result := make([]int, 0, 10)

// 	if len(s) < len(p) {
// 		return result
// 	}

// 	maskP := CalcMask(p)
// 	maskS := CalcMask(s[0:len(p)])

// 	if maskP == maskS {
// 		result = append(result, 0)
// 	}

// 	for i, j := 1, len(p); j < len(s); {
// 		maskS[s[i-1]-'a']--
// 		maskS[s[j]-'a']++

// 		//
// 		if maskP == maskS {
// 			result = append(result, i)
// 		}
// 		//
// 		i++
// 		j++
// 	}
// 	return result
// }
