package template

// LRUCache define
type LRUCache struct {
	head, tail *Node
	keys       map[int]*Node
	capacity   int
}

// Node define
type Node struct {
	key, val   int
	prev, next *Node
}

// ConstructorLRU define
func ConstructorLRU(capacity int) LRUCache {
	return LRUCache{keys: make(map[int]*Node), capacity: capacity}
}

// Get define
func (lruCache *LRUCache) Get(key int) int {
	if node, ok := lruCache.keys[key]; ok {
		lruCache.Remove(node)
		lruCache.Add(node)
		return node.val
	}
	return -1
}

// Put define
func (lruCache *LRUCache) Put(key int, value int) {
	node, ok := lruCache.keys[key]
	if ok {
		node.val = value
		lruCache.Remove(node)
		lruCache.Add(node)
		return
	}
	node = &Node{key: key, val: value}
	lruCache.keys[key] = node
	lruCache.Add(node)
	if len(lruCache.keys) > lruCache.capacity {
		delete(lruCache.keys, lruCache.tail.key)
		lruCache.Remove(lruCache.tail)
	}
}

// Add define
func (lruCache *LRUCache) Add(node *Node) {
	node.prev = nil
	node.next = lruCache.head
	if lruCache.head != nil {
		lruCache.head.prev = node
	}
	lruCache.head = node
	if lruCache.tail == nil {
		lruCache.tail = node
		lruCache.tail.next = nil
	}
}

// Remove define
func (lruCache *LRUCache) Remove(node *Node) {
	if node == lruCache.head {
		lruCache.head = node.next
		if node.next != nil {
			node.next.prev = nil
		}
		node.next = nil
		return
	}
	if node == lruCache.tail {
		lruCache.tail = node.prev
		node.prev.next = nil
		node.prev = nil
		return
	}
	node.prev.next = node.next
	node.next.prev = node.prev
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

// 22%
// import "container/list"

// type LRUCache struct {
//     Cap  int
//     Keys map[int]*list.Element
//     List *list.List
// }

// type pair struct {
//     K, V int
// }

// func Constructor(capacity int) LRUCache {
//     return LRUCache{
//         Cap: capacity,
//         Keys: make(map[int]*list.Element),
//         List: list.New(),
//     }
// }

// func (c *LRUCache) Get(key int) int {
//     if el, ok := c.Keys[key]; ok {
//         c.List.MoveToFront(el)
//         return el.Value.(pair).V
//     }
//     return -1
// }

// func (c *LRUCache) Put(key int, value int) {
// 	if el, ok := c.Keys[key]; ok {
// 		el.Value = pair{K: key, V: value}
// 		c.List.MoveToFront(el)
// 	} else {
// 		el := c.List.PushFront(pair{K: key, V: value})
// 		c.Keys[key] = el
// 	}
// 	if c.List.Len() > c.Cap {
// 		el := c.List.Back()
// 		c.List.Remove(el)
// 		delete(c.Keys, el.Value.(pair).K)
// 	}
// }

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
