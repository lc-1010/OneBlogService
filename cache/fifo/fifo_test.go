package fifo_test

import (
	"testing"

	"github.com/lc-1010/OneBlogService/cache/fifo"
	"github.com/matryer/is"
)

func TestGet(t *testing.T) {
	is := is.New(t)
	cache := fifo.New(24, nil)
	cache.DelOldest()
	cache.Set("k1", 1)
	v := cache.Get("k1")
	is.Equal(v, 1)

	cache.Del("k1")
	is.Equal(0, cache.Len())
}

func TestOnEvicted(t *testing.T) {
	is := is.New(t)
	keys := make([]string, 0, 8)
	// 自动删除
	onEvicted := func(key string, value any) {
		keys = append(keys, key)
	}
	//空间小一点
	cache := fifo.New(8, onEvicted)

	cache.Set("k1", 1)
	cache.Set("k2", 2)
	cache.Get("k1")
	cache.Set("k3", 3)
	//超过容量会执行删除操作并调用onEvicted
	cache.Get("k1")
	cache.Set("k4", 4)
	expected := []string{"k1", "k2"}
	is.Equal(expected, keys)
	is.Equal(2, cache.Len())
}
