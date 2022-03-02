package leetcode

type MyLinkedList struct {
	head *Node
}

type Node struct {
	Val  int
	Next *Node
	Prev *Node
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	curr := this.head
	for i := 0; i < index && curr != nil; i++ {
		curr = curr.Next
	}
	if curr != nil {
		return curr.Val
	} else {
		return -1
	}
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	node := &Node{Val: val}
	node.Next = this.head
	if this.head != nil {
		this.head.Prev = node
	}
	this.head = node
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	if this.head == nil {
		this.AddAtHead(val)
		return
	}
	node := &Node{Val: val}
	curr := this.head
	for curr != nil && curr.Next != nil {
		curr = curr.Next
	}
	node.Prev = curr
	curr.Next = node
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index == 0 {
		this.AddAtHead(val)
	} else {
		node := &Node{Val: val}
		curr := this.head
		for i := 0; i < index-1 && curr != nil; i++ {
			curr = curr.Next
		}
		if curr != nil {
			node.Next = curr.Next
			node.Prev = curr
			if node.Next != nil {
				node.Next.Prev = node
			}
			curr.Next = node
		}
	}
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index == 0 {
		this.head = this.head.Next
		if this.head != nil {
			this.head.Prev = nil
		}
	} else {
		curr := this.head
		for i := 0; i < index-1 && curr != nil; i++ {
			curr = curr.Next
		}
		if curr != nil && curr.Next != nil {
			curr.Next = curr.Next.Next
			if curr.Next != nil {
				curr.Next.Prev = curr
			}
		}
	}
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
