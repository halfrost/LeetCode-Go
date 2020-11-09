package leetcode

func getMaximumGenerated(n int) int {
	if n == 0 {
		return 0
	}
	nums, max := make([]int, n+1), 0
	nums[0], nums[1] = 0, 1
	for i := 0; i <= n; i++ {
		if nums[i] > max {
			max = nums[i]
		}
		if 2*i >= 2 && 2*i <= n {
			nums[2*i] = nums[i]
		}
		if 2*i+1 >= 2 && 2*i+1 <= n {
			nums[2*i+1] = nums[i] + nums[i+1]
		}
	}
	return max
}
