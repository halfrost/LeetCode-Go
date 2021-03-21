package leetcode

import (
	"container/list"
)

type FrontMiddleBackQueue struct {
	list   *list.List
	middle *list.Element
}

func Constructor() FrontMiddleBackQueue {
	return FrontMiddleBackQueue{list: list.New()}
}

func (this *FrontMiddleBackQueue) PushFront(val int) {
	e := this.list.PushFront(val)
	if this.middle == nil {
		this.middle = e
	} else if this.list.Len()%2 == 0 && this.middle.Prev() != nil {
		this.middle = this.middle.Prev()
	}
}

func (this *FrontMiddleBackQueue) PushMiddle(val int) {
	if this.middle == nil {
		this.PushFront(val)
	} else {
		if this.list.Len()%2 != 0 {
			this.middle = this.list.InsertBefore(val, this.middle)
		} else {
			this.middle = this.list.InsertAfter(val, this.middle)
		}
	}
}

func (this *FrontMiddleBackQueue) PushBack(val int) {
	e := this.list.PushBack(val)
	if this.middle == nil {
		this.middle = e
	} else if this.list.Len()%2 != 0 && this.middle.Next() != nil {
		this.middle = this.middle.Next()
	}
}

func (this *FrontMiddleBackQueue) PopFront() int {
	if this.list.Len() == 0 {
		return -1
	}
	e := this.list.Front()
	if this.list.Len() == 1 {
		this.middle = nil
	} else if this.list.Len()%2 == 0 && this.middle.Next() != nil {
		this.middle = this.middle.Next()
	}
	return this.list.Remove(e).(int)
}

func (this *FrontMiddleBackQueue) PopMiddle() int {
	if this.middle == nil {
		return -1
	}
	e := this.middle
	if this.list.Len()%2 != 0 {
		this.middle = e.Prev()
	} else {
		this.middle = e.Next()
	}
	return this.list.Remove(e).(int)
}

func (this *FrontMiddleBackQueue) PopBack() int {
	if this.list.Len() == 0 {
		return -1
	}
	e := this.list.Back()
	if this.list.Len() == 1 {
		this.middle = nil
	} else if this.list.Len()%2 != 0 && this.middle.Prev() != nil {
		this.middle = this.middle.Prev()
	}
	return this.list.Remove(e).(int)
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
