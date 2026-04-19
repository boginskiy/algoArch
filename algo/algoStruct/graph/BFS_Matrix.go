package main

import "fmt"

func bfs2(matrix [][]int, compare map[int]string, startVertex int) map[string]int {
	visited := make([]bool, len(matrix))          // Массив посещённых вершин
	distance := make(map[string]int, len(matrix)) // Массив расстояний от startVertex

	queue := []int{startVertex}        // Очередь для вершин
	visited[startVertex] = true        // Отмечаем, что начали с этой вершины
	distance[compare[startVertex]] = 0 // Расстояние от себя самого равно нулю

	for len(queue) > 0 {
		currenVertex := queue[0] // Получаем индекс текущего узла
		queue = queue[1:]

		// Просматриваем строку матрицы смежности, соответствующую вершине frontNode
		for nextVertex := 0; nextVertex < len(matrix); nextVertex++ {
			if matrix[currenVertex][nextVertex] == 1 && !visited[nextVertex] {

				visited[nextVertex] = true                                          // Метим новую вершину как посещённую
				distance[compare[nextVertex]] = distance[compare[currenVertex]] + 1 // Расстояние
				queue = append(queue, nextVertex)                                   // Добавляем её в очередь
			}
		}
	}
	return distance
}

func main() {

	// Матрица смежности представляет собой двунаправленную структуру связей между вершинами
	var Matrix = [][]int{
		{0, 1, 0, 0}, // A
		{0, 0, 1, 0}, // B
		{0, 0, 0, 1}, // C
		{0, 0, 0, 0}, // D
	}

	// Число вершин в нашем примере
	const verticesCount = 4

	compare := map[int]string{0: "A", 1: "B", 2: "C", 3: "D"}

	distance := bfs2(Matrix, compare, 1) // Начальная вершина (A)
	fmt.Println(distance)
}
