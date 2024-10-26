package leetcode

var dir = [][]int{
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1},
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
changes:
Base Case Check: In your searchWord function, you should check if index is equal to len(word) - 1, but also need to ensure that the character at index matches the character at the current position before returning true. You should adjust this condition to account for reaching the end of the word correctly.

Visited Check: You should also ensure that the last character in the word is checked correctly.

Here’s the corrected version of your code:

go
Copy code
package leetcode

var dir = [][]int{
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1},
}

func exist(board [][]byte, word string) bool {
	visited := make([][]bool, len(board))
	for i := range visited {
		visited[i] = make([]bool, len(board[0]))
	}
	for i := range board {
		for j := range board[i] {
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
	// Check if we have matched the entire word
	if index == len(word) {
		return true
	}
	// Check boundaries and if the cell is already visited or doesn't match the word
	if !isInBoard(board, x, y) || visited[x][y] || board[x][y] != word[index] {
		return false
	}

	// Mark the cell as visited
	visited[x][y] = true

	// Explore all 4 directions
	for i := 0; i < 4; i++ {
		nx := x + dir[i][0]
		ny := y + dir[i][1]
		if searchWord(board, visited, word, index+1, nx, ny) {
			return true
		}
	}

	// Backtrack: unmark the cell as visited
	visited[x][y] = false
	return false
}
Key Changes:
Base Case Update: The base case now checks if index is equal to len(word), indicating a full match of the word.
Boundary and Match Check: The conditions in searchWord ensure the current cell is valid, not visited, and matches the corresponding character in the word before proceeding.
Backtracking Logic: The backtracking logic is preserved but placed correctly to ensure valid states.
