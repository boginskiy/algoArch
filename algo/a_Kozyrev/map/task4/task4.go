package main

import "fmt"

// Определение чемпионов - агрегация и поиск максимумов

type Result struct {
	Name     string  // Имя участника
	Category string  // Категория соревнования
	Score    float64 // Оценочный балл
}

func main() {
	results := []Result{
		{Name: "Игорь", Category: "Категория A", Score: 9.5},
		{Name: "Ольга", Category: "Категория B", Score: 8.0},
		{Name: "Андрей", Category: "Категория A", Score: 9.8},
		{Name: "Елена", Category: "Категория C", Score: 7.5},
		{Name: "Дмитрий", Category: "Категория B", Score: 8.2},
		{Name: "Светлана", Category: "Категория C", Score: 7.7},
	}

	res := DefinResult(results)
	fmt.Println(res)

}

func DefinResult(results []Result) map[string]Result {
	total := map[string]Result{}

	for _, result := range results {

		if preResult, ok := total[result.Category]; !ok {
			total[result.Category] = result
		} else {
			if preResult.Score < result.Score {
				total[result.Category] = result
			}
		}

	}
	return total
}
