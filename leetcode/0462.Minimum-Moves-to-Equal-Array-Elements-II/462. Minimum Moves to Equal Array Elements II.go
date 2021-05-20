package leetcode

import (
	"math"
	"sort"
)

func minMoves2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	moves, mid := 0, len(nums)/2
	sort.Ints(nums)
	for i := range nums {
		if i == mid {
			continue
		}
		moves += int(math.Abs(float64(nums[mid] - nums[i])))
	}
	return moves
}
