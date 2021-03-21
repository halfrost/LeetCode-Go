package leetcode

import "sort"

func smallestRangeII(A []int, K int) int {
	n := len(A)
	sort.Ints(A)
	res := A[n-1] - A[0]
	for i := 0; i < n-1; i++ {
		a, b := A[i], A[i+1]
		high := max(A[n-1]-K, a+K)
		low := min(A[0]+K, b-K)
		res = min(res, high-low)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
