package leetcode

func transpose(A [][]int) [][]int {
	row, col, result := len(A), len(A[0]), make([][]int, len(A[0]))
	for i := range result {
		result[i] = make([]int, row)
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			result[j][i] = A[i][j]
		}
	}
	return result
}
