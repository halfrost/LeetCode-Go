package leetcode

import (
	"container/heap"
)

// 解法一 贪心 + 二分搜索
func maxProfit(inventory []int, orders int) int {
	maxItem, thresholdValue, count, res, mod := 0, -1, 0, 0, 1000000007
	for i := 0; i < len(inventory); i++ {
		if inventory[i] > maxItem {
			maxItem = inventory[i]
		}
	}
	low, high := 0, maxItem
	for low <= high {
		mid := low + ((high - low) >> 1)
		for i := 0; i < len(inventory); i++ {
			count += max(inventory[i]-mid, 0)
		}
		if count <= orders {
			thresholdValue = mid
			high = mid - 1
		} else {
			low = mid + 1
		}
		count = 0
	}
	count = 0
	for i := 0; i < len(inventory); i++ {
		count += max(inventory[i]-thresholdValue, 0)
	}
	count = orders - count
	for i := 0; i < len(inventory); i++ {
		if inventory[i] >= thresholdValue {
			if count > 0 {
				res += (thresholdValue + inventory[i]) * (inventory[i] - thresholdValue + 1) / 2
				count--
			} else {
				res += (thresholdValue + 1 + inventory[i]) * (inventory[i] - thresholdValue) / 2
			}
		}
	}
	return res % mod
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

// 解法二 优先队列，超时！
func maxProfit_(inventory []int, orders int) int {
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
