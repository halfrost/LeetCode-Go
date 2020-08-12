package leetcode

import "math"

func minMoves(nums []int) int {
	sum, min, l := 0, math.MaxInt32, len(nums)
	for _, v := range nums {
		sum += v
		if min > v {
			min = v
		}
	}
	return sum - min*l
}
