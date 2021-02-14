package leetcode

var dir = [][]int{
	{-1, -1},
	{-1, 0},
	{-1, 1},
	{0, 1},
	{0, -1},
	{1, -1},
	{1, 0},
	{1, 1},
}

func shortestPathBinaryMatrix(grid [][]int) int {
	visited := make([][]bool, 0)
	for range make([]int, len(grid)) {
		visited = append(visited, make([]bool, len(grid[0])))
	}
	dis := make([][]int, 0)
	for range make([]int, len(grid)) {
		dis = append(dis, make([]int, len(grid[0])))
	}
	if grid[0][0] == 1 {
		return -1
	}
	if len(grid) == 1 && len(grid[0]) == 1 {
		return 1
	}

	queue := []int{0}
	visited[0][0], dis[0][0] = true, 1
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		curx, cury := cur/len(grid[0]), cur%len(grid[0])
		for d := 0; d < 8; d++ {
			nextx := curx + dir[d][0]
			nexty := cury + dir[d][1]
			if isInBoard(grid, nextx, nexty) && !visited[nextx][nexty] && grid[nextx][nexty] == 0 {
				queue = append(queue, nextx*len(grid[0])+nexty)
				visited[nextx][nexty] = true
				dis[nextx][nexty] = dis[curx][cury] + 1
				if nextx == len(grid)-1 && nexty == len(grid[0])-1 {
					return dis[nextx][nexty]
				}
			}
		}
	}
	return -1
}

func isInBoard(board [][]int, x, y int) bool {
	return x >= 0 && x < len(board) && y >= 0 && y < len(board[0])
}
