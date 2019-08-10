package leetcode

// 解法一 DP
func minCostClimbingStairs(cost []int) int {
	dp := make([]int, len(cost))
	dp[0], dp[1] = cost[0], cost[1]
	for i := 2; i < len(cost); i++ {
		dp[i] = cost[i] + min(dp[i-2], dp[i-1])
	}
	return min(dp[len(cost)-2], dp[len(cost)-1])
}

// 解法二 DP 优化辅助空间
func minCostClimbingStairs1(cost []int) int {
	var cur, last int
	for i := 2; i < len(cost)+1; i++ {
		if last+cost[i-1] > cur+cost[i-2] {
			cur, last = last, cur+cost[i-2]
		} else {
			cur, last = last, last+cost[i-1]
		}
	}
	return last
}
