package leetcode

import "math"

// 解法一 动态规划
func calculateMinimumHP(dungeon [][]int) int {
	if len(dungeon) == 0 {
		return 0
	}
	m, n := len(dungeon), len(dungeon[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	dp[m-1][n-1] = max(1-dungeon[m-1][n-1], 1)
	for i := n - 2; i >= 0; i-- {
		dp[m-1][i] = max(1, dp[m-1][i+1]-dungeon[m-1][i])
	}
	for i := m - 2; i >= 0; i-- {
		dp[i][n-1] = max(1, dp[i+1][n-1]-dungeon[i][n-1])
	}
	for i := m - 2; i >= 0; i-- {
		for j := n - 2; j >= 0; j-- {
			dp[i][j] = min(max(1, dp[i][j+1]-dungeon[i][j]), max(1, dp[i+1][j]-dungeon[i][j]))
		}
	}
	return dp[0][0]
}

// 解法二 二分搜索
func calculateMinimumHP1(dungeon [][]int) int {
	low, high := 1, math.MaxInt64
	for low < high {
		mid := low + (high-low)>>1
		if canCross(dungeon, mid) {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return low
}

func canCross(dungeon [][]int, start int) bool {
	m, n := len(dungeon), len(dungeon[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	for i := 0; i < len(dp); i++ {
		for j := 0; j < len(dp[i]); j++ {
			if i == 0 && j == 0 {
				dp[i][j] = start + dungeon[0][0]
			} else {
				a, b := math.MinInt64, math.MinInt64
				if i > 0 && dp[i-1][j] > 0 {
					a = dp[i-1][j] + dungeon[i][j]
				}
				if j > 0 && dp[i][j-1] > 0 {
					b = dp[i][j-1] + dungeon[i][j]
				}
				dp[i][j] = max(a, b)
			}
		}
	}
	return dp[m-1][n-1] > 0
}
