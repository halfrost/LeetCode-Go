package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNode(node *ListNode) {
	if node == nil {
		return
	}
	cur := node
	for cur.Next.Next != nil {
		cur.Val = cur.Next.Val
		cur = cur.Next
	}
	cur.Val = cur.Next.Val
	cur.Next = nil
}
