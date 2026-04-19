package main

import "fmt"

func dfs(matrix [][]byte, row, col int) {
	matrix[row][col] = '#'
	maxRows, maxCols := len(matrix), len(matrix[row])

	stack := [][]int{{row, col}}

	for len(stack) > 0 {
		currV := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		for _, shift := range [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}} {
			shiftRow := currV[0] + shift[0]
			shiftCol := currV[1] + shift[1]

			if shiftRow >= 0 && shiftRow < maxRows &&
				shiftCol >= 0 && shiftCol < maxCols &&
				matrix[shiftRow][shiftCol] == '1' {
				matrix[shiftRow][shiftCol] = '#'
				stack = append(stack, []int{shiftRow, shiftCol})
			}
		}
	}
}

func bfs(matrix [][]byte, row, col int) {
	matrix[row][col] = '#'
	maxRows, maxCols := len(matrix), len(matrix[row])

	queue := [][]int{{row, col}}

	for len(queue) > 0 {
		currElem := queue[0]
		queue = queue[1:]

		for _, shift := range [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}} {
			shiftRow := currElem[0] + shift[0]
			shiftCol := currElem[1] + shift[1]

			if (shiftRow >= 0 && shiftRow < maxRows) &&
				(shiftCol >= 0 && shiftCol < maxCols) &&
				matrix[shiftRow][shiftCol] == '1' {

				matrix[shiftRow][shiftCol] = '#'
				queue = append(queue, []int{shiftRow, shiftCol})
			}
		}
	}
}

func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	rows, cols := len(grid), len(grid[0])
	cnt := 0

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '1' {
				// bfs(grid, i, j)
				dfs(grid, i, j)
				cnt++
			}
		}
	}
	return cnt
}

func main() {
	grid := [][]byte{
		{'1', '1', '1', '1'},
		{'1', '1', '0', '1'},
		{'1', '0', '0', '0'},
		{'0', '1', '0', '1'}}

	res := numIslands(grid)
	fmt.Println(res)
}

// Мое решение
// func bfs(matrix [][]byte, row, col int) {
// 	matrix[row][col] = '#'
// 	maxRows := len(matrix)
// 	maxCols := len(matrix[row])

// 	queue := make([][]int, 0, 10)
// 	queue = append(queue, []int{row, col})

// 	for len(queue) > 0 {
// 		currElem := queue[0]
// 		queue = queue[1:]

// 		for _, shift := range [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}} {
// 			shiftRow := currElem[0] + shift[0]
// 			shiftCol := currElem[1] + shift[1]

// 			if (shiftRow >= 0 && shiftRow < maxRows) &&
// 				(shiftCol >= 0 && shiftCol < maxCols) &&
// 				matrix[shiftRow][shiftCol] == '1' {

// 				matrix[shiftRow][shiftCol] = '#'
// 				queue = append(queue, []int{shiftRow, shiftCol})
// 			}
// 		}
// 	}
// }

// func numIslands(grid [][]byte) int {
// 	if len(grid) == 0 || len(grid[0]) == 0 {
// 		return 0
// 	}
// 	rows, cols := len(grid), len(grid[0])
// 	cnt := 0
// 	for i := 0; i < rows; i++ {
// 		for j := 0; j < cols; j++ {
// 			if grid[i][j] == '1' {
// 				bfs(grid, i, j)
// 				cnt++
// 			}
// 		}
// 	}
// 	return cnt
// }
