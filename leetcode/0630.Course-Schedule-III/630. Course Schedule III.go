package leetcode

import (
	"container/heap"
	"sort"
)

func scheduleCourse(courses [][]int) int {
	sort.Slice(courses, func(i, j int) bool {
		return courses[i][1] < courses[j][1]
	})
	maxHeap, time := &Schedule{}, 0
	heap.Init(maxHeap)
	for _, c := range courses {
		if time+c[0] <= c[1] {
			time += c[0]
			heap.Push(maxHeap, c[0])
		} else if (*maxHeap).Len() > 0 && (*maxHeap)[0] > c[0] {
			time -= heap.Pop(maxHeap).(int) - c[0]
			heap.Push(maxHeap, c[0])
		}
	}
	return (*maxHeap).Len()
}

type Schedule []int

func (s Schedule) Len() int           { return len(s) }
func (s Schedule) Less(i, j int) bool { return s[i] > s[j] }
func (s Schedule) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s *Schedule) Pop() interface{} {
	n := len(*s)
	t := (*s)[n-1]
	*s = (*s)[:n-1]
	return t
}
func (s *Schedule) Push(x interface{}) {
	*s = append(*s, x.(int))
}
