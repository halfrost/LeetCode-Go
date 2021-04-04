package leetcode

type MyCircularQueue struct {
	cap   int
	size  int
	queue []int
	left  int
	right int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{cap: k, size: 0, left: 0, right: 0, queue: make([]int, k)}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.size == this.cap {
		return false
	}
	this.size++
	this.queue[this.right] = value
	this.right++
	this.right %= this.cap
	return true

}

func (this *MyCircularQueue) DeQueue() bool {
	if this.size == 0 {
		return false
	}
	this.size--
	this.left++
	this.left %= this.cap
	return true
}

func (this *MyCircularQueue) Front() int {
	if this.size == 0 {
		return -1
	}
	return this.queue[this.left]
}

func (this *MyCircularQueue) Rear() int {
	if this.size == 0 {
		return -1
	}
	if this.right == 0 {
		return this.queue[this.cap-1]
	}
	return this.queue[this.right-1]
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.size == 0
}

func (this *MyCircularQueue) IsFull() bool {
	return this.size == this.cap
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
