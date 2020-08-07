package leetcode

func isToeplitzMatrix(matrix [][]int) bool {
	rows, columns := len(matrix), len(matrix[0])
	for i := 1; i < rows; i++ {
		for j := 1; j < columns; j++ {
			if matrix[i-1][j-1] != matrix[i][j] {
				return false
			}
		}
	}
	return true
}
