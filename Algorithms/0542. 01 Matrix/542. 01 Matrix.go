package leetcode

import (
	"math"
)

// 解法一 BFS
func updateMatrixBFS(matrix [][]int) [][]int {
	res := make([][]int, len(matrix))
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return res
	}
	queue := make([][]int, 0)
	for i := range matrix {
		res[i] = make([]int, len(matrix[0]))
		for j := range res[i] {
			if matrix[i][j] == 0 {
				res[i][j] = -1
				queue = append(queue, []int{i, j})
			}
		}
	}
	level := 1
	for len(queue) > 0 {
		size := len(queue)
		for size > 0 {
			size--
			node := queue[0]
			queue = queue[1:]
			i, j := node[0], node[1]
			for _, direction := range [][]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}} {
				x := i + direction[0]
				y := j + direction[1]
				if x < 0 || x >= len(matrix) || y < 0 || y >= len(matrix[0]) || res[x][y] < 0 || res[x][y] > 0 {
					continue
				}
				res[x][y] = level
				queue = append(queue, []int{x, y})
			}
		}
		level++
	}
	for i, row := range res {
		for j, cell := range row {
			if cell == -1 {
				res[i][j] = 0
			}
		}
	}
	return res
}

// 解法二 DFS
func updateMatrixDFS(matrix [][]int) [][]int {
	result := [][]int{}
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return result
	}
	maxRow, maxCol := len(matrix), len(matrix[0])
	for r := 0; r < maxRow; r++ {
		for c := 0; c < maxCol; c++ {
			if matrix[r][c] == 1 && hasZero(matrix, r, c) == false {
				// 将四周没有 0 的 1 特殊处理为最大值
				matrix[r][c] = math.MaxInt64
			}
		}
	}
	for r := 0; r < maxRow; r++ {
		for c := 0; c < maxCol; c++ {
			if matrix[r][c] == 1 {
				dfsMatrix(matrix, r, c, -1)
			}
		}
	}
	return (matrix)
}

// 判断四周是否有 0
func hasZero(matrix [][]int, row, col int) bool {
	if row > 0 && matrix[row-1][col] == 0 {
		return true
	}
	if col > 0 && matrix[row][col-1] == 0 {
		return true
	}
	if row < len(matrix)-1 && matrix[row+1][col] == 0 {
		return true
	}
	if col < len(matrix[0])-1 && matrix[row][col+1] == 0 {
		return true
	}
	return false
}

func dfsMatrix(matrix [][]int, row, col, val int) {
	// 不超过棋盘氛围，且 val 要比 matrix[row][col] 小
	if row < 0 || row >= len(matrix) || col < 0 || col >= len(matrix[0]) || (matrix[row][col] <= val) {
		return
	}
	if val > 0 {
		matrix[row][col] = val
	}
	dfsMatrix(matrix, row-1, col, matrix[row][col]+1)
	dfsMatrix(matrix, row, col-1, matrix[row][col]+1)
	dfsMatrix(matrix, row+1, col, matrix[row][col]+1)
	dfsMatrix(matrix, row, col+1, matrix[row][col]+1)
}

// 解法三 DP
func updateMatrixDP(matrix [][]int) [][]int {
	for i, row := range matrix {
		for j, val := range row {
			if val == 0 {
				continue
			}
			left, top := math.MaxInt16, math.MaxInt16
			if i > 0 {
				top = matrix[i-1][j] + 1
			}
			if j > 0 {
				left = matrix[i][j-1] + 1
			}
			matrix[i][j] = min(top, left)
		}
	}
	for i := len(matrix) - 1; i >= 0; i-- {
		for j := len(matrix[0]) - 1; j >= 0; j-- {
			if matrix[i][j] == 0 {
				continue
			}
			right, bottom := math.MaxInt16, math.MaxInt16
			if i < len(matrix)-1 {
				bottom = matrix[i+1][j] + 1
			}
			if j < len(matrix[0])-1 {
				right = matrix[i][j+1] + 1
			}
			matrix[i][j] = min(matrix[i][j], min(bottom, right))
		}
	}
	return matrix
}
