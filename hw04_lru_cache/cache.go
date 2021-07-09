package hw04lrucache

import "sync"

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
	m        sync.Mutex
}

func (l *lruCache) Set(key Key, value interface{}) bool {
	l.m.Lock()
	val, ok := l.items[key]

	if ok {
		item := val.Value.(*cacheItem)
		item.value = value
		l.queue.MoveToFront(val)
	} else {
		if l.queue.Len() == l.capacity {
			l.queue.Remove(l.queue.Back())
		}

		newItem := &cacheItem{key: string(key), value: value}
		item := l.queue.PushFront(newItem)
		l.items[key] = item
	}

	l.m.Unlock()

	return ok
}

func (l *lruCache) Get(key Key) (interface{}, bool) {
	l.m.Lock()
	val, ok := l.items[key]
	l.m.Unlock()

	if !ok {
		return nil, false
	}

	return val.Value.(*cacheItem).value, true
}

func (l *lruCache) Clear() {
	l.m.Lock()

	l.queue = NewList()
	l.items = make(map[Key]*ListItem, l.capacity)

	l.m.Unlock()
}

type cacheItem struct {
	key   string
	value interface{}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
