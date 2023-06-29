package limiter

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
)

// LimiterInterface Limiter interface
// Key：获取对应的限流器的键值对名称。
// Key returns the key for the given gin context.
// GetBucket：获取令牌桶。
// AddBuckets：新增多个令牌桶
type LimiterInterface interface {
	Key(c *gin.Context) string //
	GetBucket(key string) (*ratelimit.Bucket, bool)
	AddBucket(rules ...LimiterBucketRule) LimiterInterface
}

// Limiter bucket of ratelimit
type Limiter struct {
	LimiterBuckets map[string]*ratelimit.Bucket
}

// LimiterBucketRule Limiter bucket rule
// capacity is the number of requests in the bucket
// fill interval is the interval of the bucket
// Key：自定义键值对名称。
// FillInterval：间隔多久时间放 N 个令牌。
// Capacity：令牌桶的容量。
// Quantum：每次到达间隔时间后所放的具体令牌数量。
type LimiterBucketRule struct {
	Key          string
	FillInterval time.Duration
	Capacity     int64
	Quantum      int64
}
