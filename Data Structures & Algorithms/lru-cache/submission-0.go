type LRUCache struct {
    c int
	m map[int]*list.Element
	l *list.List
}

type kv struct {
	k int
	v int
}

func Constructor(capacity int) LRUCache {
    return LRUCache{
		c: capacity,
		m: make(map[int]*list.Element, capacity),
		l: list.New(),
	}
}

func (this *LRUCache) Get(key int) int {
	el, exists := this.m[key]
	if !exists {
		return -1
	}

	this.l.MoveToFront(el)

    return el.Value.(kv).v
}

func (this *LRUCache) Put(key int, value int) {
    el, exists := this.m[key]
	if exists {
		el.Value = kv{k: key, v: value}
		this.l.MoveToFront(el)
		return
	}

	if this.l.Len() == this.c {
		last := this.l.Back()
		this.l.Remove(last)
		delete(this.m, last.Value.(kv).k)
	}

	el = this.l.PushFront(kv{k: key, v: value})
	this.m[key] = el
}
