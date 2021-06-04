package leetcode

func findMaxLength(nums []int) int {
	dict := map[int]int{}
	dict[0] = -1
	count, res := 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			count--
		} else {
			count++
		}
		if idx, ok := dict[count]; ok {
			res = max(res, i-idx)
		} else {
			dict[count] = i
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
