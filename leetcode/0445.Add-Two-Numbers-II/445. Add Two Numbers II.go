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
func addTwoNumbers445(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	l1Length := getLength(l1)
	l2Length := getLength(l2)
	newHeader := &ListNode{Val: 1, Next: nil}
	if l1Length < l2Length {
		newHeader.Next = addNode(l2, l1, l2Length-l1Length)
	} else {
		newHeader.Next = addNode(l1, l2, l1Length-l2Length)
	}
	if newHeader.Next.Val > 9 {
		newHeader.Next.Val = newHeader.Next.Val % 10
		return newHeader
	}
	return newHeader.Next
}

func addNode(l1 *ListNode, l2 *ListNode, offset int) *ListNode {
	if l1 == nil {
		return nil
	}
	var (
		res, node *ListNode
	)
	if offset == 0 {
		res = &ListNode{Val: l1.Val + l2.Val, Next: nil}
		node = addNode(l1.Next, l2.Next, 0)
	} else {
		res = &ListNode{Val: l1.Val, Next: nil}
		node = addNode(l1.Next, l2, offset-1)
	}
	if node != nil && node.Val > 9 {
		res.Val++
		node.Val = node.Val % 10
	}
	res.Next = node
	return res
}

func getLength(l *ListNode) int {
	count := 0
	cur := l
	for cur != nil {
		count++
		cur = cur.Next
	}
	return count
}

func addTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {
	reservedL1 := reverseList(l1)
	reservedL2 := reverseList(l2)

	dummyHead := &ListNode{}
	head := dummyHead
	carry := 0
	for reservedL1 != nil || reservedL2 != nil || carry > 0 {
		val := carry
		if reservedL1 != nil {
			val = reservedL1.Val + val
			reservedL1 = reservedL1.Next
		}
		if reservedL2 != nil {
			val = reservedL2.Val + val
			reservedL2 = reservedL2.Next
		}
		carry = val / 10
		head.Next = &ListNode{Val: val % 10}
		head = head.Next
	}
	return reverseList(dummyHead.Next)
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		tmp := head.Next
		head.Next = prev

		prev = head
		head = tmp
	}
	return prev
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	stack1 := pushStack(l1)
	stack2 := pushStack(l2)

	dummyHead := &ListNode{}
	head := dummyHead
	carry := 0
	for len(stack1) > 0 || len(stack2) > 0 || carry > 0 {
		val := carry
		if len(stack1) > 0 {
			val = val + stack1[len(stack1)-1]
			stack1 = stack1[:len(stack1)-1]
		}
		if len(stack2) > 0 {
			val = val + stack2[len(stack2)-1]
			stack2 = stack2[:len(stack2)-1]
		}
		carry = val / 10
		tmp := head.Next
		head.Next = &ListNode{Val: val % 10, Next: tmp}
	}
	return dummyHead.Next
}

func pushStack(l *ListNode) []int {
	var stack []int
	for l != nil {
		stack = append(stack, l.Val)
		l = l.Next
	}
	return stack
}
