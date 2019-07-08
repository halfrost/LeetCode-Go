package leetcode

var dir = [][]int{
	[]int{-1, 0},
	[]int{0, 1},
	[]int{1, 0},
	[]int{0, -1},
}

func exist(board [][]byte, word string) bool {
	visited := make([][]bool, len(board))
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]bool, len(board[0]))
	}
	for i, v := range board {
		for j := range v {
			if searchWord(board, visited, word, 0, i, j) {
				return true
			}
		}
	}
	return false
}

func isInBoard(board [][]byte, x, y int) bool {
	return x >= 0 && x < len(board) && y >= 0 && y < len(board[0])
}

func searchWord(board [][]byte, visited [][]bool, word string, index, x, y int) bool {
	if index == len(word)-1 {
		return board[x][y] == word[index]
	}
	if board[x][y] == word[index] {
		visited[x][y] = true
		for i := 0; i < 4; i++ {
			nx := x + dir[i][0]
			ny := y + dir[i][1]
			if isInBoard(board, nx, ny) && !visited[nx][ny] && searchWord(board, visited, word, index+1, nx, ny) {
				return true
			}
		}
		visited[x][y] = false
	}
	return false
}
