package algos

import "container/list"

type LRUCache struct {
	capacity int
	cache    map[int]*list.Element
	ll       *list.List
}

func InitCache(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*list.Element),
		ll:       list.New(),
	}
}

func (c *LRUCache) Get(key int) int {
	if elem, ok := c.cache[key]; ok {
		c.ll.MoveToFront(elem)
		return elem.Value.(int)
	}
	return -1
}

func (c *LRUCache) Put(key int, value int) {
	if elem, ok := c.cache[key]; ok {
		elem.Value = value
		c.ll.MoveToFront(elem)
	} else {
		if c.ll.Len() == c.capacity {
			tail := c.ll.Back()
			delete(c.cache, tail.Value.(int))
			c.ll.Remove(tail)
		}

		elem := c.ll.PushFront(value)
		c.cache[key] = elem
	}
}

func (c *LRUCache) Size() int {
	if nil != c.ll {
		return c.ll.Len()
	}
	return 0
}
