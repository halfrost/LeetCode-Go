package leetcode

func findMaxConsecutiveOnes(nums []int) int {
	maxCount, currentCount := 0, 0
	for _, v := range nums {
		if v == 1 {
			currentCount++
		} else {
			currentCount = 0
		}
		if currentCount > maxCount {
			maxCount = currentCount
		}
	}
	return maxCount
}
