package leetcode

// 单调栈
func mostCompetitive(nums []int, k int) []int {
	stack := make([]int, 0, len(nums))
	for i := 0; i < len(nums); i++ {
		for len(stack)+len(nums)-i > k && len(stack) > 0 && nums[i] < stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, nums[i])
	}
	return stack[:k]
}
