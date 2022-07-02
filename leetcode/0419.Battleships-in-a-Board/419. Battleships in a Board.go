package leetcode

func countBattleships(board [][]byte) (ans int) {
	if len(board) == 0 || len(board[0]) == 0 {
		return 0
	}
	for i := range board {
		for j := range board[i] {
			if board[i][j] == 'X' {
				if i > 0 && board[i-1][j] == 'X' {
					continue
				}
				if j > 0 && board[i][j-1] == 'X' {
					continue
				}
				ans++
			}
		}
	}
	return
}
