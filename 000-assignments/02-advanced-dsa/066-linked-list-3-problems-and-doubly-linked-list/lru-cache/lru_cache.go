package lrucache

import "container/list"

type LRUCache struct {
	capacity int
	l        list.List
	m        map[int]*list.Element
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		l:        list.List{},
		m:        make(map[int]*list.Element),
	}
}

func (this *LRUCache) Get(key int) int {
	if v, ok := this.m[key]; ok {
		this.l.MoveToFront(v)
		return v.Value.(pair).second
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	if v, ok := this.m[key]; ok {
		v.Value = pair{key, value}
		this.l.MoveToFront(v)
		return
	}

	if this.l.Len() == this.capacity {
		d_key := this.l.Back().Value.(pair).first
		delete(this.m, d_key)
		this.l.Remove(this.l.Back())
	}

	this.l.PushFront(pair{key, value})
	this.m[key] = this.l.Front()
}

type pair struct {
	first  int
	second int
}
