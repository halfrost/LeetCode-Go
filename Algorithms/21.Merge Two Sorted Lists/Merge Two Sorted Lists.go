package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists(l1, l2.Next)
		return l2
	}
	// if l1 == nil && l2 == nil {
	// 	return nil
	// }
	// head := &ListNode{Val: 0, Next: nil}
	// current := head
	// var currentl1, currentl2 *ListNode
	// for l1 != nil || l2 != nil {
	// 	if l1 != nil {
	// 		currentl1 = &ListNode{Val: l1.Val, Next: nil}
	// 	}
	// 	if l2 != nil {
	// 		currentl2 = &ListNode{Val: l2.Val, Next: nil}
	// 	}

	// 	if l1 == nil && l2 != nil {
	// 		current.Next = currentl2
	// 		current = currentl2
	// 		l2 = l2.Next
	// 		continue
	// 	} else if l2 == nil && l1 != nil {
	// 		current.Next = currentl1
	// 		current = currentl1
	// 		l1 = l1.Next
	// 		continue
	// 	} else if l2 == nil && l1 == nil {
	// 		return head.Next
	// 	} else {
	// 		if currentl1.Val < currentl2.Val {
	// 			current.Next = currentl1
	// 			current = currentl1
	// 			l1 = l1.Next
	// 			continue
	// 		} else if currentl1.Val == currentl2.Val {
	// 			current.Next = currentl2
	// 			currentl2.Next = currentl1
	// 			current = currentl1
	// 			l1 = l1.Next
	// 			l2 = l2.Next
	// 			continue
	// 		} else {
	// 			current.Next = currentl2
	// 			current = currentl2
	// 			l2 = l2.Next
	// 			continue
	// 		}
	// 	}
	// }
	// return head.Next
}
