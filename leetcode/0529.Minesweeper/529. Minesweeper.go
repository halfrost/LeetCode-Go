package leetcode

var dir8 = [][]int{
	{-1, -1},
	{-1, 0},
	{-1, 1},
	{0, 1},
	{1, 1},
	{1, 0},
	{1, -1},
	{0, -1},
}

func updateBoard(board [][]byte, click []int) [][]byte {
	if board[click[0]][click[1]] == 'M' {
		board[click[0]][click[1]] = 'X'
		return board
	}
	dfs(board, click[0], click[1])
	return board
}

func dfs(board [][]byte, x, y int) {
	cnt := 0
	for i := 0; i < 8; i++ {
		nx, ny := x+dir8[i][0], y+dir8[i][1]
		if isInBoard(board, nx, ny) && board[nx][ny] == 'M' {
			cnt++
		}
	}
	if cnt > 0 {
		board[x][y] = byte(cnt + '0')
		return
	}
	board[x][y] = 'B'
	for i := 0; i < 8; i++ {
		nx, ny := x+dir8[i][0], y+dir8[i][1]
		if isInBoard(board, nx, ny) && board[nx][ny] != 'B' {
			dfs(board, nx, ny)
		}
	}
}

func isInBoard(board [][]byte, x, y int) bool {
	return x >= 0 && x < len(board) && y >= 0 && y < len(board[0])
}
