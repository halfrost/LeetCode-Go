package leetcode

// 解法一 二分搜索
func superEggDrop(K int, N int) int {
	low, high := 1, N
	for low < high {
		mid := low + (high-low)>>1
		if counterF(K, N, mid) >= N {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return low
}

// 计算二项式和，特殊的第一项 C(t,0) = 1
func counterF(k, n, mid int) int {
	res, sum := 1, 0
	for i := 1; i <= k && sum < n; i++ {
		res *= mid - i + 1
		res /= i
		sum += res
	}
	return sum
}

// 解法二 动态规划 DP
func superEggDrop1(K int, N int) int {
	dp, step := make([]int, K+1), 0
	for ; dp[K] < N; step++ {
		for i := K; i > 0; i-- {
			dp[i] = (1 + dp[i] + dp[i-1])
		}
	}
	return step
}
