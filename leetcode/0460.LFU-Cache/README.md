# [460. LFU Cache](https://leetcode.com/problems/lfu-cache/)


## 题目

Design and implement a data structure for [Least Frequently Used (LFU)](https://en.wikipedia.org/wiki/Least_frequently_used) cache.

Implement the `LFUCache` class:

- `LFUCache(int capacity)` Initializes the object with the `capacity` of the data structure.
- `int get(int key)` Gets the value of the `key` if the `key` exists in the cache. Otherwise, returns `1`.
- `void put(int key, int value)` Sets or inserts the value if the `key` is not already present. When the cache reaches its `capacity`, it should invalidate the least frequently used item before inserting a new item. For this problem, when there is a tie (i.e., two or more keys with the same frequency), **the least recently** used `key` would be evicted.

**Notice that** the number of times an item is used is the number of calls to the `get` and `put` functions for that item since it was inserted. This number is set to zero when the item is removed.

**Example 1:**

```
Input
["LFUCache", "put", "put", "get", "put", "get", "get", "put", "get", "get", "get"]
[[2], [1, 1], [2, 2], [1], [3, 3], [2], [3], [4, 4], [1], [3], [4]]
Output
[null, null, null, 1, null, -1, 3, null, -1, 3, 4]

Explanation
LFUCache lfu = new LFUCache(2);
lfu.put(1, 1);
lfu.put(2, 2);
lfu.get(1);      // return 1
lfu.put(3, 3);   // evicts key 2
lfu.get(2);      // return -1 (not found)
lfu.get(3);      // return 3
lfu.put(4, 4);   // evicts key 1.
lfu.get(1);      // return -1 (not found)
lfu.get(3);      // return 3
lfu.get(4);      // return 4

```

**Constraints:**

- `0 <= capacity, key, value <= 104`
- At most `10^5` calls will be made to `get` and `put`.

**Follow up:** Could you do both operations in `O(1)` time complexity?

## 题目大意

请你为 最不经常使用（LFU）缓存算法设计并实现数据结构。

实现 LFUCache 类：

- LFUCache(int capacity) - 用数据结构的容量 capacity 初始化对象
- int get(int key) - 如果键存在于缓存中，则获取键的值，否则返回 -1。
- void put(int key, int value) - 如果键已存在，则变更其值；如果键不存在，请插入键值对。当缓存达到其容量时，则应该在插入新项之前，使最不经常使用的项无效。在此问题中，当存在平局（即两个或更多个键具有相同使用频率）时，应该去除 最久未使用 的键。

注意「项的使用次数」就是自插入该项以来对其调用 get 和 put 函数的次数之和。使用次数会在对应项被移除后置为 0 。

进阶：你是否可以在 O(1) 时间复杂度内执行两项操作？

## 解题思路

- 这一题是 LFU 经典面试题，详细解释见[第三章 LFUCache 模板](https://books.halfrost.com/leetcode/ChapterThree/LFUCache/)。

## 代码

```go
package leetcode

import "container/list"

type LFUCache struct {
	nodes    map[int]*list.Element
	lists    map[int]*list.List
	capacity int
	min      int
}

type node struct {
	key       int
	value     int
	frequency int
}

func Constructor(capacity int) LFUCache {
	return LFUCache{nodes: make(map[int]*list.Element),
		lists:    make(map[int]*list.List),
		capacity: capacity,
		min:      0,
	}
}

func (this *LFUCache) Get(key int) int {
	value, ok := this.nodes[key]
	if !ok {
		return -1
	}
	currentNode := value.Value.(*node)
	this.lists[currentNode.frequency].Remove(value)
	currentNode.frequency++
	if _, ok := this.lists[currentNode.frequency]; !ok {
		this.lists[currentNode.frequency] = list.New()
	}
	newList := this.lists[currentNode.frequency]
	newNode := newList.PushBack(currentNode)
	this.nodes[key] = newNode
	if currentNode.frequency-1 == this.min && this.lists[currentNode.frequency-1].Len() == 0 {
		this.min++
	}
	return currentNode.value
}

func (this *LFUCache) Put(key int, value int) {
	if this.capacity == 0 {
		return
	}
	if currentValue, ok := this.nodes[key]; ok {
		currentNode := currentValue.Value.(*node)
		currentNode.value = value
		this.Get(key)
		return
	}
	if this.capacity == len(this.nodes) {
		currentList := this.lists[this.min]
		frontNode := currentList.Front()
		delete(this.nodes, frontNode.Value.(*node).key)
		currentList.Remove(frontNode)
	}
	this.min = 1
	currentNode := &node{
		key:       key,
		value:     value,
		frequency: 1,
	}
	if _, ok := this.lists[1]; !ok {
		this.lists[1] = list.New()
	}
	newList := this.lists[1]
	newNode := newList.PushBack(currentNode)
	this.nodes[key] = newNode
}
```