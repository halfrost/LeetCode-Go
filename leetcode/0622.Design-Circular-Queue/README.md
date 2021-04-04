# [622. Design Circular Queue](https://leetcode.com/problems/design-circular-queue/)


## 题目

Design your implementation of the circular queue. The circular queue is a linear data structure in which the operations are performed based on FIFO (First In First Out) principle and the last position is connected back to the first position to make a circle. It is also called "Ring Buffer".

One of the benefits of the circular queue is that we can make use of the spaces in front of the queue. In a normal queue, once the queue becomes full, we cannot insert the next element even if there is a space in front of the queue. But using the circular queue, we can use the space to store new values.

Implementation the `MyCircularQueue` class:

- `MyCircularQueue(k)` Initializes the object with the size of the queue to be `k`.
- `int Front()` Gets the front item from the queue. If the queue is empty, return `1`.
- `int Rear()` Gets the last item from the queue. If the queue is empty, return `1`.
- `boolean enQueue(int value)` Inserts an element into the circular queue. Return `true` if the operation is successful.
- `boolean deQueue()` Deletes an element from the circular queue. Return `true` if the operation is successful.
- `boolean isEmpty()` Checks whether the circular queue is empty or not.
- `boolean isFull()` Checks whether the circular queue is full or not.

**Example 1:**

```
Input
["MyCircularQueue", "enQueue", "enQueue", "enQueue", "enQueue", "Rear", "isFull", "deQueue", "enQueue", "Rear"]
[[3], [1], [2], [3], [4], [], [], [], [4], []]
Output
[null, true, true, true, false, 3, true, true, true, 4]

Explanation
MyCircularQueue myCircularQueue = new MyCircularQueue(3);
myCircularQueue.enQueue(1); // return True
myCircularQueue.enQueue(2); // return True
myCircularQueue.enQueue(3); // return True
myCircularQueue.enQueue(4); // return False
myCircularQueue.Rear();     // return 3
myCircularQueue.isFull();   // return True
myCircularQueue.deQueue();  // return True
myCircularQueue.enQueue(4); // return True
myCircularQueue.Rear();     // return 4

```

**Constraints:**

- `1 <= k <= 1000`
- `0 <= value <= 1000`
- At most `3000` calls will be made to `enQueue`, `deQueue`, `Front`, `Rear`, `isEmpty`, and `isFull`.

**Follow up:**

Could you solve the problem without using the built-in queue?

## 题目大意

设计你的循环队列实现。 循环队列是一种线性数据结构，其操作表现基于 FIFO（先进先出）原则并且队尾被连接在队首之后以形成一个循环。它也被称为“环形缓冲器”。

循环队列的一个好处是我们可以利用这个队列之前用过的空间。在一个普通队列里，一旦一个队列满了，我们就不能插入下一个元素，即使在队列前面仍有空间。但是使用循环队列，我们能使用这些空间去存储新的值。

你的实现应该支持如下操作：

- MyCircularQueue(k): 构造器，设置队列长度为 k 。
- Front: 从队首获取元素。如果队列为空，返回 -1 。
- Rear: 获取队尾元素。如果队列为空，返回 -1 。
- enQueue(value): 向循环队列插入一个元素。如果成功插入则返回真。
- deQueue(): 从循环队列中删除一个元素。如果成功删除则返回真。
- isEmpty(): 检查循环队列是否为空。
- isFull(): 检查循环队列是否已满。

## 解题思路

- 简单题。设计一个环形队列，底层用数组实现。额外维护 4 个变量，队列的总 cap，队列当前的 size，前一元素下标 left，后一个元素下标 right。每添加一个元素便维护 left，right，size，下标需要对 cap 取余，因为超过 cap 大小之后，需要循环存储。代码实现没有难度，具体sh见下面代码。

## 代码

```go
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
```