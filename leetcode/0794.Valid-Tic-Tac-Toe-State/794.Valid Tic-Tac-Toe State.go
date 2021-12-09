package leetcode

func validTicTacToe(board []string) bool {
	cntX, cntO := 0, 0
	for i := range board {
		for j := range board[i] {
			if board[i][j] == 'X' {
				cntX++
			} else if board[i][j] == 'O' {
				cntO++
			}
		}
	}
	if cntX < cntO || cntX > cntO+1 {
		return false
	}
	if cntX == cntO {
		return process(board, 'X')
	}
	return process(board, 'O')
}

func process(board []string, c byte) bool {
	//某一行是"ccc"
	if board[0] == string([]byte{c, c, c}) || board[1] == string([]byte{c, c, c}) || board[2] == string([]byte{c, c, c}) {
		return false
	}
	//某一列是"ccc"
	if (board[0][0] == c && board[1][0] == c && board[2][0] == c) ||
		(board[0][1] == c && board[1][1] == c && board[2][1] == c) ||
		(board[0][2] == c && board[1][2] == c && board[2][2] == c) {
		return false
	}
	//某一对角线是"ccc"
	if (board[0][0] == c && board[1][1] == c && board[2][2] == c) ||
		(board[0][2] == c && board[1][1] == c && board[2][0] == c) {
		return false
	}
	return true
}
