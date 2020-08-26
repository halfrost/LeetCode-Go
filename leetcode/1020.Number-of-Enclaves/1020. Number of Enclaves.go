package leetcode

var dir = [][]int{
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1},
}

func numEnclaves(A [][]int) int {
	m, n := len(A), len(A[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || i == m-1 || j == 0 || j == n-1 {
				if A[i][j] == 1 {
					dfsNumEnclaves(A, i, j)
				}
			}
		}
	}
	count := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if A[i][j] == 1 {
				count++
			}
		}
	}
	return count
}

func dfsNumEnclaves(A [][]int, x, y int) {
	if !isInGrid(A, x, y) || A[x][y] == 0 {
		return
	}
	A[x][y] = 0
	for i := 0; i < 4; i++ {
		nx := x + dir[i][0]
		ny := y + dir[i][1]
		dfsNumEnclaves(A, nx, ny)
	}
}

func isInGrid(grid [][]int, x, y int) bool {
	return x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0])
}
