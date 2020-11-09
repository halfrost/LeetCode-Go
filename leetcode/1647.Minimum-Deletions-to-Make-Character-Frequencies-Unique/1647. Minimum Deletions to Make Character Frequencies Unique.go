package leetcode

import (
	"sort"
)

func minDeletions(s string) int {
	frequency, res := make([]int, 26), 0
	for i := 0; i < len(s); i++ {
		frequency[s[i]-'a']++
	}
	sort.Sort(sort.Reverse(sort.IntSlice(frequency)))
	for i := 1; i <= 25; i++ {
		if frequency[i] == frequency[i-1] && frequency[i] != 0 {
			res++
			frequency[i]--
			sort.Sort(sort.Reverse(sort.IntSlice(frequency)))
			i--
		}
	}
	return res
}
