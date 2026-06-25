---
title: 3.4 LFUCache
type: docs
weight: 4
---

# Least Frequently Used LFUCache

![](https://img.halfrost.com/Blog/ArticleImage/146_1_.png)

LFU is the abbreviation of Least Frequently Used, which is also a common page replacement algorithm. It selects the page with the smallest access counter for eviction. As shown below, each page in the cache has an access counter.


![](https://img.halfrost.com/Blog/ArticleImage/146_3.png)

According to the LFU strategy, the access counter must be updated on every access. When inserting B, B is found in the cache, so the access counter is incremented, and B is moved to the position sorted by access counter in descending order. Then insert D; similarly, first update the counter, then move it to its sorted position. When inserting F, F does not exist in the cache, so the page with the smallest counter is evicted, which means page A is evicted. At this point, F is at the bottom, with a count of 1.

![](https://img.halfrost.com/Blog/ArticleImage/146_8_.png)

There is something special here compared with LRU. If there are multiple pages to be evicted with the same number of accesses, choose the one closest to the tail. As shown above, A, B, and C all have the same number of accesses, all 1 time. To insert F, F is not in the cache, so page A must be evicted. F is the newly inserted page, with an access count of 1, and is placed before C. That is to say, for the same number of accesses, they are ordered by recency, and the oldest page is evicted. This is the biggest difference from LRU.

It can be seen that **LFU updates and new page insertions can occur at any position in the linked list, while page deletions all occur at the tail of the list**.


## Solution One Get O(1) / Put O(1)

LFU likewise requires queries to be as efficient as possible, querying within O(1). A map is still chosen for lookup. Modification and deletion also need to be completed in O(1), so a doubly linked list is still used, continuing to reuse the list data structure in the container package. LFU needs to record the number of accesses, so each node must store frequency, the number of accesses, in addition to key and value.

There is also one more issue to consider: how to sort by frequency? For the same frequency, sort by order of arrival. If you start considering sorting algorithms, your line of thinking has deviated from the best answer. Sorting is at least O(nlogn). Looking back again at how LFU works, you will find that it only cares about the minimum frequency. It does not care about the order among other frequencies. Therefore, sorting is not needed. Use a min variable to store the minimum frequency; when evicting, reading this minimum value can find the node to delete. The requirement that the same frequency be ordered by arrival order is still implemented with a doubly linked list. The insertion order of the doubly linked list reflects the order of nodes. The same frequency corresponds to one doubly linked list. Since there may be multiple identical frequencies, there may be multiple doubly linked lists. Use a map to maintain the correspondence between access frequency and doubly linked list. When deleting the minimum frequency, use min to find the minimum frequency, then find the doubly linked list corresponding to this frequency in this map, and delete the oldest node in the doubly linked list. This solves the LFU delete operation.

The update operation of LFU is similar to LRU. It also needs a map to store the mapping between key and doubly linked list node. This doubly linked list node stores a tuple of the three elements key-value-frequency. In this way, through the key and frequency in the node, the key in the map can be deleted in reverse.

Define the data structure of LFUCache as follows:

```go

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

```

The Get operation of LFUCache involves updating the frequency value and 2 maps. In the nodes map, node information is obtained by key. In lists, delete the node at the current frequency. After deletion, frequency ++. If the new frequency exists in lists, add it to the head of the doubly linked list; if it does not exist, a new doubly linked list needs to be created and the current node added to the head. Then update the map whose value is the doubly linked list node. Finally, update the min value, and determine whether the doubly linked list corresponding to the old frequency is already empty. If it is empty, min++.

```go
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
	newNode := newList.PushFront(currentNode)
	this.nodes[key] = newNode
	if currentNode.frequency-1 == this.min && this.lists[currentNode.frequency-1].Len() == 0 {
		this.min++
	}
	return currentNode.value
}

```

The logic of LFU's Put operation is a bit more involved. First, query whether the key exists in the nodes map. If it exists, obtain this node, update its value, and then manually call the Get operation once, because the update logic below is consistent with the Get operation. If it does not exist in the map, proceed with insertion or deletion. Determine whether capacity is full. If it is full, perform the delete operation. Delete the tail node in the doubly linked list corresponding to min, and correspondingly delete the key-value pair in the nodes map as well.

Since the access count of a newly inserted page must be 1, min is set to 1 at this time. Create a new node and insert it into the 2 maps.

```go

func (this *LFUCache) Put(key int, value int) {
	if this.capacity == 0 {
		return
	}
	// If it exists, update the access count
	if currentValue, ok := this.nodes[key]; ok {
		currentNode := currentValue.Value.(*node)
		currentNode.value = value
		this.Get(key)
		return
	}
	// If it does not exist and the cache is full, deletion is required
	if this.capacity == len(this.nodes) {
		currentList := this.lists[this.min]
		backNode := currentList.Back()
		delete(this.nodes, backNode.Value.(*node).key)
		currentList.Remove(backNode)
	}
	// Create a new node and insert it into 2 maps
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
	newNode := newList.PushFront(currentNode)
	this.nodes[key] = newNode
}

```

In summary, LFU is a data structure composed of two maps and a min pointer. In one map, the key stores the access count, and the corresponding value is a doubly linked list. Here, the role of the doubly linked list is to evict the oldest page at the tail when the frequency is the same. In the other map, the value corresponding to the key is a node of the doubly linked list. Compared with LRU, the node stores one additional value, the number of accesses, that is, the node stores a key-value-frequency tuple. Here, the role of the doubly linked list is similar to LRU: it can update the value and frequency values in the doubly linked list node according to the key in the map, and it can also update the corresponding relationship in the map in reverse according to the key and frequency in the doubly linked list node. As shown below:

![](https://img.halfrost.com/Blog/ArticleImage/146_10_0.png)

After submitting the code, it successfully passed all test cases.


![](https://img.halfrost.com/Blog/ArticleImage/146_5.png)


## Solution Two Get O(capacity) / Put O(capacity)

Another idea for LFU is to use the [Index Priority Queue](https://algs4.cs.princeton.edu/24pq/) data structure. Don't be intimidated by the name: Index Priority Queue = map + Priority Queue, nothing more.

Use Priority Queue to maintain a min-heap, where the heap top is the element with the smallest number of accesses. The value in the map stores the node in the priority queue.

```go
import "container/heap"

type LFUCache struct {
	capacity int
	pq       PriorityQueue
	hash     map[int]*Item
	counter  int
}

func Constructor(capacity int) LFUCache {
	lfu := LFUCache{
		pq:       PriorityQueue{},
		hash:     make(map[int]*Item, capacity),
		capacity: capacity,
	}
	return lfu
}

```

Get and Put operations should be as fast as possible, and there are 2 problems to solve. When the number of accesses is the same, how do we delete the oldest element? When an element's number of accesses changes, how do we quickly adjust the heap? To solve these 2 problems, define the following data structure:

```go
// An Item is something we manage in a priority queue.
type Item struct {
	value     int // The value of the item; arbitrary.
	key       int
	frequency int // The priority of the item in the queue.
	count     int // use for evicting the oldest element
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

```

The nodes in the heap store these 5 values. The count value is used to determine which element is the oldest, similar to an operation timestamp. The index value is used to re-heapify and adjust the heap. Next, implement the methods of PriorityQueue.

```go
// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	if pq[i].frequency == pq[j].frequency {
		return pq[i].count < pq[j].count
	}
	return pq[i].frequency < pq[j].frequency
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, value int, frequency int, count int) {
	item.value = value
	item.count = count
	item.frequency = frequency
	heap.Fix(pq, item.index)
}
```

In the Less() method, frequency is sorted from small to large; when frequency is the same, count is sorted from small to large. According to the heap-building rules of a priority queue, the smallest frequency will be at the heap top, and for the same frequency, the smaller the count, the closer it is to the heap top.

In the Swap() method, remember to update the index value. In the Push() method, when inserting, the length of the queue is the index value of this element, so remember to update the index value here as well. The update() method calls the Fix() function. The Fix() function costs less than first Remove() and then Push() a new value. Therefore, the Fix() function is called here, and the time complexity of this operation is O(log n).

This maintains the minimum Index Priority Queue. The Get operation is very simple:

```go
func (this *LFUCache) Get(key int) int {
	if this.capacity == 0 {
		return -1
	}
	if item, ok := this.hash[key]; ok {
		this.counter++
		this.pq.update(item, item.value, item.frequency+1, this.counter)
		return item.value
	}
	return -1
}

```

Query the key in the hashmap. If it exists, increment the counter timestamp and call the update method of Priority Queue to adjust the heap.

```go
func (this *LFUCache) Put(key int, value int) {
	if this.capacity == 0 {
		return
	}
	this.counter++
	// If it exists, increase frequency, then adjust the heap
	if item, ok := this.hash[key]; ok {
		this.pq.update(item, value, item.frequency+1, this.counter)
		return
	}
	// If it does not exist and the cache is full, deletion is required. Delete it from hashmap and pq.
	if len(this.pq) == this.capacity {
		item := heap.Pop(&this.pq).(*Item)
		delete(this.hash, item.key)
	}
	// Create a new node and add it to hashmap and pq.
	item := &Item{
		value: value,
		key:   key,
		count: this.counter,
	}
	heap.Push(&this.pq, item)
	this.hash[key] = item
}
```


For LFU implemented with a min-heap, the Put time complexity is O(capacity), and the Get time complexity is O(capacity), which is not as good as the version implemented with 2 maps. Coincidentally, the min-heap version actually beat 100%.

![](https://img.halfrost.com/Blog/ArticleImage/146_7.png)


## Template


```go
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
