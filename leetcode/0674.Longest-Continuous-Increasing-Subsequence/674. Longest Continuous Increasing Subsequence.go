package leetcode

func findLengthOfLCIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	res, length := 1, 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			length++
		} else {
			res = max(res, length)
			length = 1
		}
	}
	return max(res, length)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
