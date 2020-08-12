package leetcode

func maxCount(m int, n int, ops [][]int) int {
	minM, minN := m, n
	for _, op := range ops {
		minM = min(minM, op[0])
		minN = min(minN, op[1])
	}
	return minM * minN
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
