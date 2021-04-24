package leetcode

func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 1; i <= target; i++ {
		for _, num := range nums {
			if i-num >= 0 {
				dp[i] += dp[i-num]
			}
		}
	}
	return dp[target]
}

// 暴力解法超时
func combinationSum41(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	c, res := []int{}, 0
	findcombinationSum4(nums, target, 0, c, &res)
	return res
}

func findcombinationSum4(nums []int, target, index int, c []int, res *int) {
	if target <= 0 {
		if target == 0 {
			*res++
		}
		return
	}
	for i := 0; i < len(nums); i++ {
		c = append(c, nums[i])
		findcombinationSum4(nums, target-nums[i], i, c, res)
		c = c[:len(c)-1]
	}
}
