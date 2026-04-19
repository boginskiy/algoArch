package main

import "fmt"

// Поиск всех анаграмм в строке - скользящее окно с хешированием

func main() {

	s := "cbaebabacd"
	p := "abc"

	res := SearchAllAnagrams(s, p)
	fmt.Println(res)

}

func SearchAllAnagrams(s, p string) []int {
	result := make([]int, 0, 5)

	hashComp := [26]int{}
	for _, ch := range p {
		hashComp[ch-'a']++
	}

	hashComp2 := [26]int{}
	for _, ch := range s[:len(p)] {
		hashComp2[ch-'a']++
	}

	if hashComp == hashComp2 {
		result = append(result, 0)
	}

	l := 0

	for r := 3; r < len(s); r++ {
		hashComp2[s[l]-'a']--
		hashComp2[s[r]-'a']++

		if hashComp == hashComp2 {
			result = append(result, l+1)
		}
		l++
	}
	return result
}
