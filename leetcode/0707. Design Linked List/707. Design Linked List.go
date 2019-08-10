package leetcode

type MyLinkedList struct {
	Val  int
	Next *MyLinkedList
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{Val: -999, Next: nil}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	cur := this
	for i := 0; cur != nil; i++ {
		if i == index {
			if cur.Val == -999 {
				return -1
			} else {
				return cur.Val
			}
		}
		cur = cur.Next
	}
	return -1
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	if this.Val == -999 {
		this.Val = val
	} else {
		tmp := &MyLinkedList{Val: this.Val, Next: this.Next}
		this.Val = val
		this.Next = tmp
	}
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	cur := this
	for cur.Next != nil {
		cur = cur.Next
	}
	tmp := &MyLinkedList{Val: val, Next: nil}
	cur.Next = tmp
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	cur := this
	if index == 0 {
		this.AddAtHead(val)
		return
	}
	for i := 0; cur != nil; i++ {
		if i == index-1 {
			break
		}
		cur = cur.Next
	}
	if cur != nil && cur.Val != -999 {
		tmp := &MyLinkedList{Val: val, Next: cur.Next}
		cur.Next = tmp
	}
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	cur := this
	for i := 0; cur != nil; i++ {
		if i == index-1 {
			break
		}
		cur = cur.Next
	}
	if cur != nil && cur.Next != nil {
		cur.Next = cur.Next.Next
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
