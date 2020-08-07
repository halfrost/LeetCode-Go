package leetcode

func canJump(nums []int) bool {
	n := len(nums)
	if n == 0 {
		return false
	}
	if n == 1 {
		return true
	}
	maxJump := 0
	for i, v := range nums {
		if i > maxJump {
			return false
		}
		maxJump = max(maxJump, i+v)
	}
	return true
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
