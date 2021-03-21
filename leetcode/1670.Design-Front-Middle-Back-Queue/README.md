# [1670. Design Front Middle Back Queue](https://leetcode.com/problems/design-front-middle-back-queue/)


## 题目

Design a queue that supports `push` and `pop` operations in the front, middle, and back.

Implement the `FrontMiddleBack` class:

- `FrontMiddleBack()` Initializes the queue.
- `void pushFront(int val)` Adds `val` to the **front** of the queue.
- `void pushMiddle(int val)` Adds `val` to the **middle** of the queue.
- `void pushBack(int val)` Adds `val` to the **back** of the queue.
- `int popFront()` Removes the **front** element of the queue and returns it. If the queue is empty, return `1`.
- `int popMiddle()` Removes the **middle** element of the queue and returns it. If the queue is empty, return `1`.
- `int popBack()` Removes the **back** element of the queue and returns it. If the queue is empty, return `1`.

**Notice** that when there are **two** middle position choices, the operation is performed on the **frontmost** middle position choice. For example:

- Pushing `6` into the middle of `[1, 2, 3, 4, 5]` results in `[1, 2, 6, 3, 4, 5]`.
- Popping the middle from `[1, 2, 3, 4, 5, 6]` returns `3` and results in `[1, 2, 4, 5, 6]`.

**Example 1:**

```
Input:
["FrontMiddleBackQueue", "pushFront", "pushBack", "pushMiddle", "pushMiddle", "popFront", "popMiddle", "popMiddle", "popBack", "popFront"]
[[], [1], [2], [3], [4], [], [], [], [], []]
Output:
[null, null, null, null, null, 1, 3, 4, 2, -1]

Explanation:
FrontMiddleBackQueue q = new FrontMiddleBackQueue();
q.pushFront(1);   // [1]
q.pushBack(2);    // [1, 2]
q.pushMiddle(3);  // [1, 3, 2]
q.pushMiddle(4);  // [1, 4, 3, 2]
q.popFront();     // return 1 -> [4, 3, 2]
q.popMiddle();    // return 3 -> [4, 2]
q.popMiddle();    // return 4 -> [2]
q.popBack();      // return 2 -> []
q.popFront();     // return -1 -> [] (The queue is empty)

```

**Constraints:**

- `1 <= val <= 109`
- At most `1000` calls will be made to `pushFront`, `pushMiddle`, `pushBack`, `popFront`, `popMiddle`, and `popBack`.

## 题目大意

请你设计一个队列，支持在前，中，后三个位置的 push 和 pop 操作。

请你完成 FrontMiddleBack 类：

- FrontMiddleBack() 初始化队列。
- void pushFront(int val) 将 val 添加到队列的 最前面 。
- void pushMiddle(int val) 将 val 添加到队列的 正中间 。
- void pushBack(int val) 将 val 添加到队里的 最后面 。
- int popFront() 将 最前面 的元素从队列中删除并返回值，如果删除之前队列为空，那么返回 -1 。
- int popMiddle() 将 正中间 的元素从队列中删除并返回值，如果删除之前队列为空，那么返回 -1 。
- int popBack() 将 最后面 的元素从队列中删除并返回值，如果删除之前队列为空，那么返回 -1 。

请注意当有 两个 中间位置的时候，选择靠前面的位置进行操作。比方说：

- 将 6 添加到 [1, 2, 3, 4, 5] 的中间位置，结果数组为 [1, 2, 6, 3, 4, 5] 。
- 从 [1, 2, 3, 4, 5, 6] 的中间位置弹出元素，返回 3 ，数组变为 [1, 2, 4, 5, 6] 。

## 解题思路

- 简单题，利用 go 原生的双向队列 list 的实现，可以轻松实现这个“前中后队列”。
- 具体实现见代码，几组特殊测试用例见测试文件。

## 代码

```go
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
```