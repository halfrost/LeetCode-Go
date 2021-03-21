package leetcode

func maximumUniqueSubarray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	result, left, right, freq := 0, 0, -1, map[int]int{}
	for left < len(nums) {
		if right+1 < len(nums) && freq[nums[right+1]] == 0 {
			freq[nums[right+1]]++
			right++
		} else {
			freq[nums[left]]--
			left++
		}
		sum := 0
		for i := left; i <= right; i++ {
			sum += nums[i]
		}
		result = max(result, sum)
	}
	return result
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
