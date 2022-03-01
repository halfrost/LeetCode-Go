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
func mergeNodes(head *ListNode) *ListNode {
	res := &ListNode{}
	h := res
	if head.Next == nil {
		return &structures.ListNode{}
	}
	cur := head
	sum := 0
	for cur.Next != nil {
		if cur.Next.Val != 0 {
			sum += cur.Next.Val
		} else {
			h.Next = &ListNode{Val: sum, Next: nil}
			h = h.Next
			sum = 0
		}
		cur = cur.Next
	}
	return res.Next
}
