package main

func dfs(graph map[string][]string, startV string) {
	stack := make([]string, 0, len(graph))
	visited := make(map[string]bool)

	// Добавляем стартовую вершину в стек
	stack = append(stack, startV)

	for len(stack) > 0 {
		currV := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if visited[currV] {
			continue
		}

		visited[currV] = true

		// Получаем соседей текущей вершины
		neighbors := graph[currV]

		// Добавляем соседей в стек
		for _, neighbor := range neighbors {
			if !visited[neighbor] {
				stack = append(stack, neighbor)
			}
		}
	}
}

func main() {
	graph := map[string][]string{
		"A": {"B", "C"},
		"B": {"D", "E"},
		"C": {"F"},
		"D": {},
		"E": {"F"},
		"F": {},
	}

	startV := "A"
	dfs(graph, startV)
}
