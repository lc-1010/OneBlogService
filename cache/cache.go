package cache

type Cache interface {
	Set(key string, value any)
	Get(key string) any
	Del(key string)
	DelOldest()
	Len() int
}

const DefaultMaxBytes = 1 << 29

type safeCache struct{}
