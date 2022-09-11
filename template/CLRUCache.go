package template

import (
	"container/list"
	"hash/fnv"
	"sync"
)

type command int

const (
	// MoveToFront define
	MoveToFront command = iota
	// PushFront define
	PushFront
	// Delete define
	Delete
)

type clear struct {
	done chan struct{}
}

// CLRUCache define: High Concurrency LRUCache
type CLRUCache struct {
	sync.RWMutex
	cap         int
	list        *list.List
	buckets     []*bucket
	bucketMask  uint32
	deletePairs chan *list.Element
	movePairs   chan *list.Element
	control     chan interface{}
}

// Pair define
type Pair struct {
	key   string
	value interface{}
	cmd   command
}

// New define
func New(capacity int) *CLRUCache {
	c := &CLRUCache{
		cap:        capacity,
		list:       list.New(),
		bucketMask: uint32(1024) - 1,
		buckets:    make([]*bucket, 1024),
	}
	for i := 0; i < 1024; i++ {
		c.buckets[i] = &bucket{
			keys: make(map[string]*list.Element),
		}
	}
	c.restart()
	return c
}

// Get define
func (c *CLRUCache) Get(key string) interface{} {
	el := c.bucket(key).get(key)
	if el == nil {
		return nil
	}
	c.move(el)
	return el.Value.(Pair).value
}

// Put define
func (c *CLRUCache) Put(key string, value interface{}) {
	el, exist := c.bucket(key).set(key, value)
	if exist != nil {
		c.deletePairs <- exist
	}
	c.move(el)
}

func (c *CLRUCache) move(el *list.Element) {
	select {
	case c.movePairs <- el:
	default:
	}
}

// Delete define
func (c *CLRUCache) Delete(key string) bool {
	el := c.bucket(key).delete(key)
	if el != nil {
		c.deletePairs <- el
		return true
	}
	return false
}

// Clear define
func (c *CLRUCache) Clear() {
	done := make(chan struct{})
	c.control <- clear{done: done}
	<-done
}

// Count define
func (c *CLRUCache) Count() int {
	count := 0
	for _, b := range c.buckets {
		count += b.pairCount()
	}
	return count
}

func (c *CLRUCache) bucket(key string) *bucket {
	h := fnv.New32a()
	h.Write([]byte(key))
	return c.buckets[h.Sum32()&c.bucketMask]
}

func (c *CLRUCache) stop() {
	close(c.movePairs)
	<-c.control
}

func (c *CLRUCache) restart() {
	c.deletePairs = make(chan *list.Element, 128)
	c.movePairs = make(chan *list.Element, 128)
	c.control = make(chan interface{})
	go c.worker()
}

func (c *CLRUCache) worker() {
	defer close(c.control)
	for {
		select {
		case el, ok := <-c.movePairs:
			if !ok {
				goto clean
			}
			if c.doMove(el) && c.list.Len() > c.cap {
				el := c.list.Back()
				c.list.Remove(el)
				c.bucket(el.Value.(Pair).key).delete(el.Value.(Pair).key)
			}
		case el := <-c.deletePairs:
			c.list.Remove(el)
		case control := <-c.control:
			switch msg := control.(type) {
			case clear:
				for _, bucket := range c.buckets {
					bucket.clear()
				}
				c.list = list.New()
				msg.done <- struct{}{}
			}
		}
	}
clean:
	for {
		select {
		case el := <-c.deletePairs:
			c.list.Remove(el)
		default:
			close(c.deletePairs)
			return
		}
	}
}

func (c *CLRUCache) doMove(el *list.Element) bool {
	if el.Value.(Pair).cmd == MoveToFront {
		c.list.MoveToFront(el)
		return false
	}
	newel := c.list.PushFront(el.Value.(Pair))
	c.bucket(el.Value.(Pair).key).update(el.Value.(Pair).key, newel)
	return true
}
