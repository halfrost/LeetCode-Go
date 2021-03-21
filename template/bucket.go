package template

import (
	"container/list"
	"sync"
)

type bucket struct {
	sync.RWMutex
	keys map[string]*list.Element
}

func (b *bucket) pairCount() int {
	b.RLock()
	defer b.RUnlock()
	return len(b.keys)
}

func (b *bucket) get(key string) *list.Element {
	b.RLock()
	defer b.RUnlock()
	if el, ok := b.keys[key]; ok {
		return el
	}
	return nil
}

func (b *bucket) set(key string, value interface{}) (*list.Element, *list.Element) {
	el := &list.Element{Value: Pair{key: key, value: value, cmd: PushFront}}
	b.Lock()
	exist := b.keys[key]
	b.keys[key] = el
	b.Unlock()
	return el, exist
}

func (b *bucket) update(key string, el *list.Element) {
	b.Lock()
	b.keys[key] = el
	b.Unlock()
}

func (b *bucket) delete(key string) *list.Element {
	b.Lock()
	el := b.keys[key]
	delete(b.keys, key)
	b.Unlock()
	return el
}

func (b *bucket) clear() {
	b.Lock()
	b.keys = make(map[string]*list.Element)
	b.Unlock()
}
