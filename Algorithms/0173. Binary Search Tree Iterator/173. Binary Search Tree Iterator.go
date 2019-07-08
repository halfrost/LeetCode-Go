package leetcode

import "container/heap"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// BSTIterator define
type BSTIterator struct {
	pq    PriorityQueueOfInt
	count int
}

// Constructor173 define
func Constructor173(root *TreeNode) BSTIterator {
	result, pq := []int{}, PriorityQueueOfInt{}
	postorder(root, &result)
	for _, v := range result {
		heap.Push(&pq, v)
	}
	bs := BSTIterator{pq: pq, count: len(result)}
	return bs
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	this.count--
	return heap.Pop(&this.pq).(int)
}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	return this.count != 0
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
type PriorityQueueOfInt []int

func (pq PriorityQueueOfInt) Len() int {
	return len(pq)
}

func (pq PriorityQueueOfInt) Less(i, j int) bool {
	return pq[i] < pq[j]
}

func (pq PriorityQueueOfInt) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueueOfInt) Push(x interface{}) {
	item := x.(int)
	*pq = append(*pq, item)
}

func (pq *PriorityQueueOfInt) Pop() interface{} {
	n := len(*pq)
	item := (*pq)[n-1]
	*pq = (*pq)[:n-1]
	return item
}
