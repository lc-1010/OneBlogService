package limiter

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
)

// MethodLimiter for router limiter
// implement LimiterInterface
type MethodLimiter struct {
	*Limiter
}

func NewMethodLimiter() LimiterInterface {
	l := &Limiter{
		LimiterBuckets: make(map[string]*ratelimit.Bucket),
	}
	return MethodLimiter{
		Limiter: l,
	}
}

// Key use ? to split uri get key
// key is the uri request
func (l MethodLimiter) Key(c *gin.Context) string {
	uri := c.Request.RequestURI
	index := strings.Index(uri, "?")
	if index == -1 {
		return uri
	}
	return uri[:index]
}

// GetBucket get bucket
func (l MethodLimiter) GetBucket(key string) (*ratelimit.Bucket, bool) {
	bucket, ok := l.LimiterBuckets[key]
	return bucket, ok
}

// AddBucket add bucket
func (l MethodLimiter) AddBucket(rules ...LimiterBucketRule) LimiterInterface {
	for _, rule := range rules {
		if _, ok := l.LimiterBuckets[rule.Key]; !ok {
			bucket := ratelimit.NewBucketWithQuantum(
				rule.FillInterval,
				rule.Capacity,
				rule.Quantum,
			)
			l.LimiterBuckets[rule.Key] = bucket
		}
	}
	return l
}
