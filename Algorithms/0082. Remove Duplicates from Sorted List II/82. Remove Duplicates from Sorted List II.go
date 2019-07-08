package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates1(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	newHead := &ListNode{Next: head, Val: -999999}
	cur := newHead
	last := newHead
	front := head
	for front.Next != nil {
		if front.Val == cur.Val {
			// fmt.Printf("相同节点front = %v | cur = %v | last = %v\n", front.Val, cur.Val, last.Val)
			front = front.Next
			continue
		} else {
			if cur.Next != front {
				// fmt.Printf("删除重复节点front = %v | cur = %v | last = %v\n", front.Val, cur.Val, last.Val)
				last.Next = front
				if front.Next != nil && front.Next.Val != front.Val {
					last = front
				}
				cur = front
				front = front.Next
			} else {
				// fmt.Printf("常规循环前front = %v | cur = %v | last = %v\n", front.Val, cur.Val, last.Val)
				last = cur
				cur = cur.Next
				front = front.Next
				// fmt.Printf("常规循环后front = %v | cur = %v | last = %v\n", front.Val, cur.Val, last.Val)

			}
		}
	}
	if front.Val == cur.Val {
		// fmt.Printf("相同节点front = %v | cur = %v | last = %v\n", front.Val, cur.Val, last.Val)
		last.Next = nil
	} else {
		if cur.Next != front {
			last.Next = front
		}
	}
	return newHead.Next
}

func deleteDuplicates2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next != nil && head.Val == head.Next.Val {
		for head.Next != nil && head.Val == head.Next.Val {
			head = head.Next
		}
		return deleteDuplicates(head.Next)
	}
	head.Next = deleteDuplicates(head.Next)
	return head
}
