package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 解法一 单链表
func partition(head *ListNode, x int) *ListNode {
	beforeHead := &ListNode{Val: 0, Next: nil}
	before := beforeHead
	afterHead := &ListNode{Val: 0, Next: nil}
	after := afterHead

	for head != nil {
		if head.Val < x {
			before.Next = head
			before = before.Next
		} else {
			after.Next = head
			after = after.Next
		}
		head = head.Next
	}
	after.Next = nil
	before.Next = afterHead.Next
	return beforeHead.Next
}

// DoublyListNode define
type DoublyListNode struct {
	Val  int
	Prev *DoublyListNode
	Next *DoublyListNode
}

// 解法二 双链表
func partition1(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	DLNHead := genDoublyListNode(head)
	cur := DLNHead
	for cur != nil {
		if cur.Val < x {
			tmp := &DoublyListNode{Val: cur.Val, Prev: nil, Next: nil}
			compareNode := cur
			for compareNode.Prev != nil {
				if compareNode.Val >= x && compareNode.Prev.Val < x {
					break
				}
				compareNode = compareNode.Prev
			}
			if compareNode == DLNHead {
				if compareNode.Val < x {
					cur = cur.Next
					continue
				} else {
					tmp.Next = DLNHead
					DLNHead.Prev = tmp
					DLNHead = tmp
				}
			} else {
				tmp.Next = compareNode
				tmp.Prev = compareNode.Prev
				compareNode.Prev.Next = tmp
				compareNode.Prev = tmp
			}
			deleteNode := cur
			if cur.Prev != nil {
				deleteNode.Prev.Next = deleteNode.Next
			}
			if cur.Next != nil {
				deleteNode.Next.Prev = deleteNode.Prev
			}
		}
		cur = cur.Next
	}
	return genListNode(DLNHead)
}

func genDoublyListNode(head *ListNode) *DoublyListNode {
	cur := head.Next
	DLNHead := &DoublyListNode{Val: head.Val, Prev: nil, Next: nil}
	curDLN := DLNHead
	for cur != nil {
		tmp := &DoublyListNode{Val: cur.Val, Prev: curDLN, Next: nil}
		curDLN.Next = tmp
		curDLN = tmp
		cur = cur.Next
	}
	return DLNHead
}

func genListNode(head *DoublyListNode) *ListNode {
	cur := head.Next
	LNHead := &ListNode{Val: head.Val, Next: nil}
	curLN := LNHead
	for cur != nil {
		tmp := &ListNode{Val: cur.Val, Next: nil}
		curLN.Next = tmp
		curLN = tmp
		cur = cur.Next
	}
	return LNHead
}
