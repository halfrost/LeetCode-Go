package leetcode

func minOperations(nums []int, x int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	target := total - x
	if target < 0 {
		return -1
	}
	if target == 0 {
		return len(nums)
	}
	left, right, sum, res := 0, 0, 0, -1
	for right < len(nums) {
		if sum < target {
			sum += nums[right]
			right++
		}
		for sum >= target {
			if sum == target {
				res = max(res, right-left)
			}
			sum -= nums[left]
			left++
		}
	}
	if res == -1 {
		return -1
	}
	return len(nums) - res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
