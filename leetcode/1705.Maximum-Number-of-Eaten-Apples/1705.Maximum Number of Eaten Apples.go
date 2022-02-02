package leetcode

import "container/heap"

func eatenApples(apples []int, days []int) int {
	h := hp{}
	i := 0
	var ans int
	for ; i < len(apples); i++ {
		for len(h) > 0 && h[0].end <= i {
			heap.Pop(&h)
		}
		if apples[i] > 0 {
			heap.Push(&h, data{apples[i], i + days[i]})
		}
		if len(h) > 0 {
			minData := heap.Pop(&h).(data)
			ans++
			if minData.left > 1 {
				heap.Push(&h, data{minData.left - 1, minData.end})
			}
		}
	}
	for len(h) > 0 {
		for len(h) > 0 && h[0].end <= i {
			heap.Pop(&h)
		}
		if len(h) == 0 {
			break
		}
		minData := heap.Pop(&h).(data)
		nums := min(minData.left, minData.end-i)
		ans += nums
		i += nums
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type data struct {
	left int
	end  int
}

type hp []data

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].end < h[j].end }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *hp) Push(x interface{}) {
	*h = append(*h, x.(data))
}

func (h *hp) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
