package leetcode

import (
	"container/heap"
)

// 解法一 二分搜索
func kthSmallest378(matrix [][]int, k int) int {
	m, n, low := len(matrix), len(matrix[0]), matrix[0][0]
	high := matrix[m-1][n-1] + 1
	for low < high {
		mid := low + (high-low)>>1
		// 如果 count 比 k 小，在大值的那一半继续二分搜索
		if counterKthSmall(m, n, mid, matrix) >= k {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return low
}

func counterKthSmall(m, n, mid int, matrix [][]int) int {
	count, j := 0, n-1
	// 每次循环统计比 mid 值小的元素个数
	for i := 0; i < m; i++ {
		// 遍历每行中比 mid 小的元素的个数
		for j >= 0 && mid < matrix[i][j] {
			j--
		}
		count += j + 1
	}
	return count
}

// 解法二 优先队列
func kthSmallest3781(matrix [][]int, k int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	pq := &pq{data: make([]interface{}, k)}
	heap.Init(pq)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if pq.Len() < k {
				heap.Push(pq, matrix[i][j])
			} else if matrix[i][j] < pq.Head().(int) {
				heap.Pop(pq)
				heap.Push(pq, matrix[i][j])
			} else {
				break
			}
		}
	}
	return heap.Pop(pq).(int)
}

type pq struct {
	data []interface{}
	len  int
}

func (p *pq) Len() int {
	return p.len
}

func (p *pq) Less(a, b int) bool {
	return p.data[a].(int) > p.data[b].(int)
}

func (p *pq) Swap(a, b int) {
	p.data[a], p.data[b] = p.data[b], p.data[a]
}

func (p *pq) Push(o interface{}) {
	p.data[p.len] = o
	p.len++
}

func (p *pq) Head() interface{} {
	return p.data[0]
}

func (p *pq) Pop() interface{} {
	p.len--
	return p.data[p.len]
}
