package main

import "fmt"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	//
	var mask [26]uint
	for i, b := range s {
		mask[b-'a']++    // s
		mask[t[i]-'a']-- // t
	}

	// Check
	for _, v := range mask {
		if v != 0 {
			return false
		}
	}
	return true
}

func main() {
	s := "anagram"
	t := "nagaram"

	res := isAnagram(s, t)
	fmt.Println(res)
}
