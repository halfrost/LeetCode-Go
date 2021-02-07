package leetcode

func findMaxAverage(nums []int, k int) float64 {
	sum := 0
	for _, v := range nums[:k] {
		sum += v
	}
	maxSum := sum
	for i := k; i < len(nums); i++ {
		sum = sum - nums[i-k] + nums[i]
		maxSum = max(maxSum, sum)
	}
	return float64(maxSum) / float64(k)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
