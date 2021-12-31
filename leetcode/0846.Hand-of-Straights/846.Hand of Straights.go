package leetcode

import "sort"

func isNStraightHand(hand []int, groupSize int) bool {
	mp := make(map[int]int)
	for _, v := range hand {
		mp[v] += 1
	}
	sort.Ints(hand)
	for _, num := range hand {
		if mp[num] == 0 {
			continue
		}
		for diff := 0; diff < groupSize; diff++ {
			if mp[num+diff] == 0 {
				return false
			}
			mp[num+diff] -= 1
		}
	}
	return true
}
