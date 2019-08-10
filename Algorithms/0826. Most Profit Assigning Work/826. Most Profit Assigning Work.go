package leetcode

import (
	"fmt"
	"sort"
)

// Task define
type Task struct {
	Difficulty int
	Profit     int
}

// Tasks define
type Tasks []Task

// Len define
func (p Tasks) Len() int { return len(p) }

// Swap define
func (p Tasks) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

// SortByDiff define
type SortByDiff struct{ Tasks }

// Less define
func (p SortByDiff) Less(i, j int) bool {
	return p.Tasks[i].Difficulty < p.Tasks[j].Difficulty
}

func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {
	if len(difficulty) == 0 || len(profit) == 0 || len(worker) == 0 {
		return 0
	}
	tasks, res, index := []Task{}, 0, 0
	for i := 0; i < len(difficulty); i++ {
		tasks = append(tasks, Task{Difficulty: difficulty[i], Profit: profit[i]})
	}
	sort.Sort(SortByDiff{tasks})
	sort.Ints(worker)
	for i := 1; i < len(tasks); i++ {
		tasks[i].Profit = max(tasks[i].Profit, tasks[i-1].Profit)
	}
	fmt.Printf("tasks = %v worker = %v\n", tasks, worker)
	for _, w := range worker {
		for index < len(difficulty) && w >= tasks[index].Difficulty {
			index++
		}
		fmt.Printf("tasks【index】 = %v\n", tasks[index])
		if index > 0 {
			res += tasks[index-1].Profit
		}
	}
	return res
}
