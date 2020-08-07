package leetcode

import "sort"

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	gi, si, res := 0, 0, 0
	for gi < len(g) && si < len(s) {
		if s[si] >= g[gi] {
			res++
			si++
			gi++
		} else {
			si++
		}
	}
	return res
}
