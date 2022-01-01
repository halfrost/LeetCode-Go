package leetcode

func construct2DArray(original []int, m int, n int) [][]int {
	if len(original) != m*n {
		return nil
	}
	ans := make([][]int, m)
	start := 0
	end := n
	for i := 0; i < m; i++ {
		ans[i] = append(ans[i], original[start:end]...)
		start += n
		end += n
	}
	return ans
}
