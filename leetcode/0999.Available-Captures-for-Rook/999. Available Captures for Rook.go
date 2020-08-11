package leetcode

func numRookCaptures(board [][]byte) int {
	num := 0
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == 'R' {
				num += caputure(board, i-1, j, -1, 0) // Up
				num += caputure(board, i+1, j, 1, 0)  // Down
				num += caputure(board, i, j-1, 0, -1) // Left
				num += caputure(board, i, j+1, 0, 1)  // Right
			}
		}
	}
	return num
}

func caputure(board [][]byte, x, y int, bx, by int) int {
	for x >= 0 && x < len(board) && y >= 0 && y < len(board[x]) && board[x][y] != 'B' {
		if board[x][y] == 'p' {
			return 1
		}
		x += bx
		y += by
	}
	return 0
}
