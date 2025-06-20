package main

import "container/list"

type LRUCache struct {
	capacity int
	data     map[int]*list.Element
	cache    *list.List
}

type kv struct {
	k, v int
}

func ConstructorXI(capacity int) LRUCache {
	m := make(map[int]*list.Element)
	return LRUCache{capacity, m, list.New()}
}

func (this *LRUCache) Get(key int) int {
	if val, ok := this.data[key]; ok {
		this.cache.MoveToFront(val)
		return val.Value.(kv).v

	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if val, ok := this.data[key]; ok {
		val.Value = kv{key, value}
		this.cache.MoveToFront(val)
		return
	}
	if this.cache.Len() == this.capacity {
		last := this.cache.Back()
		delete(this.data, last.Value.(kv).k)
		this.cache.Remove(last)
	}
	this.data[key] = this.cache.PushFront(kv{key, value})
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
