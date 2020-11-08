package leetcode

import (
	"container/heap"
)

func maxProfit(inventory []int, orders int) int {
	res, mod := 0, 1000000007
	q := PriorityQueue{}
	for i := 0; i < len(inventory); i++ {
		heap.Push(&q, &Item{count: inventory[i]})
	}
	for ; orders > 0; orders-- {
		item := heap.Pop(&q).(*Item)
		res = (res + item.count) % mod
		heap.Push(&q, &Item{count: item.count - 1})
	}
	return res
}

// Item define
type Item struct {
	count int
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	// 注意：因为golang中的heap是按最小堆组织的，所以count越大，Less()越小，越靠近堆顶.
	return pq[i].count > pq[j].count
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

// Push define
func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	*pq = append(*pq, item)
}

// Pop define
func (pq *PriorityQueue) Pop() interface{} {
	n := len(*pq)
	item := (*pq)[n-1]
	*pq = (*pq)[:n-1]
	return item
}
