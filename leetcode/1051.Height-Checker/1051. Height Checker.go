package leetcode

import "sort"

func heightChecker(heights []int) int {
	result, checker := 0, []int{}
	checker = append(checker, heights...)
	sort.Ints(checker)
	for i := 0; i < len(heights); i++ {
		if heights[i] != checker[i] {
			result++
		}
	}
	return result
}
