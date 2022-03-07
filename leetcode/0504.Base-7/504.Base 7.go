package leetcode

import "strconv"

func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}
	negative := false
	if num < 0 {
		negative = true
		num = -num
	}
	var ans string
	var nums []int
	for num != 0 {
		remainder := num % 7
		nums = append(nums, remainder)
		num = num / 7
	}
	if negative {
		ans += "-"
	}
	for i := len(nums) - 1; i >= 0; i-- {
		ans += strconv.Itoa(nums[i])
	}
	return ans
}
