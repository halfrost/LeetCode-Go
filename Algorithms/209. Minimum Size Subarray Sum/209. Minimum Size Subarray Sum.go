package leetcode

func minSubArrayLen(s int, nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	l, r := 0, -1
	res := n + 1
	sum := 0
	for l < n {
		if (r+1) < n && sum < s {
			r++
			sum += nums[r]
		} else {
			sum -= nums[l]
			l++
		}
		if sum >= s {
			res = min(res, r-l+1)
		}
	}
	if res == n+1 {
		return 0
	}
	return res
}
