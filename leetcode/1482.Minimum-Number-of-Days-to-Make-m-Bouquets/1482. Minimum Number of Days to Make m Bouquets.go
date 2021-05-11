package leetcode

import "sort"

func minDays(bloomDay []int, m int, k int) int {
	if m*k > len(bloomDay) {
		return -1
	}
	maxDay := 0
	for _, day := range bloomDay {
		if day > maxDay {
			maxDay = day
		}
	}
	return sort.Search(maxDay, func(days int) bool {
		flowers, bouquets := 0, 0
		for _, d := range bloomDay {
			if d > days {
				flowers = 0
			} else {
				flowers++
				if flowers == k {
					bouquets++
					flowers = 0
				}
			}
		}
		return bouquets >= m
	})
}
