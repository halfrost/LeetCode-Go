package leetcode

var dir = [][]int{
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1},
}

func findPaths(m int, n int, maxMove int, startRow int, startColumn int) int {
	visited := make([][][]int, m)
	for i := range visited {
		visited[i] = make([][]int, n)
		for j := range visited[i] {
			visited[i][j] = make([]int, maxMove+1)
			for l := range visited[i][j] {
				visited[i][j][l] = -1
			}
		}
	}
	return dfs(startRow, startColumn, maxMove, m, n, visited)
}

func dfs(x, y, maxMove, m, n int, visited [][][]int) int {
	if x < 0 || x >= m || y < 0 || y >= n {
		return 1
	}
	if maxMove == 0 {
		visited[x][y][maxMove] = 0
		return 0
	}
	if visited[x][y][maxMove] >= 0 {
		return visited[x][y][maxMove]
	}
	res := 0
	for i := 0; i < 4; i++ {
		nx := x + dir[i][0]
		ny := y + dir[i][1]
		res += (dfs(nx, ny, maxMove-1, m, n, visited) % 1000000007)
	}
	visited[x][y][maxMove] = res % 1000000007
	return visited[x][y][maxMove]
}
