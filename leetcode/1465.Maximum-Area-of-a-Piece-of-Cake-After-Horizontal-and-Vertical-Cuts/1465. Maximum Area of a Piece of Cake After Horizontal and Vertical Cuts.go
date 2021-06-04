package leetcode

import "sort"

func maxArea(h int, w int, horizontalCuts []int, verticalCuts []int) int {
	sort.Ints(horizontalCuts)
	sort.Ints(verticalCuts)
	maxw, maxl := horizontalCuts[0], verticalCuts[0]
	for i, c := range horizontalCuts[1:] {
		if c-horizontalCuts[i] > maxw {
			maxw = c - horizontalCuts[i]
		}
	}
	if h-horizontalCuts[len(horizontalCuts)-1] > maxw {
		maxw = h - horizontalCuts[len(horizontalCuts)-1]
	}
	for i, c := range verticalCuts[1:] {
		if c-verticalCuts[i] > maxl {
			maxl = c - verticalCuts[i]
		}
	}
	if w-verticalCuts[len(verticalCuts)-1] > maxl {
		maxl = w - verticalCuts[len(verticalCuts)-1]
	}
	return (maxw * maxl) % (1000000007)
}
