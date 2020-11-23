package leetcode

import (
	"sort"
)

func minimumEffort(tasks [][]int) int {
	sort.Sort(Task(tasks))
	res, cur := 0, 0
	for _, t := range tasks {
		if t[1] > cur {
			res += t[1] - cur
			cur = t[1] - t[0]
		} else {
			cur -= t[0]
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Task define
type Task [][]int

func (task Task) Len() int {
	return len(task)
}

func (task Task) Less(i, j int) bool {
	t1, t2 := task[i][1]-task[i][0], task[j][1]-task[j][0]
	if t1 != t2 {
		return t2 < t1
	}
	return task[j][1] < task[i][1]
}

func (task Task) Swap(i, j int) {
	t := task[i]
	task[i] = task[j]
	task[j] = t
}
