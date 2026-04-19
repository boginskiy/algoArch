package main

import "fmt"

func isValid(s string) bool {
	res := make([]rune, 0, len(s))

	for _, c := range s {

		if c == '(' || c == '[' || c == '{' {
			res = append(res, c)

		} else {
			if len(res) == 0 {
				return false
			}

			sub := rune(c) - res[len(res)-1]

			if sub > 2 || sub < 1 {
				return false
			} else {
				res = res[:len(res)-1]
			}
		}
	}
	return len(res) == 0
}

func main() {
	s := "(]"

	res := isValid(s)
	fmt.Println(res)
}

// ### Стандарт
// func isValid(s string) bool {
// 	res := make([]rune, 0, len(s))

// 	for _, c := range s {

// 		if c == '(' || c == '[' || c == '{' {
// 			res = append(res, c)

// 		} else {
// 			if len(res) == 0 {
// 				return false
// 			}

// 			sub := rune(c) - res[len(res)-1]

// 			if sub > 2 || sub < 1 {
// 				return false
// 			} else {
// 				res = res[:len(res)-1]
// 			}
// 		}
// 	}
// 	return len(res) == 0
// }
