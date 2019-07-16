package leetcode

const Len int = 100000

type MyHashMap struct {
	content [Len]*HashNode
}

type HashNode struct {
	key  int
	val  int
	next *HashNode
}

func (N *HashNode) Put(key int, value int) {
	if N.key == key {
		N.val = value
		return
	}
	if N.next == nil {
		N.next = &HashNode{key, value, nil}
		return
	}
	N.next.Put(key, value)
}

func (N *HashNode) Get(key int) int {
	if N.key == key {
		return N.val
	}
	if N.next == nil {
		return -1
	}
	return N.next.Get(key)
}

func (N *HashNode) Remove(key int) *HashNode {
	if N.key == key {
		p := N.next
		N.next = nil
		return p
	}
	if N.next != nil {
		return N.next.Remove(key)
	}
	return nil
}

/** Initialize your data structure here. */
func Constructor706() MyHashMap {
	return MyHashMap{}
}

/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int) {
	node := this.content[this.Hash(key)]
	if node == nil {
		this.content[this.Hash(key)] = &HashNode{key: key, val: value, next: nil}
		return
	}
	node.Put(key, value)
}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	HashNode := this.content[this.Hash(key)]
	if HashNode == nil {
		return -1
	}
	return HashNode.Get(key)
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int) {
	HashNode := this.content[this.Hash(key)]
	if HashNode == nil {
		return
	}
	this.content[this.Hash(key)] = HashNode.Remove(key)
}

func (this *MyHashMap) Hash(value int) int {
	return value % Len
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
