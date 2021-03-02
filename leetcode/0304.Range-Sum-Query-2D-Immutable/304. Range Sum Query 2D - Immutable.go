package leetcode

type NumMatrix struct {
	cumsum [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	if len(matrix) == 0 {
		return NumMatrix{nil}
	}
	cumsum := make([][]int, len(matrix)+1)
	cumsum[0] = make([]int, len(matrix[0])+1)
	for i := range matrix {
		cumsum[i+1] = make([]int, len(matrix[i])+1)
		for j := range matrix[i] {
			cumsum[i+1][j+1] = matrix[i][j] + cumsum[i][j+1] + cumsum[i+1][j] - cumsum[i][j]
		}
	}
	return NumMatrix{cumsum}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	cumsum := this.cumsum
	return cumsum[row2+1][col2+1] - cumsum[row1][col2+1] - cumsum[row2+1][col1] + cumsum[row1][col1]
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
