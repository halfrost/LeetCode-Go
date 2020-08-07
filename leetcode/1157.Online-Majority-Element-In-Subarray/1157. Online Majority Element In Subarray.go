package leetcode

import (
	"sort"
)

type segmentItem struct {
	candidate int
	count     int
}

// MajorityChecker define
type MajorityChecker struct {
	segmentTree []segmentItem
	data        []int
	merge       func(i, j segmentItem) segmentItem
	count       map[int][]int
}

// Constructor1157 define
func Constructor1157(arr []int) MajorityChecker {
	data, tree, mc, count := make([]int, len(arr)), make([]segmentItem, 4*len(arr)), MajorityChecker{}, make(map[int][]int)
	// 这个 merge 函数就是摩尔投票算法
	mc.merge = func(i, j segmentItem) segmentItem {
		if i.candidate == j.candidate {
			return segmentItem{candidate: i.candidate, count: i.count + j.count}
		}
		if i.count > j.count {
			return segmentItem{candidate: i.candidate, count: i.count - j.count}
		}
		return segmentItem{candidate: j.candidate, count: j.count - i.count}
	}
	for i := 0; i < len(arr); i++ {
		data[i] = arr[i]
	}
	for i := 0; i < len(arr); i++ {
		if _, ok := count[arr[i]]; !ok {
			count[arr[i]] = []int{}
		}
		count[arr[i]] = append(count[arr[i]], i)
	}
	mc.data, mc.segmentTree, mc.count = data, tree, count
	if len(arr) > 0 {
		mc.buildSegmentTree(0, 0, len(arr)-1)
	}
	return mc
}

func (mc *MajorityChecker) buildSegmentTree(treeIndex, left, right int) {
	if left == right {
		mc.segmentTree[treeIndex] = segmentItem{candidate: mc.data[left], count: 1}
		return
	}
	leftTreeIndex, rightTreeIndex := mc.leftChild(treeIndex), mc.rightChild(treeIndex)
	midTreeIndex := left + (right-left)>>1
	mc.buildSegmentTree(leftTreeIndex, left, midTreeIndex)
	mc.buildSegmentTree(rightTreeIndex, midTreeIndex+1, right)
	mc.segmentTree[treeIndex] = mc.merge(mc.segmentTree[leftTreeIndex], mc.segmentTree[rightTreeIndex])
}

func (mc *MajorityChecker) leftChild(index int) int {
	return 2*index + 1
}

func (mc *MajorityChecker) rightChild(index int) int {
	return 2*index + 2
}

// Query define
func (mc *MajorityChecker) query(left, right int) segmentItem {
	if len(mc.data) > 0 {
		return mc.queryInTree(0, 0, len(mc.data)-1, left, right)
	}
	return segmentItem{candidate: -1, count: -1}
}

func (mc *MajorityChecker) queryInTree(treeIndex, left, right, queryLeft, queryRight int) segmentItem {
	midTreeIndex, leftTreeIndex, rightTreeIndex := left+(right-left)>>1, mc.leftChild(treeIndex), mc.rightChild(treeIndex)
	if queryLeft <= left && queryRight >= right { // segment completely inside range
		return mc.segmentTree[treeIndex]
	}
	if queryLeft > midTreeIndex {
		return mc.queryInTree(rightTreeIndex, midTreeIndex+1, right, queryLeft, queryRight)
	} else if queryRight <= midTreeIndex {
		return mc.queryInTree(leftTreeIndex, left, midTreeIndex, queryLeft, queryRight)
	}
	// merge query results
	return mc.merge(mc.queryInTree(leftTreeIndex, left, midTreeIndex, queryLeft, midTreeIndex),
		mc.queryInTree(rightTreeIndex, midTreeIndex+1, right, midTreeIndex+1, queryRight))
}

// Query define
func (mc *MajorityChecker) Query(left int, right int, threshold int) int {
	res := mc.query(left, right)
	if _, ok := mc.count[res.candidate]; !ok {
		return -1
	}
	start := sort.Search(len(mc.count[res.candidate]), func(i int) bool { return left <= mc.count[res.candidate][i] })
	end := sort.Search(len(mc.count[res.candidate]), func(i int) bool { return right < mc.count[res.candidate][i] }) - 1
	if (end - start + 1) >= threshold {
		return res.candidate
	}
	return -1
}

/**
 * Your MajorityChecker object will be instantiated and called as such:
 * obj := Constructor(arr);
 * param_1 := obj.Query(left,right,threshold);
 */
