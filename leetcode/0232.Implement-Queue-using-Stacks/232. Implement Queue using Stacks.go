package leetcode

type MyQueue struct {
	Stack *[]int
	Queue *[]int
}

/** Initialize your data structure here. */
func Constructor232() MyQueue {
	tmp1, tmp2 := []int{}, []int{}
	return MyQueue{Stack: &tmp1, Queue: &tmp2}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	*this.Stack = append(*this.Stack, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if len(*this.Queue) == 0 {
		this.fromStackToQueue(this.Stack, this.Queue)
	}

	popped := (*this.Queue)[len(*this.Queue)-1]
	*this.Queue = (*this.Queue)[:len(*this.Queue)-1]
	return popped
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if len(*this.Queue) == 0 {
		this.fromStackToQueue(this.Stack, this.Queue)
	}

	return (*this.Queue)[len(*this.Queue)-1]
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(*this.Stack)+len(*this.Queue) == 0
}

func (this *MyQueue) fromStackToQueue(s, q *[]int) {
	for len(*s) > 0 {
		popped := (*s)[len(*s)-1]
		*s = (*s)[:len(*s)-1]

		*q = append(*q, popped)
	}
}
