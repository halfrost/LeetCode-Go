package leetcode

import (
	"math"
	"strings"
)

// 解法一 BFS，利用状态压缩来过滤筛选状态
func shortestPathAllKeys(grid []string) int {
	if len(grid) == 0 {
		return 0
	}
	board, visited, startx, starty, res, fullKeys := make([][]byte, len(grid)), make([][][]bool, len(grid)), 0, 0, 0, 0
	for i := 0; i < len(grid); i++ {
		board[i] = make([]byte, len(grid[0]))
	}
	for i, g := range grid {
		board[i] = []byte(g)
		for _, v := range g {
			if v == 'a' || v == 'b' || v == 'c' || v == 'd' || v == 'e' || v == 'f' {
				fullKeys |= (1 << uint(v-'a'))
			}
		}
		if strings.Contains(g, "@") {
			startx, starty = i, strings.Index(g, "@")
		}
	}
	for i := 0; i < len(visited); i++ {
		visited[i] = make([][]bool, len(board[0]))
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			visited[i][j] = make([]bool, 64)
		}
	}
	queue := []int{}
	queue = append(queue, (starty<<16)|(startx<<8))
	visited[startx][starty][0] = true
	for len(queue) != 0 {
		qLen := len(queue)
		for i := 0; i < qLen; i++ {
			state := queue[0]
			queue = queue[1:]
			starty, startx = state>>16, (state>>8)&0xFF
			keys := state & 0xFF
			if keys == fullKeys {
				return res
			}
			for i := 0; i < 4; i++ {
				newState := keys
				nx := startx + dir[i][0]
				ny := starty + dir[i][1]
				if !isInBoard(board, nx, ny) {
					continue
				}
				if board[nx][ny] == '#' {
					continue
				}
				flag, canThroughLock := keys&(1<<(board[nx][ny]-'A')), false
				if flag != 0 {
					canThroughLock = true
				}
				if isLock(board, nx, ny) && !canThroughLock {
					continue
				}
				if isKey(board, nx, ny) {
					newState |= (1 << (board[nx][ny] - 'a'))
				}
				if visited[nx][ny][newState] {
					continue
				}
				queue = append(queue, (ny<<16)|(nx<<8)|newState)
				visited[nx][ny][newState] = true
			}
		}
		res++
	}
	return -1
}

// 解法二 DFS，但是超时了，剪枝条件不够强
func shortestPathAllKeys1(grid []string) int {
	if len(grid) == 0 {
		return 0
	}
	board, visited, startx, starty, res, fullKeys := make([][]byte, len(grid)), make([][][]bool, len(grid)), 0, 0, math.MaxInt64, 0
	for i := 0; i < len(grid); i++ {
		board[i] = make([]byte, len(grid[0]))
	}
	for i, g := range grid {
		board[i] = []byte(g)
		for _, v := range g {
			if v == 'a' || v == 'b' || v == 'c' || v == 'd' || v == 'e' || v == 'f' {
				fullKeys |= (1 << uint(v-'a'))
			}
		}
		if strings.Contains(g, "@") {
			startx, starty = i, strings.Index(g, "@")
		}
	}
	for i := 0; i < len(visited); i++ {
		visited[i] = make([][]bool, len(board[0]))
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			visited[i][j] = make([]bool, 64)
		}
	}
	searchKeys(board, &visited, fullKeys, 0, (starty<<16)|(startx<<8), &res, []int{})
	if res == math.MaxInt64 {
		return -1
	}
	return res - 1
}

func searchKeys(board [][]byte, visited *[][][]bool, fullKeys, step, state int, res *int, path []int) {
	y, x := state>>16, (state>>8)&0xFF
	keys := state & 0xFF

	if keys == fullKeys {
		*res = min(*res, step)
		return
	}

	flag, canThroughLock := keys&(1<<(board[x][y]-'A')), false
	if flag != 0 {
		canThroughLock = true
	}
	newState := keys
	//fmt.Printf("x = %v y = %v fullKeys = %v keys = %v step = %v res = %v path = %v state = %v\n", x, y, fullKeys, keys, step, *res, path, state)
	if (board[x][y] != '#' && !isLock(board, x, y)) || (isLock(board, x, y) && canThroughLock) {
		if isKey(board, x, y) {
			newState |= (1 << uint(board[x][y]-'a'))
		}
		(*visited)[x][y][newState] = true
		path = append(path, x)
		path = append(path, y)

		for i := 0; i < 4; i++ {
			nx := x + dir[i][0]
			ny := y + dir[i][1]
			if isInBoard(board, nx, ny) && !(*visited)[nx][ny][newState] {
				searchKeys(board, visited, fullKeys, step+1, (ny<<16)|(nx<<8)|newState, res, path)
			}
		}
		(*visited)[x][y][keys] = false
		path = path[:len(path)-1]
		path = path[:len(path)-1]
	}
}

func isLock(board [][]byte, x, y int) bool {
	if (board[x][y] == 'A') || (board[x][y] == 'B') ||
		(board[x][y] == 'C') || (board[x][y] == 'D') ||
		(board[x][y] == 'E') || (board[x][y] == 'F') {
		return true
	}
	return false
}

func isKey(board [][]byte, x, y int) bool {
	if (board[x][y] == 'a') || (board[x][y] == 'b') ||
		(board[x][y] == 'c') || (board[x][y] == 'd') ||
		(board[x][y] == 'e') || (board[x][y] == 'f') {
		return true
	}
	return false
}
