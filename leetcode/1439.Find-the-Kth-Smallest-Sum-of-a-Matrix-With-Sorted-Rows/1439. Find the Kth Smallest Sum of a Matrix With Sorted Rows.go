package leetcode

import "container/heap"

func kthSmallest(mat [][]int, k int) int {
	if len(mat) == 0 || len(mat[0]) == 0 || k == 0 {
		return 0
	}
	prev := mat[0]
	for i := 1; i < len(mat); i++ {
		prev = kSmallestPairs(prev, mat[i], k)
	}
	if k < len(prev) {
		return -1
	}
	return prev[k-1]
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) []int {
	res := []int{}
	if len(nums2) == 0 {
		return res
	}
	pq := newPriorityQueue()
	for i := 0; i < len(nums1) && i < k; i++ {
		heap.Push(pq, &pddata{
			n1:    nums1[i],
			n2:    nums2[0],
			n2Idx: 0,
		})
	}
	for pq.Len() > 0 {
		i := heap.Pop(pq)
		data := i.(*pddata)
		res = append(res, data.n1+data.n2)
		k--
		if k <= 0 {
			break
		}
		idx := data.n2Idx
		idx++
		if idx >= len(nums2) {
			continue
		}
		heap.Push(pq, &pddata{
			n1:    data.n1,
			n2:    nums2[idx],
			n2Idx: idx,
		})
	}
	return res
}

type pddata struct {
	n1    int
	n2    int
	n2Idx int
}

type priorityQueue []*pddata

func newPriorityQueue() *priorityQueue {
	pq := priorityQueue([]*pddata{})
	heap.Init(&pq)
	return &pq
}

func (pq priorityQueue) Len() int           { return len(pq) }
func (pq priorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }
func (pq priorityQueue) Less(i, j int) bool { return pq[i].n1+pq[i].n2 < pq[j].n1+pq[j].n2 }
func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	val := old[len(old)-1]
	old[len(old)-1] = nil
	*pq = old[0 : len(old)-1]
	return val
}

func (pq *priorityQueue) Push(i interface{}) {
	val := i.(*pddata)
	*pq = append(*pq, val)
}
