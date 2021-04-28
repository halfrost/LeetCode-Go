---
title: 3.3 LRUCache
type: docs
weight: 3
---

# 最近最少使用 LRUCache

![](https://img.halfrost.com/Blog/ArticleImage/146_1_.png)

LRU 是 Least Recently Used 的缩写，即最近最少使用，是一种常用的页面置换算法，选择最近最久未使用的页面予以淘汰。如上图，要插入 F 的时候，此时需要淘汰掉原来的一个页面。

![](https://img.halfrost.com/Blog/ArticleImage/146_2_0.png)

根据 LRU 的策略，每次都淘汰最近最久未使用的页面，所以先淘汰 A 页面。再插入 C 的时候，发现缓存中有 C 页面，这个时候需要把 C 页面放到首位，因为它被使用了。以此类推，插入 G 页面，G 页面是新页面，不在缓存中，所以淘汰掉 B 页面。插入 H 页面，H 页面是新页面，不在缓存中，所以淘汰掉 D 页面。插入 E 的时候，发现缓存中有 E 页面，这个时候需要把 E 页面放到首位。插入 I 页面，I 页面是新页面，不在缓存中，所以淘汰掉 F 页面。

可以发现，**LRU 更新和插入新页面都发生在链表首，删除页面都发生在链表尾**。

## 解法一 Get O(1) / Put O(1)

LRU 要求查询尽量高效，O(1) 内查询。那肯定选用 map 查询。修改，删除也要尽量 O(1) 完成。搜寻常见的数据结构，链表，栈，队列，树，图。树和图排除，栈和队列无法任意查询中间的元素，也排除。所以选用链表来实现。但是如果选用单链表，删除这个结点，需要 O(n) 遍历一遍找到前驱结点。所以选用双向链表，在删除的时候也能 O(1) 完成。

由于 Go 的 container 包中的 list 底层实现是双向链表，所以可以直接复用这个数据结构。定义 LRUCache 的数据结构如下：

```go
import "container/list"

type LRUCache struct {
    Cap  int
    Keys map[int]*list.Element
    List *list.List
}

type pair struct {
    K, V int
}

func Constructor(capacity int) LRUCache {
    return LRUCache{
        Cap: capacity,
        Keys: make(map[int]*list.Element),
        List: list.New(),
    }
}

```

这里需要解释 2 个问题，list 中的值存的是什么？pair 这个结构体有什么用？

```go
type Element struct {
	// Next and previous pointers in the doubly-linked list of elements.
	// To simplify the implementation, internally a list l is implemented
	// as a ring, such that &l.root is both the next element of the last
	// list element (l.Back()) and the previous element of the first list
	// element (l.Front()).
	next, prev *Element

	// The list to which this element belongs.
	list *List

	// The value stored with this element.
	Value interface{}
}
```

在 container/list 中，这个双向链表的每个结点的类型是 Element。Element 中存了 4 个值，前驱和后继结点，双向链表的头结点，value 值。这里的 value 是 interface 类型。笔者在这个 value 里面存了 pair 这个结构体。这就解释了 list 里面存的是什么数据。

为什么要存 pair 呢？单单只存 v 不行么，为什么还要存一份 key ？原因是在 LRUCache 执行删除操作的时候，需要维护 2 个数据结构，一个是 map，一个是双向链表。在双向链表中删除淘汰出去的 value，在 map 中删除淘汰出去 value 对应的 key。如果在双向链表的 value 中不存储 key，那么再删除 map 中的 key 的时候有点麻烦。如果硬要实现，需要先获取到双向链表这个结点 Element 的地址。然后遍历 map，在 map 中找到存有这个 Element 元素地址对应的 key，再删除。这样做时间复杂度是 O(n)，做不到 O(1)。所以双向链表中的 Value 需要存储这个 pair。

LRUCache 的 Get 操作很简单，在 map 中直接读取双向链表的结点。如果 map 中存在，将它移动到双向链表的表头，并返回它的 value 值，如果 map 中不存在，返回 -1。

```go 
func (c *LRUCache) Get(key int) int {
	if el, ok := c.Keys[key]; ok {
		c.List.MoveToFront(el)
		return el.Value.(pair).V
	}
	return -1
}
```

LRUCache 的 Put 操作也不难。先查询 map 中是否存在 key，如果存在，更新它的 value，并且把该结点移到双向链表的表头。如果 map 中不存在，新建这个结点加入到双向链表和 map 中。最后别忘记还需要维护双向链表的 cap，如果超过 cap，需要淘汰最后一个结点，双向链表中删除这个结点，map 中删掉这个结点对应的 key。

```go
func (c *LRUCache) Put(key int, value int) {
	if el, ok := c.Keys[key]; ok {
		el.Value = pair{K: key, V: value}
		c.List.MoveToFront(el)
	} else {
		el := c.List.PushFront(pair{K: key, V: value})
		c.Keys[key] = el
	}
	if c.List.Len() > c.Cap {
		el := c.List.Back()
		c.List.Remove(el)
		delete(c.Keys, el.Value.(pair).K)
	}
}

```

总结，LRU 是由一个 map 和一个双向链表组成的数据结构。map 中 key 对应的 value 是双向链表的结点。双向链表中存储 key-value 的 pair。双向链表表首更新缓存，表尾淘汰缓存。如下图：

![](https://img.halfrost.com/Blog/ArticleImage/146_9.png)

提交代码以后，成功通过所有测试用例。

![](https://img.halfrost.com/Blog/ArticleImage/146_4_.png)


## 解法二 Get O(1) / Put O(1)

数据结构上想不到其他解法了，但从打败的百分比上，看似还有常数的优化空间。笔者反复思考，觉得可能导致运行时间变长的地方是在 interface{} 类型推断，其他地方已无优化的空间。手写一个双向链表提交试试，代码如下：

```go

type LRUCache struct {
	head, tail *Node
	keys       map[int]*Node
	capacity   int
}

type Node struct {
	key, val   int
	prev, next *Node
}

func ConstructorLRU(capacity int) LRUCache {
	return LRUCache{keys: make(map[int]*Node), capacity: capacity}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.keys[key]; ok {
		this.Remove(node)
		this.Add(node)
		return node.val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.keys[key]; ok {
		node.val = value
		this.Remove(node)
		this.Add(node)
		return
	} else {
		node = &Node{key: key, val: value}
		this.keys[key] = node
		this.Add(node)
	}
	if len(this.keys) > this.capacity {
		delete(this.keys, this.tail.key)
		this.Remove(this.tail)
	}
}

func (this *LRUCache) Add(node *Node) {
	node.prev = nil
	node.next = this.head
	if this.head != nil {
		this.head.prev = node
	}
	this.head = node
	if this.tail == nil {
		this.tail = node
		this.tail.next = nil
	}
}

func (this *LRUCache) Remove(node *Node) {
	if node == this.head {
		this.head = node.next
		if node.next != nil {
			node.next.prev = nil
		}
		node.next = nil
		return
	}
	if node == this.tail {
		this.tail = node.prev
		node.prev.next = nil
		node.prev = nil
		return
	}
	node.prev.next = node.next
	node.next.prev = node.prev
}

```

提交以后还真的 100% 了。

![](https://img.halfrost.com/Blog/ArticleImage/146_6.png)

上述代码实现的 LRU 本质并没有优化，只是换了一个写法，没有用 container 包而已。


## 模板

```go
type LRUCache struct {
	head, tail *Node
	Keys       map[int]*Node
	Cap        int
}

type Node struct {
	Key, Val   int
	Prev, Next *Node
}

func Constructor(capacity int) LRUCache {
	return LRUCache{Keys: make(map[int]*Node), Cap: capacity}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.Keys[key]; ok {
		this.Remove(node)
		this.Add(node)
		return node.Val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.Keys[key]; ok {
		node.Val = value
		this.Remove(node)
		this.Add(node)
		return
	} else {
		node = &Node{Key: key, Val: value}
		this.Keys[key] = node
		this.Add(node)
	}
	if len(this.Keys) > this.Cap {
		delete(this.Keys, this.tail.Key)
		this.Remove(this.tail)
	}
}

func (this *LRUCache) Add(node *Node) {
	node.Prev = nil
	node.Next = this.head
	if this.head != nil {
		this.head.Prev = node
	}
	this.head = node
	if this.tail == nil {
		this.tail = node
		this.tail.Next = nil
	}
}

func (this *LRUCache) Remove(node *Node) {
	if node == this.head {
		this.head = node.Next
		node.Next = nil
		return
	}
	if node == this.tail {
		this.tail = node.Prev
		node.Prev.Next = nil
		node.Prev = nil
		return
	}
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

```


----------------------------------------------
<div style="display: flex;justify-content: space-between;align-items: center;">
<p><a href="https://books.halfrost.com/leetcode/ChapterThree/UnionFind/">⬅️上一页</a></p>
<p><a href="https://books.halfrost.com/leetcode/ChapterThree/LFUCache/">下一页➡️</a></p>
</div>
