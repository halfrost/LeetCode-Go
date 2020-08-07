package leetcode

import (
	"sort"
)

func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)
	left, right, res := 0, len(people)-1, 0
	for left <= right {
		if left == right {
			res++
			return res
		}
		if people[left]+people[right] <= limit {
			left++
			right--
		} else {
			right--
		}
		res++
	}
	return res
}
