package leetcode

import (
	"sort"
	"strconv"
)

func findRelativeRanks(score []int) []string {
	mp := make(map[int]int)
	for i, v := range score {
		mp[v] = i
	}
	sort.Slice(score, func(i, j int) bool {
		return score[i] > score[j]
	})
	ans := make([]string, len(score))
	for i, v := range score {
		if i == 0 {
			ans[mp[v]] = "Gold Medal"
		} else if i == 1 {
			ans[mp[v]] = "Silver Medal"
		} else if i == 2 {
			ans[mp[v]] = "Bronze Medal"
		} else {
			ans[mp[v]] = strconv.Itoa(i + 1)
		}
	}
	return ans
}
