package main

import "fmt"

// Алгоритм Breadth-First Search: начиная с начальной вершины startVertex,
// выводит порядок посещения вершин и возвращает массив расстояний
// от стартовой вершины до каждой другой вершины.

func bfs(graph map[int][]int, startVertex int) ([]bool, []int, []int) {

	visited := make([]bool, len(graph))    // Массив посещённых вершин
	distance := make([]int, len(graph))    // Массив расстояний от startVertex
	vertexes := make([]int, 0, len(graph)) // Узлы которые обошли

	// Изначально расстояние бесконечно (-1)
	for i := range distance {
		distance[i] = -1
	}

	// Очередь
	queue := []int{startVertex}

	visited[startVertex] = true
	distance[startVertex] = 0 // Расстояние от себя самого равно нулю

	//
	for len(queue) > 0 {
		valueVertex := queue[0]
		queue = queue[1:]

		//
		vertexes = append(vertexes, valueVertex)

		// Обход соседей
		neighbors := graph[valueVertex]
		for _, neighbor := range neighbors {
			if !visited[neighbor] {

				visited[neighbor] = true
				distance[neighbor] = distance[valueVertex] + 1
				queue = append(queue, neighbor)
			}
		}
	}
	return visited, distance, vertexes
}

func main() {
	graph := map[int][]int{
		0: {1, 2},
		1: {0, 3, 4},
		2: {0, 5, 6},
		3: {1, 7},
		4: {1, 7},
		5: {2},
		6: {2},
		7: {3, 4},
	}

	visited, distance, vertexes := bfs(graph, 2)

	fmt.Println(visited, distance, vertexes)
}
