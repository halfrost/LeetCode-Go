package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 此题和 143 题 Reorder List 思路基本一致
func isPalindrome234(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	res := true
	// 寻找中间结点
	p1 := head
	p2 := head
	for p2.Next != nil && p2.Next.Next != nil {
		p1 = p1.Next
		p2 = p2.Next.Next
	}

	// 反转链表后半部分  1->2->3->4->5->6 to 1->2->3->6->5->4
	preMiddle := p1
	preCurrent := p1.Next
	for preCurrent.Next != nil {
		current := preCurrent.Next
		preCurrent.Next = current.Next
		current.Next = preMiddle.Next
		preMiddle.Next = current
	}

	// 扫描表，判断是否是回文
	p1 = head
	p2 = preMiddle.Next
	// fmt.Printf("p1 = %v p2 = %v preMiddle = %v head = %v\n", p1.Val, p2.Val, preMiddle.Val, L2ss(head))
	for p1 != preMiddle {
		// fmt.Printf("*****p1 = %v p2 = %v preMiddle = %v head = %v\n", p1, p2, preMiddle, L2ss(head))
		if p1.Val == p2.Val {
			p1 = p1.Next
			p2 = p2.Next
			// fmt.Printf("-------p1 = %v p2 = %v preMiddle = %v head = %v\n", p1, p2, preMiddle, L2ss(head))
		} else {
			res = false
			break
		}
	}
	if p1 == preMiddle {
		if p2 != nil && p1.Val != p2.Val {
			return false
		}
	}

	return res
}

// L2ss define
func L2ss(head *ListNode) []int {
	res := []int{}

	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}

	return res
}
