package leetcode

import (
	"sort"
)

func maxEvents(events [][]int) int {
	sort.Slice(events, func(i, j int) bool {
		if events[i][0] == events[j][0] {
			return events[i][1] < events[j][1]
		}
		return events[i][0] < events[j][0]
	})
	attended, current := 1, events[0]
	for i := 1; i < len(events); i++ {
		prev, event := events[i-1], events[i]
		if event[0] == prev[0] && event[1] == prev[1] && event[1] == event[0] {
			continue
		}
		start, end := max(current[0], event[0]-1), max(current[1], event[1])
		if end-start > 0 {
			current[0] = start + 1
			current[1] = end
			attended++
		}
	}
	return attended
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
