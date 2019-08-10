package leetcode

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
