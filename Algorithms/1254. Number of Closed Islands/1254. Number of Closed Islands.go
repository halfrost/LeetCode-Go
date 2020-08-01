package leetcode

func closedIsland(grid [][]int) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])
	if n == 0 {
		return 0
	}
	res, visited := 0, make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			isEdge := false
			if grid[i][j] == 0 && !visited[i][j] {
				checkIslands(grid, &visited, i, j, &isEdge)
				if !isEdge {
					res++
				}

			}
		}
	}
	return res
}

func checkIslands(grid [][]int, visited *[][]bool, x, y int, isEdge *bool) {
	if (x == 0 || x == len(grid)-1 || y == 0 || y == len(grid[0])-1) && grid[x][y] == 0 {
		*isEdge = true
	}
	(*visited)[x][y] = true
	for i := 0; i < 4; i++ {
		nx := x + dir[i][0]
		ny := y + dir[i][1]
		if isIntInBoard(grid, nx, ny) && !(*visited)[nx][ny] && grid[nx][ny] == 0 {
			checkIslands(grid, visited, nx, ny, isEdge)
		}
	}
	*isEdge = *isEdge || false
}

func isIntInBoard(board [][]int, x, y int) bool {
	return x >= 0 && x < len(board) && y >= 0 && y < len(board[0])
}
