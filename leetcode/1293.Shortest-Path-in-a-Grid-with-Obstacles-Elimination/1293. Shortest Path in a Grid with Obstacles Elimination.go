package leetcode

var dir = [][]int{
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1},
}

type pos struct {
	x, y     int
	obstacle int
	step     int
}

func shortestPath(grid [][]int, k int) int {
	queue, m, n := []pos{}, len(grid), len(grid[0])
	visitor := make([][][]int, m)
	if len(grid) == 1 && len(grid[0]) == 1 {
		return 0
	}
	for i := 0; i < m; i++ {
		visitor[i] = make([][]int, n)
		for j := 0; j < n; j++ {
			visitor[i][j] = make([]int, k+1)
		}
	}
	visitor[0][0][0] = 1
	queue = append(queue, pos{x: 0, y: 0, obstacle: 0, step: 0})
	for len(queue) > 0 {
		size := len(queue)
		for size > 0 {
			size--
			node := queue[0]
			queue = queue[1:]
			for i := 0; i < len(dir); i++ {
				newX := node.x + dir[i][0]
				newY := node.y + dir[i][1]
				if newX == m-1 && newY == n-1 {
					if node.obstacle != 0 {
						if node.obstacle <= k {
							return node.step + 1
						} else {
							continue
						}
					}
					return node.step + 1
				}
				if isInBoard(grid, newX, newY) {
					if grid[newX][newY] == 1 {
						if node.obstacle+1 <= k && visitor[newX][newY][node.obstacle+1] != 1 {
							queue = append(queue, pos{x: newX, y: newY, obstacle: node.obstacle + 1, step: node.step + 1})
							visitor[newX][newY][node.obstacle+1] = 1
						}
					} else {
						if node.obstacle <= k && visitor[newX][newY][node.obstacle] != 1 {
							queue = append(queue, pos{x: newX, y: newY, obstacle: node.obstacle, step: node.step + 1})
							visitor[newX][newY][node.obstacle] = 1
						}
					}

				}
			}
		}
	}
	return -1
}

func isInBoard(board [][]int, x, y int) bool {
	return x >= 0 && x < len(board) && y >= 0 && y < len(board[0])
}
