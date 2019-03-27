package leetcode

import (
	"strconv"
)

func largestNumber(nums []int) string {
	if len(nums) == 0 {
		return ""
	}
	numStrs := toStringArray(nums)
	quickSortString(numStrs, 0, len(numStrs)-1)
	res := ""
	for _, str := range numStrs {
		if res == "0" && str == "0" {
			continue
		}
		res = res + str
	}
	return res
}

func toStringArray(nums []int) []string {
	strs := make([]string, 0)
	for _, num := range nums {
		strs = append(strs, strconv.Itoa(num))
	}
	return strs
}
func partitionString(a []string, lo, hi int) int {
	pivot := a[hi]
	i := lo - 1
	for j := lo; j < hi; j++ {
		ajStr := a[j] + pivot
		pivotStr := pivot + a[j]
		if ajStr > pivotStr { // 这里的判断条件是关键
			i++
			a[j], a[i] = a[i], a[j]
		}
	}
	a[i+1], a[hi] = a[hi], a[i+1]
	return i + 1
}
func quickSortString(a []string, lo, hi int) {
	if lo >= hi {
		return
	}
	p := partitionString(a, lo, hi)
	quickSortString(a, lo, p-1)
	quickSortString(a, p+1, hi)
}
