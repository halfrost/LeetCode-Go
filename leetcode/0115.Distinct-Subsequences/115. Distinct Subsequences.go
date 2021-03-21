package leetcode

// 解法一 压缩版 DP
func numDistinct(s string, t string) int {
	dp := make([]int, len(s)+1)
	for i, curT := range t {
		pre := 0
		for j, curS := range s {
			if i == 0 {
				pre = 1
			}
			newDP := dp[j+1]
			if curT == curS {
				dp[j+1] = dp[j] + pre
			} else {
				dp[j+1] = dp[j]
			}
			pre = newDP
		}
	}
	return dp[len(s)]
}

// 解法二 普通 DP
func numDistinct1(s, t string) int {
	m, n := len(s), len(t)
	if m < n {
		return 0
	}
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		dp[i][n] = 1
	}
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if s[i] == t[j] {
				dp[i][j] = dp[i+1][j+1] + dp[i+1][j]
			} else {
				dp[i][j] = dp[i+1][j]
			}
		}
	}
	return dp[0][0]
}
