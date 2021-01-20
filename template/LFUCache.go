package template

import "container/list"

// LFUCache define
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

// Constructor define
func Constructor(capacity int) LFUCache {
	return LFUCache{nodes: make(map[int]*list.Element),
		lists:    make(map[int]*list.List),
		capacity: capacity,
		min:      0,
	}
}

// Get define
func (lfuCache *LFUCache) Get(key int) int {
	value, ok := lfuCache.nodes[key]
	if !ok {
		return -1
	}
	currentNode := value.Value.(*node)
	lfuCache.lists[currentNode.frequency].Remove(value)
	currentNode.frequency++
	if _, ok := lfuCache.lists[currentNode.frequency]; !ok {
		lfuCache.lists[currentNode.frequency] = list.New()
	}
	newList := lfuCache.lists[currentNode.frequency]
	newNode := newList.PushFront(currentNode)
	lfuCache.nodes[key] = newNode
	if currentNode.frequency-1 == lfuCache.min && lfuCache.lists[currentNode.frequency-1].Len() == 0 {
		lfuCache.min++
	}
	return currentNode.value
}

// Put define
func (lfuCache *LFUCache) Put(key int, value int) {
	if lfuCache.capacity == 0 {
		return
	}
	if currentValue, ok := lfuCache.nodes[key]; ok {
		currentNode := currentValue.Value.(*node)
		currentNode.value = value
		lfuCache.Get(key)
		return
	}
	if lfuCache.capacity == len(lfuCache.nodes) {
		currentList := lfuCache.lists[lfuCache.min]
		backNode := currentList.Back()
		delete(lfuCache.nodes, backNode.Value.(*node).key)
		currentList.Remove(backNode)
	}
	lfuCache.min = 1
	currentNode := &node{
		key:       key,
		value:     value,
		frequency: 1,
	}
	if _, ok := lfuCache.lists[1]; !ok {
		lfuCache.lists[1] = list.New()
	}
	newList := lfuCache.lists[1]
	newNode := newList.PushFront(currentNode)
	lfuCache.nodes[key] = newNode
}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

// Index Priority Queue
// import "container/heap"

// type LFUCache struct {
// 	capacity int
// 	pq       PriorityQueue
// 	hash     map[int]*Item
// 	counter  int
// }

// func Constructor(capacity int) LFUCache {
// 	lfu := LFUCache{
// 		pq:       PriorityQueue{},
// 		hash:     make(map[int]*Item, capacity),
// 		capacity: capacity,
// 	}
// 	return lfu
// }

// func (this *LFUCache) Get(key int) int {
// 	if this.capacity == 0 {
// 		return -1
// 	}
// 	if item, ok := this.hash[key]; ok {
// 		this.counter++
// 		this.pq.update(item, item.value, item.frequency+1, this.counter)
// 		return item.value
// 	}
// 	return -1
// }

// func (this *LFUCache) Put(key int, value int) {
// 	if this.capacity == 0 {
// 		return
// 	}
// 	// fmt.Printf("Put %d\n", key)
// 	this.counter++
// 	// 如果存在，增加 frequency，再调整堆
// 	if item, ok := this.hash[key]; ok {
// 		this.pq.update(item, value, item.frequency+1, this.counter)
// 		return
// 	}
// 	// 如果不存在且缓存满了，需要删除。在 hashmap 和 pq 中删除。
// 	if len(this.pq) == this.capacity {
// 		item := heap.Pop(&this.pq).(*Item)
// 		delete(this.hash, item.key)
// 	}
// 	// 新建结点，在 hashmap 和 pq 中添加。
// 	item := &Item{
// 		value: value,
// 		key:   key,
// 		count: this.counter,
// 	}
// 	heap.Push(&this.pq, item)
// 	this.hash[key] = item
// }

// // An Item is something we manage in a priority queue.
// type Item struct {
// 	value     int // The value of the item; arbitrary.
// 	key       int
// 	frequency int // The priority of the item in the queue.
// 	count     int // use for evicting the oldest element
// 	// The index is needed by update and is maintained by the heap.Interface methods.
// 	index int // The index of the item in the heap.
// }

// // A PriorityQueue implements heap.Interface and holds Items.
// type PriorityQueue []*Item

// func (pq PriorityQueue) Len() int { return len(pq) }

// func (pq PriorityQueue) Less(i, j int) bool {
// 	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
// 	if pq[i].frequency == pq[j].frequency {
// 		return pq[i].count < pq[j].count
// 	}
// 	return pq[i].frequency < pq[j].frequency
// }

// func (pq PriorityQueue) Swap(i, j int) {
// 	pq[i], pq[j] = pq[j], pq[i]
// 	pq[i].index = i
// 	pq[j].index = j
// }

// func (pq *PriorityQueue) Push(x interface{}) {
// 	n := len(*pq)
// 	item := x.(*Item)
// 	item.index = n
// 	*pq = append(*pq, item)
// }

// func (pq *PriorityQueue) Pop() interface{} {
// 	old := *pq
// 	n := len(old)
// 	item := old[n-1]
// 	old[n-1] = nil  // avoid memory leak
// 	item.index = -1 // for safety
// 	*pq = old[0 : n-1]
// 	return item
// }

// // update modifies the priority and value of an Item in the queue.
// func (pq *PriorityQueue) update(item *Item, value int, frequency int, count int) {
// 	item.value = value
// 	item.count = count
// 	item.frequency = frequency
// 	heap.Fix(pq, item.index)
// }
