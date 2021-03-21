package leetcode

func kLengthApart(nums []int, k int) bool {
	prevIndex := -1
	for i, num := range nums {
		if num == 1 {
			if prevIndex != -1 && i-prevIndex-1 < k {
				return false
			}
			prevIndex = i
		}
	}
	return true
}
