package leetcode

type MyStack struct {
	enque []int
	deque []int
}

/** Initialize your data structure here. */
func Constructor225() MyStack {
	return MyStack{[]int{}, []int{}}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.enque = append(this.enque, x)
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	length := len(this.enque)
	for i := 0; i < length-1; i++ {
		this.deque = append(this.deque, this.enque[0])
		this.enque = this.enque[1:]
	}
	topEle := this.enque[0]
	this.enque = this.deque
	this.deque = nil

	return topEle
}

/** Get the top element. */
func (this *MyStack) Top() int {
	topEle := this.Pop()
	this.enque = append(this.enque, topEle)

	return topEle
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	if len(this.enque) == 0 {
		return true
	}

	return false
}
