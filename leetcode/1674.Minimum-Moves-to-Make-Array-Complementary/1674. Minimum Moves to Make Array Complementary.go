package leetcode

func minMoves(nums []int, limit int) int {
	diff := make([]int, limit*2+2) // nums[i] <= limit, b+limit+1 is maximum limit+limit+1
	for j := 0; j < len(nums)/2; j++ {
		a, b := min(nums[j], nums[len(nums)-j-1]), max(nums[j], nums[len(nums)-j-1])
		// using prefix sum: most interesting point, and is the key to reduce complexity
		diff[2] += 2
		diff[a+1]--
		diff[a+b]--
		diff[a+b+1]++
		diff[b+limit+1]++
	}
	cur, res := 0, len(nums)
	for i := 2; i <= 2*limit; i++ {
		cur += diff[i]
		res = min(res, cur)
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
