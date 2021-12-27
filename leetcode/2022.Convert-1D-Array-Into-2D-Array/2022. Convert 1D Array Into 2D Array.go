package leetcode

func construct2DArray(original []int, m int, n int) [][]int {
	if m*n != len(original) {
		return [][]int{}
	}
	res := make([][]int, m)
	for i := 0; i < m; i++ {
		res[i] = original[n*i : n*(i+1)]
	}
	return res
}
