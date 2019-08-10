package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	s := head.Next
	var behind *ListNode
	for head.Next != nil {
		headNext := head.Next
		if behind != nil && behind.Next != nil {
			behind.Next = headNext
		}
		var next *ListNode
		if head.Next.Next != nil {
			next = head.Next.Next
		}
		if head.Next.Next != nil {
			head.Next = next
		} else {
			head.Next = nil
		}
		headNext.Next = head
		behind = head
		if head.Next != nil {
			head = next
		}
	}
	return s
}
