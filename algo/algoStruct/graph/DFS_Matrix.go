package main

import "fmt"

func dfs2(graph [][]int, start int) []string {
	visited := make([]bool, len(graph))
	result := make([]string, 0)
	stack := []int{start}

	for len(stack) > 0 {
		currV := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if visited[currV] {
			continue
		}

		visited[currV] = true

		// Преобразование индекса в букву
		result = append(result, fmt.Sprintf("%c", 'A'+currV))

		// Соседи текущей вершины
		for i, neighbor := range graph[currV] {
			if neighbor == 1 && !visited[i] {
				stack = append(stack, i)
			}
		}

	}
	return result
}

func main() {

	graph := [][]int{
		{0, 1, 1, 0}, // A
		{1, 0, 0, 1}, // B
		{1, 0, 0, 1}, // C
		{0, 1, 1, 0}, // D
	}

	startV := 0
	res := dfs2(graph, startV)
	fmt.Println(res)
}
