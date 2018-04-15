package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// ListNode define
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var behind *ListNode
	for head != nil {
		next := head.Next
		head.Next = behind
		behind = head
		head = next
	}
	return behind
}
