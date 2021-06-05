package leetcode

import (
	"container/heap"
	"sort"
)

func maxPerformance(n int, speed []int, efficiency []int, k int) int {
	indexes := make([]int, n)
	for i := range indexes {
		indexes[i] = i
	}
	sort.Slice(indexes, func(i, j int) bool {
		return efficiency[indexes[i]] > efficiency[indexes[j]]
	})
	ph := speedHeap{}
	heap.Init(&ph)
	speedSum := 0
	var max int64
	for _, index := range indexes {
		if ph.Len() == k {
			speedSum -= heap.Pop(&ph).(int)
		}
		speedSum += speed[index]
		heap.Push(&ph, speed[index])
		max = Max(max, int64(speedSum)*int64(efficiency[index]))
	}
	return int(max % (1e9 + 7))
}

type speedHeap []int

func (h speedHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h speedHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h speedHeap) Len() int            { return len(h) }
func (h *speedHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *speedHeap) Pop() interface{} {
	res := (*h)[len(*h)-1]
	*h = (*h)[:h.Len()-1]
	return res
}

func Max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
