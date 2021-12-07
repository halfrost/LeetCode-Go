package leetcode

type point struct {
	x int
	y int
}

var (
	borders       []point
	m             int
	n             int
	vis           [][]bool
	dirs          []point
	q             []point
	originalColor int
)

func colorBorder(grid [][]int, row, col, color int) [][]int {
	m, n = len(grid), len(grid[0])
	vis = make([][]bool, m)
	for i := range vis {
		vis[i] = make([]bool, n)
	}
	dirs = []point{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	originalColor = grid[row][col]
	borders = []point{}
	q = []point{{row, col}}
	bfs(q, grid)
	for _, p := range borders {
		grid[p.x][p.y] = color
	}
	return grid
}

func bfs(q []point, grid [][]int) {
	for len(q) != 0 {
		ele := q[0]
		vis[ele.x][ele.y] = true
		q = q[1:]
		isBorder := false
		for _, dir := range dirs {
			nx, ny := ele.x+dir.x, ele.y+dir.y
			if !(0 <= nx && nx < m && 0 <= ny && ny < n && grid[nx][ny] == originalColor) {
				isBorder = true
			} else if !vis[nx][ny] {
				q = append(q, point{nx, ny})
			}
		}
		if isBorder {
			borders = append(borders, point{ele.x, ele.y})
		}
	}
}
