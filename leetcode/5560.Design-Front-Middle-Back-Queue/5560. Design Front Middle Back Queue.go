package leetcode

type FrontMiddleBackQueue struct {
	Queue  []int
	Length int
}

func Constructor() FrontMiddleBackQueue {
	return FrontMiddleBackQueue{Queue: make([]int, 0), Length: 0}
}

func (this *FrontMiddleBackQueue) PushFront(val int) {
	tmp := make([]int, this.Length+1)
	copy(tmp[1:], this.Queue)
	tmp[0] = val
	this.Queue = tmp
	this.Length++
}

func (this *FrontMiddleBackQueue) PushMiddle(val int) {
	tmp := make([]int, this.Length+1)
	idx := this.Length / 2
	copy(tmp[:idx], this.Queue[:idx])
	tmp[idx] = val
	copy(tmp[idx+1:], this.Queue[idx:])
	this.Queue = tmp
	this.Length++
}

func (this *FrontMiddleBackQueue) PushBack(val int) {
	this.Queue = append(this.Queue, val)
	this.Length++
}

func (this *FrontMiddleBackQueue) PopFront() int {
	if this.Length == 0 {
		return -1
	}
	res := this.Queue[0]
	this.Queue = this.Queue[1:]
	this.Length--
	return res
}

func (this *FrontMiddleBackQueue) PopMiddle() int {
	if this.Length == 0 {
		return -1
	}
	mid := (this.Length - 1) / 2
	res := this.Queue[mid]
	tmp := make([]int, len(this.Queue)-1)
	copy(tmp[:mid], this.Queue[:mid])
	copy(tmp[mid:], this.Queue[mid+1:])
	this.Queue = tmp
	this.Length--
	return res
}

func (this *FrontMiddleBackQueue) PopBack() int {
	if this.Length == 0 {
		return -1
	}
	res := this.Queue[this.Length-1]
	this.Queue = this.Queue[:this.Length-1]
	this.Length--
	return res
}

/**
 * Your FrontMiddleBackQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PushFront(val);
 * obj.PushMiddle(val);
 * obj.PushBack(val);
 * param_4 := obj.PopFront();
 * param_5 := obj.PopMiddle();
 * param_6 := obj.PopBack();
 */
