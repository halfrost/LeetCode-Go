package leetcode

func uniquePathsIII(grid [][]int) int {
	visited := make([][]bool, len(grid))
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]bool, len(grid[0]))
	}
	res, empty, startx, starty, endx, endy, path := 0, 0, 0, 0, 0, 0, []int{}
	for i, v := range grid {
		for j, vv := range v {
			switch vv {
			case 0:
				empty++
			case 1:
				startx, starty = i, j
			case 2:
				endx, endy = i, j
			}
		}
	}
	findUniquePathIII(grid, visited, path, empty+1, startx, starty, endx, endy, &res) // 可走的步数要加一，因为终点格子也算一步，不然永远走不到终点！
	return res
}

func isInPath(board [][]int, x, y int) bool {
	return x >= 0 && x < len(board) && y >= 0 && y < len(board[0])
}

func findUniquePathIII(board [][]int, visited [][]bool, path []int, empty, startx, starty, endx, endy int, res *int) {
	if startx == endx && starty == endy {
		if empty == 0 {
			*res++
		}
		return
	}
	if board[startx][starty] >= 0 {
		visited[startx][starty] = true
		empty--
		path = append(path, startx)
		path = append(path, starty)
		for i := 0; i < 4; i++ {
			nx := startx + dir[i][0]
			ny := starty + dir[i][1]
			if isInPath(board, nx, ny) && !visited[nx][ny] {
				findUniquePathIII(board, visited, path, empty, nx, ny, endx, endy, res)
			}
		}
		visited[startx][starty] = false
		empty++
		path = path[:len(path)-2]
	}
	return
}
