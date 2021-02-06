package template

import (
	"container/list"
	"math/rand"
	"strconv"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_CLRUCache(t *testing.T) {
	obj := New(2)
	time.Sleep(150 * time.Millisecond)
	obj.Put("1", 1)
	time.Sleep(150 * time.Millisecond)
	obj.Put("2", 2)
	time.Sleep(150 * time.Millisecond)
	param1 := obj.Get("1")
	time.Sleep(150 * time.Millisecond)
	assert.Equal(t, 1, param1)
	obj.Put("3", 3)
	time.Sleep(150 * time.Millisecond)
	param1 = obj.Get("2")
	assert.Equal(t, nil, param1)
	obj.Put("4", 4)
	time.Sleep(150 * time.Millisecond)
	param1 = obj.Get("1")
	time.Sleep(150 * time.Millisecond)
	assert.Equal(t, nil, param1)
	param1 = obj.Get("3")
	time.Sleep(150 * time.Millisecond)
	assert.Equal(t, 3, param1)
	param1 = obj.Get("4")
	time.Sleep(150 * time.Millisecond)
	assert.Equal(t, 4, param1)
}

func MList2Ints(lru *CLRUCache) [][]interface{} {
	res := [][]interface{}{}
	for head := lru.list.Front(); head != nil; head = head.Next() {
		tmp := []interface{}{head.Value.(Pair).key, head.Value.(Pair).value}
		res = append(res, tmp)
	}
	return res
}

func BenchmarkGetAndPut1(b *testing.B) {
	b.ResetTimer()
	obj := New(128)
	wg := sync.WaitGroup{}
	wg.Add(b.N * 2)
	for i := 0; i < b.N; i++ {
		go func() {
			defer wg.Done()
			obj.Get(strconv.Itoa(rand.Intn(200)))
		}()
		go func() {
			defer wg.Done()
			obj.Put(strconv.Itoa(rand.Intn(200)), strconv.Itoa(rand.Intn(200)))
		}()
	}
	wg.Wait()
}

type Cache struct {
	sync.RWMutex
	Cap  int
	Keys map[string]*list.Element
	List *list.List
}

type pair struct {
	K, V string
}

func NewLRUCache(capacity int) Cache {
	return Cache{
		Cap:  capacity,
		Keys: make(map[string]*list.Element),
		List: list.New(),
	}
}

func (c *Cache) Get(key string) interface{} {
	c.Lock()
	if el, ok := c.Keys[key]; ok {
		c.List.MoveToFront(el)
		return el.Value.(pair).V
	}
	c.Unlock()
	return nil
}

func (c *Cache) Put(key string, value string) {
	c.Lock()
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
	c.Unlock()
}

func BenchmarkGetAndPut2(b *testing.B) {
	b.ResetTimer()
	obj := NewLRUCache(128)
	wg := sync.WaitGroup{}
	wg.Add(b.N * 2)
	for i := 0; i < b.N; i++ {
		go func() {
			defer wg.Done()
			obj.Get(strconv.Itoa(rand.Intn(200)))
		}()
		go func() {
			defer wg.Done()
			obj.Put(strconv.Itoa(rand.Intn(200)), strconv.Itoa(rand.Intn(200)))
		}()
	}
	wg.Wait()
}
