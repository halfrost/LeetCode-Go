package leetcode

func sumZero(n int) []int {
	res, left, right, start := make([]int, n), 0, n-1, 1
	for left < right {
		res[left] = start
		res[right] = -start
		start++
		left = left + 1
		right = right - 1
	}
	return res
}
