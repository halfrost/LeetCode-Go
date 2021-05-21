package leetcode

func diagonalSum(mat [][]int) int {
	n := len(mat)
	ans := 0
	for pi := 0; pi < n; pi++ {
		ans += mat[pi][pi]
	}
	for si, sj := n-1, 0; sj < n; si, sj = si-1, sj+1 {
		ans += mat[si][sj]
	}
	if n%2 == 0 {
		return ans
	}
	return ans - mat[n/2][n/2]
}
