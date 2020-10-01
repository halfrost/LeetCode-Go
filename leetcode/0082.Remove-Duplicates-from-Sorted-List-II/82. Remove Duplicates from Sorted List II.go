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

func deleteDuplicates(head *ListNode) *ListNode {
	cur := head
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	for cur.Next != nil {
		if cur.Next.Val == cur.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}

// 双循环简单解法 O(n*m)
func deleteDuplicates3(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	nilNode := &ListNode{Val: 0, Next: head}
	head = nilNode

	lastVal := 0
	for head.Next != nil && head.Next.Next != nil {
		if head.Next.Val == head.Next.Next.Val {
			lastVal = head.Next.Val
			for head.Next != nil && lastVal == head.Next.Val {
				head.Next = head.Next.Next
			}
		} else {
			head = head.Next
		}
	}
	return nilNode.Next
}

// 双指针+删除标志位，单循环解法 O(n)
func deleteDuplicates4(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	nilNode := &ListNode{Val: 0, Next: head}
	// 上次遍历有删除操作的标志位
	lastIsDel := false
	// 虚拟空结点
	head = nilNode
	// 前后指针用于判断
	pre, back := head.Next, head.Next.Next
	// 每次只删除前面的一个重复的元素，留一个用于下次遍历判重
	// pre, back 指针的更新位置和值比较重要和巧妙
	for head.Next != nil && head.Next.Next != nil {
		if pre.Val != back.Val && lastIsDel {
			head.Next = head.Next.Next
			pre, back = head.Next, head.Next.Next
			lastIsDel = false
			continue
		}

		if pre.Val == back.Val {
			head.Next = head.Next.Next
			pre, back = head.Next, head.Next.Next
			lastIsDel = true
		} else {
			head = head.Next
			pre, back = head.Next, head.Next.Next
			lastIsDel = false
		}
	}
	// 处理 [1,1] 这种删除还剩一个的情况
	if lastIsDel && head.Next != nil {
		head.Next = nil
	}
	return nilNode.Next
}
