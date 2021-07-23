package leetcode

import "sort"

func minPairSum(nums []int) int {
	sort.Ints(nums)
	n, res := len(nums), 0
	for i, val := range nums[:n/2] {
		res = max(res, val+nums[n-1-i])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
