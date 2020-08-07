package leetcode

import (
	"math"
	"sort"
)

func minimumAbsDifference(arr []int) [][]int {
	minDiff, res := math.MaxInt32, [][]int{}
	sort.Ints(arr)
	for i := 1; i < len(arr); i++ {
		if arr[i]-arr[i-1] < minDiff {
			minDiff = arr[i] - arr[i-1]
		}
		if minDiff == 1 {
			break
		}
	}
	for i := 1; i < len(arr); i++ {
		if arr[i]-arr[i-1] == minDiff {
			res = append(res, []int{arr[i-1], arr[i]})
		}
	}
	return res
}
