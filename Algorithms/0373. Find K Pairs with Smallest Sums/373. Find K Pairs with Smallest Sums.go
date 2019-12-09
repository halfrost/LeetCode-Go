package leetcode

import (
	"container/heap"
	"sort"
)

// 解法一 优先队列
func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	result, h := [][]int{}, &minHeap{}
	if len(nums1) == 0 || len(nums2) == 0 || k == 0 {
		return result
	}
	if len(nums1)*len(nums2) < k {
		k = len(nums1) * len(nums2)
	}
	heap.Init(h)
	for _, num := range nums1 {
		heap.Push(h, []int{num, nums2[0], 0})
	}
	for len(result) < k {
		min := heap.Pop(h).([]int)
		result = append(result, min[:2])
		if min[2] < len(nums2)-1 {
			heap.Push(h, []int{min[0], nums2[min[2]+1], min[2] + 1})
		}
	}
	return result
}

type minHeap [][]int

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i][0]+h[i][1] < h[j][0]+h[j][1] }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.([]int))
}

func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// 解法二 暴力解法
func kSmallestPairs1(nums1 []int, nums2 []int, k int) [][]int {
	size1, size2, res := len(nums1), len(nums2), [][]int{}
	if size1 == 0 || size2 == 0 || k < 0 {
		return nil
	}
	for i := 0; i < size1; i++ {
		for j := 0; j < size2; j++ {
			res = append(res, []int{nums1[i], nums2[j]})
		}
	}
	sort.Slice(res, func(i, j int) bool {
		return res[i][0]+res[i][1] < res[j][0]+res[j][1]
	})
	if len(res) >= k {
		return res[:k]
	}
	return res
}
