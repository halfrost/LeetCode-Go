---
title: 3.4 LFUCache
type: docs
weight: 4
---

# 最不经常最少使用 LFUCache

![](https://img.halfrost.com/Blog/ArticleImage/146_1_.png)

LFU 是 Least Frequently Used 的缩写，即最不经常最少使用，也是一种常用的页面置换算法，选择访问计数器最小的页面予以淘汰。如下图，缓存中每个页面带一个访问计数器。


![](https://img.halfrost.com/Blog/ArticleImage/146_3.png)

根据 LFU 的策略，每访问一次都要更新访问计数器。当插入 B 的时候，发现缓存中有 B，所以增加访问计数器的计数，并把 B 移动到访问计数器从大到小排序的地方。再插入 D，同理先更新计数器，再移动到它排序以后的位置。当插入 F 的时候，缓存中不存在 F，所以淘汰计数器最小的页面的页面，所以淘汰 A 页面。此时 F 排在最下面，计数为 1。

![](https://img.halfrost.com/Blog/ArticleImage/146_8_.png)

这里有一个比 LRU 特别的地方。如果淘汰的页面访问次数有多个相同的访问次数，选择最靠尾部的。如上图中，A、B、C 三者的访问次数相同，都是 1 次。要插入 F，F 不在缓存中，此时要淘汰 A 页面。F 是新插入的页面，访问次数为 1，排在 C 的前面。也就是说相同的访问次数，按照新旧顺序排列，淘汰掉最旧的页面。这一点是和 LRU 最大的不同的地方。

可以发现，**LFU 更新和插入新页面可以发生在链表中任意位置，删除页面都发生在表尾**。


## 解法一 Get O(1) / Put O(1)

LFU 同样要求查询尽量高效，O(1) 内查询。依旧选用 map 查询。修改和删除也需要 O(1) 完成，依旧选用双向链表，继续复用 container 包中的 list 数据结构。LFU 需要记录访问次数，所以每个结点除了存储 key，value，需要再多存储 frequency 访问次数。

还有 1 个问题需要考虑，一个是如何按频次排序？相同频次，按照先后顺序排序。如果你开始考虑排序算法的话，思考方向就偏离最佳答案了。排序至少 O(nlogn)。重新回看 LFU 的工作原理，会发现它只关心最小频次。其他频次之间的顺序并不关心。所以不需要排序。用一个 min 变量保存最小频次，淘汰时读取这个最小值能找到要删除的结点。相同频次按照先后顺序排列，这个需求还是用双向链表实现，双向链表插入的顺序体现了结点的先后顺序。相同频次对应一个双向链表，可能有多个相同频次，所以可能有多个双向链表。用一个 map 维护访问频次和双向链表的对应关系。删除最小频次时，通过 min 找到最小频次，然后再这个 map 中找到这个频次对应的双向链表，在双向链表中找到最旧的那个结点删除。这就解决了 LFU 删除操作。

LFU 的更新操作和 LRU 类似，也需要用一个 map 保存 key 和双向链表结点的映射关系。这个双向链表结点中存储的是 key-value-frequency 三个元素的元组。这样通过结点中的 key 和 frequency 可以反过来删除 map 中的 key。

定义 LFUCache 的数据结构如下：

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

LFUCache 的 Get 操作涉及更新 frequency 值和 2 个 map。在 nodes map 中通过 key 获取到结点信息。在 lists 删除结点当前 frequency 结点。删完以后 frequency ++。新的 frequency 如果在 lists 中存在，添加到双向链表表首，如果不存在，需要新建一个双向链表并把当前结点加到表首。再更新双向链表结点作为 value 的 map。最后更新 min 值，判断老的 frequency 对应的双向链表中是否已经为空，如果空了，min++。

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

LFU 的 Put 操作逻辑稍微多一点。先在 nodes map 中查询 key 是否存在，如果存在，获取这个结点，更新它的 value 值，然后手动调用一次 Get 操作，因为下面的更新逻辑和 Get 操作一致。如果 map 中不存在，接下来进行插入或者删除操作。判断 capacity 是否装满，如果装满，执行删除操作。在 min 对应的双向链表中删除表尾的结点，对应的也要删除 nodes map 中的键值。

由于新插入的页面访问次数一定为 1，所以 min 此时置为 1。新建结点，插入到 2 个 map 中。

```go

func (this *LFUCache) Put(key int, value int) {
	if this.capacity == 0 {
		return
	}
	// 如果存在，更新访问次数
	if currentValue, ok := this.nodes[key]; ok {
		currentNode := currentValue.Value.(*node)
		currentNode.value = value
		this.Get(key)
		return
	}
	// 如果不存在且缓存满了，需要删除
	if this.capacity == len(this.nodes) {
		currentList := this.lists[this.min]
		backNode := currentList.Back()
		delete(this.nodes, backNode.Value.(*node).key)
		currentList.Remove(backNode)
	}
	// 新建结点，插入到 2 个 map 中
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

总结，LFU 是由两个 map 和一个 min 指针组成的数据结构。一个 map 中 key 存的是访问次数，对应的 value 是一个个的双向链表，此处双向链表的作用是在相同频次的情况下，淘汰表尾最旧的那个页面。另一个 map 中 key 对应的 value 是双向链表的结点，结点中比 LRU 多存储了一个访问次数的值，即结点中存储 key-value-frequency 的元组。此处双向链表的作用和 LRU 是类似的，可以根据 map 中的 key 更新双向链表结点中的 value 和 frequency 的值，也可以根据双向链表结点中的 key 和 frequency 反向更新 map 中的对应关系。如下图：

![](https://img.halfrost.com/Blog/ArticleImage/146_10_0.png)

提交代码以后，成功通过所有测试用例。


![](https://img.halfrost.com/Blog/ArticleImage/146_5.png)


## 解法二 Get O(capacity) / Put O(capacity)

LFU 的另外一个思路是利用 [Index Priority Queue](https://algs4.cs.princeton.edu/24pq/) 这个数据结构。别被名字吓到，Index Priority Queue = map + Priority Queue，仅此而已。

利用 Priority Queue 维护一个最小堆，堆顶是访问次数最小的元素。map 中的 value 存储的是优先队列中结点。

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

Get 和 Put 操作要尽量的快，有 2 个问题需要解决。当访问次数相同时，如何删除掉最久的元素？当元素的访问次数发生变化时，如何快速调整堆？为了解决这 2 个问题，定义如下的数据结构：

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

堆中的结点存储这 5 个值。count 值用来决定哪个是最老的元素，类似一个操作时间戳。index 值用来 re-heapify 调整堆的。接下来实现 PriorityQueue 的方法。

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

在 Less() 方法中，frequency 从小到大排序，frequency 相同的，按 count 从小到大排序。按照优先队列建堆规则，可以得到，frequency 最小的在堆顶，相同的 frequency，count 最小的越靠近堆顶。

在 Swap() 方法中，记得要更新 index 值。在 Push() 方法中，插入时队列的长度即是该元素的 index 值，此处也要记得更新 index 值。update() 方法调用 Fix() 函数。Fix() 函数比先 Remove() 再 Push() 一个新的值，花销要小。所以此处调用 Fix() 函数，这个操作的时间复杂度是 O(log n)。

这样就维护了最小 Index Priority Queue。Get 操作非常简单：

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

在 hashmap 中查询 key，如果存在，counter 时间戳累加，调用 Priority Queue 的 update 方法，调整堆。

```go
func (this *LFUCache) Put(key int, value int) {
	if this.capacity == 0 {
		return
	}
	this.counter++
	// 如果存在，增加 frequency，再调整堆
	if item, ok := this.hash[key]; ok {
		this.pq.update(item, value, item.frequency+1, this.counter)
		return
	}
	// 如果不存在且缓存满了，需要删除。在 hashmap 和 pq 中删除。
	if len(this.pq) == this.capacity {
		item := heap.Pop(&this.pq).(*Item)
		delete(this.hash, item.key)
	}
	// 新建结点，在 hashmap 和 pq 中添加。
	item := &Item{
		value: value,
		key:   key,
		count: this.counter,
	}
	heap.Push(&this.pq, item)
	this.hash[key] = item
}
```


用最小堆实现的 LFU，Put 时间复杂度是 O(capacity)，Get 时间复杂度是 O(capacity)，不及 2 个 map 实现的版本。巧的是最小堆的版本居然打败了 100%。

![](https://img.halfrost.com/Blog/ArticleImage/146_7.png)


## 模板


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


----------------------------------------------
<div style="display: flex;justify-content: space-between;align-items: center;">
<p><a href="https://books.halfrost.com/leetcode/ChapterThree/LRUCache/">⬅️上一页</a></p>
<p><a href="https://books.halfrost.com/leetcode/ChapterFour/">下一章➡️</a></p>
</div>
