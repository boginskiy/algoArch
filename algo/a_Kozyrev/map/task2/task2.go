package main

import "fmt"

// Группировка элементов - построение map по критерию

func main() {

	users := []struct {
		Name string
		Age  int
	}{
		{"Иван", 25},
		{"Марина", 30},
		{"Алексей", 25},
		{"Анна", 30},
		{"Сергей", 40},
	}

	res := GroupElems(users)
	fmt.Println(res)
}

func GroupElems(list []struct {
	Name string
	Age  int
}) map[int][]string {

	result := make(map[int][]string)

	for _, v := range list {
		result[v.Age] = append(result[v.Age], v.Name)
	}

	return result
}
