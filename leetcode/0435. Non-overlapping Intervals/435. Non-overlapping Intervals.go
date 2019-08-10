package leetcode

import (
	"sort"
)

// 解法一 DP O(n^2) 思路是仿造最长上升子序列的思路
func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}
	sort.Sort(Intervals(intervals))
	dp, res := make([]int, len(intervals)), 0
	for i := range dp {
		dp[i] = 1
	}
	for i := 1; i < len(intervals); i++ {
		for j := 0; j < i; j++ {
			if intervals[i][0] >= intervals[j][1] {
				dp[i] = max(dp[i], 1+dp[j])
			}
		}
	}
	for _, v := range dp {
		res = max(res, v)
	}
	return len(intervals) - res
}

// Intervals define
type Intervals [][]int

func (a Intervals) Len() int {
	return len(a)
}
func (a Intervals) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a Intervals) Less(i, j int) bool {
	for k := 0; k < len(a[i]); k++ {
		if a[i][k] < a[j][k] {
			return true
		} else if a[i][k] == a[j][k] {
			continue
		} else {
			return false
		}
	}
	return true
}

// 解法二 贪心 O(n)
func eraseOverlapIntervals1(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}
	sort.Sort(Intervals(intervals))
	pre, res := 0, 1

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] >= intervals[pre][1] {
			res++
			pre = i
		} else if intervals[i][1] < intervals[pre][1] {
			pre = i
		}
	}
	return len(intervals) - res
}
