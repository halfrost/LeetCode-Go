package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	isCycle, slow := hasCycle142(head)
	if !isCycle {
		return nil
	}
	fast := head
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}
	return fast
}

func hasCycle142(head *ListNode) (bool, *ListNode) {
	fast := head
	slow := head
	for slow != nil && fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true, slow
		}
	}
	return false, nil
}
