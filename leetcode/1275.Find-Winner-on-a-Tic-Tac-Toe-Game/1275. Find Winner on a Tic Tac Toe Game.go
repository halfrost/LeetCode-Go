package leetcode

func tictactoe(moves [][]int) string {
	board := [3][3]byte{}
	for i := 0; i < len(moves); i++ {
		if i%2 == 0 {
			board[moves[i][0]][moves[i][1]] = 'X'
		} else {
			board[moves[i][0]][moves[i][1]] = 'O'
		}
	}
	for i := 0; i < 3; i++ {
		if board[i][0] == 'X' && board[i][1] == 'X' && board[i][2] == 'X' {
			return "A"
		}
		if board[i][0] == 'O' && board[i][1] == 'O' && board[i][2] == 'O' {
			return "B"
		}
		if board[0][i] == 'X' && board[1][i] == 'X' && board[2][i] == 'X' {
			return "A"
		}
		if board[0][i] == 'O' && board[1][i] == 'O' && board[2][i] == 'O' {
			return "B"
		}
	}
	if board[0][0] == 'X' && board[1][1] == 'X' && board[2][2] == 'X' {
		return "A"
	}
	if board[0][0] == 'O' && board[1][1] == 'O' && board[2][2] == 'O' {
		return "B"
	}
	if board[0][2] == 'X' && board[1][1] == 'X' && board[2][0] == 'X' {
		return "A"
	}
	if board[0][2] == 'O' && board[1][1] == 'O' && board[2][0] == 'O' {
		return "B"
	}
	if len(moves) < 9 {
		return "Pending"
	}
	return "Draw"
}
