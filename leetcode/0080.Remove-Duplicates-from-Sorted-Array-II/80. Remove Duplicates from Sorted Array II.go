package leetcode

func removeDuplicates(nums []int) int {
	slow := 0
	for i, v := range nums {
		if i < 2 || nums[slow-2] != v {
			nums[slow] = v
			slow++
		}
	}
	return slow
}
