package leetcode

import "math"

func pacificAtlantic(matrix [][]int) [][]int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil
	}
	row, col, res := len(matrix), len(matrix[0]), make([][]int, 0)
	pacific, atlantic := make([][]bool, row), make([][]bool, row)
	for i := 0; i < row; i++ {
		pacific[i] = make([]bool, col)
		atlantic[i] = make([]bool, col)
	}
	for i := 0; i < row; i++ {
		dfs(matrix, i, 0, &pacific, math.MinInt32)
		dfs(matrix, i, col-1, &atlantic, math.MinInt32)
	}
	for j := 0; j < col; j++ {
		dfs(matrix, 0, j, &pacific, math.MinInt32)
		dfs(matrix, row-1, j, &atlantic, math.MinInt32)
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if atlantic[i][j] && pacific[i][j] {
				res = append(res, []int{i, j})
			}
		}
	}
	return res
}

func dfs(matrix [][]int, row, col int, visited *[][]bool, height int) {
	if row < 0 || row >= len(matrix) || col < 0 || col >= len(matrix[0]) {
		return
	}
	if (*visited)[row][col] || matrix[row][col] < height {
		return
	}
	(*visited)[row][col] = true
	dfs(matrix, row+1, col, visited, matrix[row][col])
	dfs(matrix, row-1, col, visited, matrix[row][col])
	dfs(matrix, row, col+1, visited, matrix[row][col])
	dfs(matrix, row, col-1, visited, matrix[row][col])
}
