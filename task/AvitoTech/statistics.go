package main

import "fmt"

// Условие:
// Нужно найти чемпионов соревнований по шагам и они должны соответствовать данным критериям:
// 1. Прошли наибольшее общее количество шагов за все дни соревнования.
// 2. Не пропустили ни одного дня (то есть у пользователя есть запись в каждом дне).

// Входные данные:
// - statistics: список дней соревнования
// - Каждый день — это список словарей вида: {userId: int, steps: int}

// Выходные данные:
// - champions: объект вида {userIds: []int, steps: int}

// Пояснение:
// - userIds — список победителей (если несколько с одинаковым максимумом)
// - steps — общее количество шагов победителя(ей)

// Пример 1:
// statistics = [
//  [ { userId: 1, steps: 1000 }, { userId: 2, steps: 1500 } ],
//  [ { userId: 2, steps: 1000 } ]
// ]

// Вывод:
// champions = { userIds: [2], steps: 2500 }

// Пример 2:
// statistics = [
//  [ { userId: 1, steps: 2000 }, { userId: 2, steps: 1500 } ],
//  [ { userId: 2, steps: 4000 }, { userId: 1, steps: 3500 } ]
// ]

// Вывод:
// champions = { userIds: [1, 2], steps: 5500 }

type Statistics struct {
	UserId int
	Steps  int
}

type Result struct {
	UserIds []int
	Steps   int
}

func getChampions(statistics [][]Statistics) Result {
	// Если соревнований не было.
	if len(statistics) == 0 {
		return Result{}
	}

	// Сохраняем предварительные данные по участникам
	resultSteps := make(map[int]int, len(statistics[0]))
	resultDays := make(map[int]int, len(statistics[0]))

	for _, statisticOfDay := range statistics {
		for _, player := range statisticOfDay {
			// Считаем количество дней
			resultDays[player.UserId]++
			// Считаем количество шагов
			resultSteps[player.UserId] += player.Steps
		}
	}

	max := 0
	maxArr := make([]int, 0, len(statistics))

	for userID, days := range resultDays {
		if days != len(statistics) || max > resultSteps[userID] {
			continue
		}

		if max < resultSteps[userID] {
			max = resultSteps[userID]
			maxArr = maxArr[:0]
		}

		maxArr = append(maxArr, userID)
	}

	return Result{
		UserIds: maxArr,
		Steps:   max,
	}
}

func main() {

	// test1 := [][]Statistics{
	// 	{{UserId: 1, Steps: 1000}, {UserId: 2, Steps: 1500}},
	// 	{{UserId: 2, Steps: 1000}}}

	test1 := [][]Statistics{
		{{UserId: 1, Steps: 2000}, {UserId: 2, Steps: 1500}},
		{{UserId: 2, Steps: 4000}, {UserId: 1, Steps: 3500}}}

	result := getChampions(test1)
	fmt.Println(result)
}

// Решение
func getChampions2(statistics [][]Statistics) Result {
	if len(statistics) == 0 {
		return Result{}
	}

	// Подсчёт общего кол-ва дней, в которых участвовал каждый пользователь
	daysParticipated := make(map[int]int)

	// Подсчёт общего кол-ва шагов каждого пользователя
	totalSteps := make(map[int]int)

	// Общее кол-во дней, за которые считались шаги
	totalDays := len(statistics)

	// Алгоритмическая сложность O(n*m)
	for _, day := range statistics { // Выделяем один из дней, пример: [ { userId: 1, steps: 1000 }, { userId: 2, steps: 1500 } ]
		for _, stat := range day {
			daysParticipated[stat.UserId]++
			totalSteps[stat.UserId] += stat.Steps
		}
	}

	// Валидируем пользователей и находим тех, которые не пропустили ни одного дня
	// Алгоритмическая сложность O(m)
	validUsers := make(map[int]bool)

	for userId, days := range daysParticipated {
		if days == totalDays {
			validUsers[userId] = true
		}
	}

	// Находим максимальное кол-во шагов среди провалидированных пользователей
	maxSteps := 0

	// Алгоритмическая сложность O(v)
	for userId, _ := range validUsers {
		if totalSteps[userId] > maxSteps {
			maxSteps = totalSteps[userId]
		}
	}

	// Собираем всех пользователей с maxSteps
	var resultIds []int

	// Алгоритмическая сложность O(v)
	for userId, _ := range validUsers {
		if totalSteps[userId] == maxSteps {
			resultIds = append(resultIds, userId)
		}
	}

	return Result{
		UserIds: resultIds,
		Steps:   maxSteps,
	}
}
