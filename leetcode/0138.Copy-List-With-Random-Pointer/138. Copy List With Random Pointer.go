package leetcode

// Node define
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	tempHead := copyNodeToLinkedList(head)
	return splitLinkedList(tempHead)
}

func splitLinkedList(head *Node) *Node {
	cur := head
	head = head.Next
	for cur != nil && cur.Next != nil {
		cur.Next, cur = cur.Next.Next, cur.Next
	}
	return head
}

func copyNodeToLinkedList(head *Node) *Node {
	cur := head
	for cur != nil {
		node := &Node{
			Val:  cur.Val,
			Next: cur.Next,
		}
		cur.Next, cur = node, cur.Next
	}
	cur = head
	for cur != nil {
		if cur.Random != nil {
			cur.Next.Random = cur.Random.Next
		}
		cur = cur.Next.Next
	}
	return head
}
