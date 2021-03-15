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
func swapNodes(head *ListNode, k int) *ListNode {
	count := 1
	var a, b *ListNode
	for node := head; node != nil; node = node.Next {
		if count == k {
			a = node
		}
		count++
	}
	length := count
	count = 1
	for node := head; node != nil; node = node.Next {
		if count == length-k {
			b = node
		}
		count++
	}
	a.Val, b.Val = b.Val, a.Val
	return head
}
