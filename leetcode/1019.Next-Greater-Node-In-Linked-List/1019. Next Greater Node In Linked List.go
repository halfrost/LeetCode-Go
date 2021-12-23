package leetcode

import (
	"github.com/halfrost/LeetCode-Go/structures"
)

// ListNode define
type ListNode = structures.ListNode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 解法一 单调栈
func nextLargerNodes(head *ListNode) []int {
	type node struct {
		index, val int
	}
	var monoStack []node
	var res []int
	for head != nil {
		for len(monoStack) > 0 && monoStack[len(monoStack)-1].val < head.Val {
			res[monoStack[len(monoStack)-1].index] = head.Val
			monoStack = monoStack[:len(monoStack)-1]
		}
		monoStack = append(monoStack, node{len(res), head.Val})
		res = append(res, 0)
		head = head.Next
	}
	return res
}
