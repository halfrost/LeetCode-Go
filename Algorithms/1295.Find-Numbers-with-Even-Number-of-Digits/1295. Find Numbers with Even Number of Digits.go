package leetcode

import "strconv"

func findNumbers(nums []int) int {
	res := 0
	for _, n := range nums {
		res += 1 - len(strconv.Itoa(n))%2
	}
	return res
}
