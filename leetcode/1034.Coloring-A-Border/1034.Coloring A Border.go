package leetcode

type point struct {
	x int
	y int
}

type gridInfo struct {
	m             int
	n             int
	grid          [][]int
	originalColor int
}

func colorBorder(grid [][]int, row, col, color int) [][]int {
	m, n := len(grid), len(grid[0])
	dirs := []point{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	vis := make([][]bool, m)
	for i := range vis {
		vis[i] = make([]bool, n)
	}
	var borders []point
	gInfo := gridInfo{
		m:             m,
		n:             n,
		grid:          grid,
		originalColor: grid[row][col],
	}
	dfs(row, col, gInfo, dirs, vis, &borders)
	for _, p := range borders {
		grid[p.x][p.y] = color
	}
	return grid
}

func dfs(x, y int, gInfo gridInfo, dirs []point, vis [][]bool, borders *[]point) {
	vis[x][y] = true
	isBorder := false
	for _, dir := range dirs {
		nx, ny := x+dir.x, y+dir.y
		if !(0 <= nx && nx < gInfo.m && 0 <= ny && ny < gInfo.n && gInfo.grid[nx][ny] == gInfo.originalColor) {
			isBorder = true
		} else if !vis[nx][ny] {
			dfs(nx, ny, gInfo, dirs, vis, borders)
		}
	}
	if isBorder {
		*borders = append(*borders, point{x, y})
	}
}
