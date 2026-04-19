package main

import "fmt"

// Валидация скобок - проверка корректности последовательности

func ValidBrackets(str string) bool {
	stack := make([]rune, 0)
	brackets := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, ch := range str {
		switch ch {
		case '(', '{', '[':
			stack = append(stack, ch)
		case ')', '}', ']':
			if len(stack) == 0 || stack[len(stack)-1] != brackets[ch] {
				return false
			}
			stack = stack[:len(stack)-1]
		default:
			continue
		}
	}

	return len(stack) == 0
}

func main() {
	line := "({[(([[[()]]])])]})"

	res := ValidBrackets(line)
	fmt.Println(res)
}

// Стандартное решение
func ValidationBrackets2(line string) bool {
	skack := make([]rune, 0, len(line)/2)

	for _, s := range line {
		// Кладем на стек элементы
		if s == '(' || s == '{' || s == '[' {
			skack = append(skack, s)

		} else {
			if len(skack) == 0 {
				return false
			}
			// Снимаем со стека элемент
			lastS := skack[len(skack)-1]
			if s-lastS > 2 {
				return false
			}
			skack = skack[:len(skack)-1]
		}
	}
	return len(skack) == 0
}
