package leetcode

func minSubArrayLen(s int, nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	left, right, res, sum := 0, -1, n+1, 0
	for left < n {
		if (right+1) < n && sum < s {
			right++
			sum += nums[right]
		} else {
			sum -= nums[left]
			left++
		}
		if sum >= s {
			res = min(res, right-left+1)
		}
	}
	if res == n+1 {
		return 0
	}
	return res
}
