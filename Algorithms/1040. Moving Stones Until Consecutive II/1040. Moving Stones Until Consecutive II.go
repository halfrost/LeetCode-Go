package leetcode

import (
	"math"
	"sort"
)

func numMovesStonesII(stones []int) []int {
	if len(stones) == 0 {
		return []int{0, 0}
	}
	sort.Ints(stones)
	n := len(stones)
	maxStep, minStep, left, right := max(stones[n-1]-stones[1]-n+2, stones[n-2]-stones[0]-n+2), math.MaxInt64, 0, 0
	for left < n {
		if right+1 < n && stones[right]-stones[left] < n {
			right++
		} else {
			if stones[right]-stones[left] >= n {
				right--
			}
			if right-left+1 == n-1 && stones[right]-stones[left]+1 == n-1 {
				minStep = min(minStep, 2)
			} else {
				minStep = min(minStep, n-(right-left+1))
			}
			if right == n-1 && stones[right]-stones[left] < n {
				break
			}
			left++
		}
	}
	return []int{minStep, maxStep}
}
