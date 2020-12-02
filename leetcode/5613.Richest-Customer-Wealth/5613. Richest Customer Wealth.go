package leetcode

func maximumWealth(accounts [][]int) int {
	res := 0
	for i := 0; i < len(accounts); i++ {
		sum := 0
		for j := 0; j < len(accounts[i]); j++ {
			sum += accounts[i][j]
		}
		res = max(res, sum)
	}
	return res
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
