package leetcode

import "sort"

func minimumDifference(nums []int, k int) int {
	sort.Ints(nums)
	minDiff := 100000 + 1
	for i := 0; i < len(nums); i++ {
		if i+k-1 >= len(nums) {
			break
		}
		diff := nums[i+k-1] - nums[i]
		if diff < minDiff {
			minDiff = diff
		}
	}
	return minDiff
}
