package main

import "fmt"

func bfs3(matrix [][]int, vertex int, mapper map[int]string) map[string]int {
	visited := make([]bool, len(matrix))
	distance := make(map[string]int)

	queue := []int{vertex}
	visited[vertex] = true
	distance[mapper[vertex]] = 0

	for len(queue) > 0 {
		currV := queue[0]
		queue = queue[1:]

		for i, neighbour := range matrix[currV] {
			if neighbour == 1 && !visited[i] {
				queue = append(queue, i)
				visited[i] = true
				distance[mapper[i]] = distance[mapper[currV]] + 1
			}
		}
	}
	return distance
}

func main() {

	var Matrix = [][]int{
		{0, 0, 1, 1}, // A
		{0, 0, 0, 0}, // B
		{0, 1, 0, 1}, // C
		{0, 1, 0, 0}, // D
	}

	Mapper := map[int]string{0: "A", 1: "B", 2: "C", 3: "D"}

	distance := bfs3(Matrix, 0, Mapper)
	fmt.Println(distance)

}
