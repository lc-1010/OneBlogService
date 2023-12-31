package fifo

import (
	"container/list"

	"github.com/lc-1010/OneBlogService/cache"
)

// fifo not concurrent security
type fifo struct {
	// max cache size
	maxBytes int
	//
	onEvicted func(key string, value any)
	// used bytes
	usedBytes int
	// 双向链表 存放具体数值
	ll *list.List
	// 存储key value 值是指针
	cache map[string]*list.Element
}

type entry struct {
	key   string
	value any
}

func (e *entry) Len() int {
	return cache.CalcLen(e.value)
}

func New(maxBytes int, onEvicted func(key string, value any)) cache.Cache {

	return &fifo{
		maxBytes:  maxBytes,
		onEvicted: onEvicted,
		ll:        list.New(),
		cache:     make(map[string]*list.Element),
	}

}

func (f *fifo) Set(key string, value any) {
	if e, ok := f.cache[key]; ok {
		f.ll.MoveToBack(e)
		en := e.Value.(*entry)
		f.usedBytes = f.usedBytes - cache.CalcLen(en.value) +
			cache.CalcLen(value)
		en.value = value
		return
	}

	en := &entry{key, value}
	e := f.ll.PushBack(en)
	f.cache[key] = e
	f.usedBytes += en.Len()
	if f.maxBytes > 0 && f.usedBytes > f.maxBytes {
		f.DelOldest()
	}
}

func (f *fifo) Get(key string) any {
	if e, ok := f.cache[key]; ok {
		return e.Value.(*entry).value
	}
	return nil
}
func (f *fifo) Del(key string) {
	if e, ok := f.cache[key]; ok {
		f.removeElement(e)
	}
}
func (f *fifo) DelOldest() {
	f.removeElement(f.ll.Front())
}
func (f *fifo) Len() int {
	return f.ll.Len()
}

func (f *fifo) removeElement(e *list.Element) {
	if e == nil {
		return
	}

	f.ll.Remove(e)
	en := e.Value.(*entry)
	f.usedBytes -= en.Len()
	delete(f.cache, en.key)
	if f.onEvicted != nil {
		f.onEvicted(en.key, en.value)
	}
}
