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

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	pre, cur, list2Cur := list1, list1.Next, list2
	for cur.Next != nil {
		if cur.Val == a {
			pre.Next = list2
			pre = cur
			break
		}
		pre = cur
		cur = cur.Next
	}
	cur = cur.Next
	for list2Cur.Next != nil {
		list2Cur = list2Cur.Next
	}
	if a == b {
		list2Cur.Next = cur
		return list1
	}
	for cur.Next != nil {
		if cur.Val == b {
			list2Cur.Next = cur.Next
			break
		}
		pre = cur
		cur = cur.Next
	}
	return list1
}
