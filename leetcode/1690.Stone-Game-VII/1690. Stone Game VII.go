package leetcode

// 解法一 优化空间版 DP
func stoneGameVII(stones []int) int {
	n := len(stones)
	sum := make([]int, n)
	dp := make([]int, n)
	for i, d := range stones {
		sum[i] = d
	}
	for i := 1; i < n; i++ {
		for j := 0; j+i < n; j++ {
			if (n-i)%2 == 1 {
				d0 := dp[j] + sum[j]
				d1 := dp[j+1] + sum[j+1]
				if d0 > d1 {
					dp[j] = d0
				} else {
					dp[j] = d1
				}
			} else {
				d0 := dp[j] - sum[j]
				d1 := dp[j+1] - sum[j+1]
				if d0 < d1 {
					dp[j] = d0
				} else {
					dp[j] = d1
				}
			}
			sum[j] = sum[j] + stones[i+j]
		}
	}
	return dp[0]
}

// 解法二 常规 DP
func stoneGameVII1(stones []int) int {
	prefixSum := make([]int, len(stones))
	for i := 0; i < len(stones); i++ {
		if i == 0 {
			prefixSum[i] = stones[i]
		} else {
			prefixSum[i] = prefixSum[i-1] + stones[i]
		}
	}
	dp := make([][]int, len(stones))
	for i := range dp {
		dp[i] = make([]int, len(stones))
		dp[i][i] = 0
	}
	n := len(stones)
	for l := 2; l <= n; l++ {
		for i := 0; i+l <= n; i++ {
			dp[i][i+l-1] = max(prefixSum[i+l-1]-prefixSum[i+1]+stones[i+1]-dp[i+1][i+l-1], prefixSum[i+l-2]-prefixSum[i]+stones[i]-dp[i][i+l-2])
		}
	}
	return dp[0][n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
