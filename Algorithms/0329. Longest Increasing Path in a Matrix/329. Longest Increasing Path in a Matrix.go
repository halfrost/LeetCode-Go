package leetcode

import (
	"math"
)

func longestIncreasingPath(matrix [][]int) int {
	cache, res := make([][]int, len(matrix)), 0
	for i := 0; i < len(cache); i++ {
		cache[i] = make([]int, len(matrix[0]))
	}
	for i, v := range matrix {
		for j := range v {
			searchPath(matrix, cache, math.MinInt64, i, j)
			res = max(res, cache[i][j])
		}
	}
	return res
}

func isInIntBoard(board [][]int, x, y int) bool {
	return x >= 0 && x < len(board) && y >= 0 && y < len(board[0])
}

func searchPath(board, cache [][]int, lastNum, x, y int) int {
	if board[x][y] <= lastNum {
		return 0
	}
	if cache[x][y] > 0 {
		return cache[x][y]
	}
	count := 1
	for i := 0; i < 4; i++ {
		nx := x + dir[i][0]
		ny := y + dir[i][1]
		if isInIntBoard(board, nx, ny) {
			count = max(count, searchPath(board, cache, board[x][y], nx, ny)+1)
		}
	}
	cache[x][y] = count
	return count
}
