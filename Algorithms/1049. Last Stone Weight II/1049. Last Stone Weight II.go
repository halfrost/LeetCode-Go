package leetcode

func lastStoneWeightII(stones []int) int {
	sum := 0
	for _, v := range stones {
		sum += v
	}
	n, C, dp := len(stones), sum/2, make([]int, sum/2+1)
	for i := 0; i <= C; i++ {
		if stones[0] <= i {
			dp[i] = stones[0]
		} else {
			dp[i] = 0
		}
	}
	for i := 1; i < n; i++ {
		for j := C; j >= stones[i]; j-- {
			dp[j] = max(dp[j], dp[j-stones[i]]+stones[i])
		}
	}
	return sum - 2*dp[C]
}
