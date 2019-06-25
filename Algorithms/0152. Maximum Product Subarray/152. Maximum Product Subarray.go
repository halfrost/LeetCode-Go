package leetcode

func maxProduct(nums []int) int {
	minimum, maximum, res := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			maximum, minimum = minimum, maximum
		}
		maximum = max(nums[i], maximum*nums[i])
		minimum = min(nums[i], minimum*nums[i])
		res = max(res, maximum)
	}
	return res
}
