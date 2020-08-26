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

	mineMap := make([][]int, len(board))
	for i := range board {
		mineMap[i] = make([]int, len(board[i]))
	}

	for i := range board {
		for j := range board[i] {
			if board[i][j] == 'M' {
				mineMap[i][j] = -1
				for _, d := range dir8 {
					nx, ny := i+d[0], j+d[1]
					if isInBoard(board, nx, ny) && mineMap[nx][ny] >= 0 {
						mineMap[nx][ny]++
					}
				}
			}
		}
	}
	mineSweeper(click[0], click[1], board, mineMap, dir8)
	return board
}

func isInBoard(board [][]byte, x, y int) bool {
	return x >= 0 && x < len(board) && y >= 0 && y < len(board[0])
}

func mineSweeper(x, y int, board [][]byte, mineMap [][]int, dir8 [][]int) {
	if board[x][y] != 'M' && board[x][y] != 'E' {
		return
	}
	if mineMap[x][y] == -1 {
		board[x][y] = 'X'
	} else if mineMap[x][y] > 0 {
		board[x][y] = '0' + byte(mineMap[x][y])
	} else {
		board[x][y] = 'B'
		for _, d := range dir8 {
			nx, ny := x+d[0], y+d[1]
			if isInBoard(board, nx, ny) && mineMap[nx][ny] >= 0 {
				mineSweeper(nx, ny, board, mineMap, dir8)
			}
		}
	}
}
