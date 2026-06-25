---
title: 3.3 LRUCache
type: docs
weight: 3
---

# Least Recently Used LRUCache

![](https://img.halfrost.com/Blog/ArticleImage/146_1_.png)

LRU is the abbreviation for Least Recently Used. It is a commonly used page replacement algorithm that selects the page that has not been used for the longest time recently for eviction. As shown above, when inserting F, one of the original pages needs to be evicted.

![](https://img.halfrost.com/Blog/ArticleImage/146_2_0.png)

According to the LRU strategy, the page that has not been used for the longest time recently is evicted each time, so page A is evicted first. When inserting C again, we find that page C is already in the cache. At this time, page C needs to be moved to the front because it has been used. By the same logic, when inserting page G, page G is a new page and is not in the cache, so page B is evicted. When inserting page H, page H is a new page and is not in the cache, so page D is evicted. When inserting E, we find that page E is already in the cache. At this time, page E needs to be moved to the front. When inserting page I, page I is a new page and is not in the cache, so page F is evicted.

It can be found that **LRU updates and inserting new pages both happen at the head of the linked list, while deleting pages happens at the tail of the linked list**.

## Solution 1 Get O(1) / Put O(1)

LRU requires queries to be as efficient as possible, querying within O(1). So a map is definitely chosen for querying. Modification and deletion should also be completed in O(1) as much as possible. Looking through common data structures: linked lists, stacks, queues, trees, and graphs. Trees and graphs are ruled out; stacks and queues cannot arbitrarily query elements in the middle, so they are also ruled out. Therefore, a linked list is chosen for implementation. However, if a singly linked list is used, deleting this node requires O(n) traversal to find the predecessor node. Therefore, a doubly linked list is chosen, so deletion can also be completed in O(1).

Since the underlying implementation of list in Go's container package is a doubly linked list, this data structure can be reused directly. The data structure of LRUCache is defined as follows:

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

Here two questions need to be explained: what is stored as the value in list? What is the use of the pair struct?

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

In container/list, the type of each node in this doubly linked list is Element. Element stores 4 values: the predecessor and successor nodes, the head node of the doubly linked list, and the value. The value here is of interface type. The author stores the pair struct in this value. This explains what data is stored in the list.

Why store pair? Wouldn't it be enough to store only v? Why also store a copy of the key? The reason is that when LRUCache performs a delete operation, it needs to maintain 2 data structures: one is the map, and the other is the doubly linked list. Delete the evicted value from the doubly linked list, and delete the key corresponding to the evicted value from the map. If the key is not stored in the value of the doubly linked list, then deleting the key from the map is a bit troublesome. If you insist on implementing it, you need to first get the address of this Element node in the doubly linked list. Then traverse the map, find the key corresponding to the address of this Element element in the map, and then delete it. The time complexity of doing this is O(n), so it cannot be O(1). Therefore, the Value in the doubly linked list needs to store this pair.

The Get operation of LRUCache is very simple: directly read the node of the doubly linked list from the map. If it exists in the map, move it to the head of the doubly linked list and return its value; if it does not exist in the map, return -1.

```go 
func (c *LRUCache) Get(key int) int {
	if el, ok := c.Keys[key]; ok {
		c.List.MoveToFront(el)
		return el.Value.(pair).V
	}
	return -1
}
```

The Put operation of LRUCache is also not difficult. First query whether the key exists in the map. If it exists, update its value and move the node to the head of the doubly linked list. If it does not exist in the map, create this new node and add it to the doubly linked list and the map. Finally, don't forget that the cap of the doubly linked list also needs to be maintained. If it exceeds cap, evict the last node, delete this node from the doubly linked list, and delete the key corresponding to this node from the map.

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

In summary, LRU is a data structure composed of a map and a doubly linked list. The value corresponding to the key in the map is the node of the doubly linked list. The doubly linked list stores key-value pairs. The head of the doubly linked list updates the cache, and the tail evicts the cache. As shown below:

![](https://img.halfrost.com/Blog/ArticleImage/146_9.png)

After submitting the code, it successfully passed all test cases.

![](https://img.halfrost.com/Blog/ArticleImage/146_4_.png)


## Solution 2 Get O(1) / Put O(1)

In terms of data structures, I couldn't think of any other solution, but from the defeated percentage, it seemed there was still room for constant-factor optimization. After repeated consideration, the author felt that the place that might cause the runtime to become longer was interface{} type assertion; there was no room for optimization elsewhere. I tried submitting a handwritten doubly linked list. The code is as follows:

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

After submitting, it really reached 100%.

![](https://img.halfrost.com/Blog/ArticleImage/146_6.png)

The LRU implemented by the above code is essentially not optimized; it is just written in another way, without using the container package.


## Template

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
