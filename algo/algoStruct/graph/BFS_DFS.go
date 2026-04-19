package main

import (
	"fmt"
)

func breadthFirstSearch(graph map[string][]string, target string) string {
	queue := []string{target}
	checkIn := make(map[string]struct{})

	for len(queue) > 0 { // Пока очередь не пуста
		name := queue[0] // Извлекаем человека
		queue = queue[1:]

		// Проверка человека, если он не проверялся ранее
		if _, ok := checkIn[name]; !ok {

			// Проверка, что человек является продавцом манго
			if name == target {
				return name
			} else {
				queue = append(queue, graph[name]...)
				checkIn[name] = struct{}{}
			}
		}
	}
	return ""
}

func main() {
	graph := map[string][]string{
		"you":    {"alice", "bob", "claire"},
		"bob":    {"anuj", "peggy"},
		"alice":  {"peggy"},
		"claire": {"thom", "jonny"},
		"anuj":   {},
		"peggy":  {},
		"thom":   {},
		"jonny":  {},
	}

	man := breadthFirstSearch(graph, "thom")
	fmt.Printf("Mango Seller: %v\n", man)

}
