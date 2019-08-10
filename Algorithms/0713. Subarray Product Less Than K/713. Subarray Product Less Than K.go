package leetcode

func numSubarrayProductLessThanK(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}
	res, left, right, prod := 0, 0, 0, 1
	for left < len(nums) {
		if right < len(nums) && prod*nums[right] < k {
			prod = prod * nums[right]
			right++
		} else if left == right {
			left++
			right++
		} else {
			res += right - left
			prod = prod / nums[left]
			left++
		}
	}
	return res
}
