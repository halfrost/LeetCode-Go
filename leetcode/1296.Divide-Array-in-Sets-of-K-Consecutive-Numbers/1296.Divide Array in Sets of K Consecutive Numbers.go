package leetcode

import "sort"

func isPossibleDivide(nums []int, k int) bool {
	mp := make(map[int]int)
	for _, v := range nums {
		mp[v] += 1
	}
	sort.Ints(nums)
	for _, num := range nums {
		if mp[num] == 0 {
			continue
		}
		for diff := 0; diff < k; diff++ {
			if mp[num+diff] == 0 {
				return false
			}
			mp[num+diff] -= 1
		}
	}
	return true
}
