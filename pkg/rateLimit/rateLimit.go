package ratelimiting

import (
	"sync"
	"time"
)

type TokenBucket struct {
	mu            sync.Mutex
	buckets       map[string]*tokenBucket
	rate          float64 //tokens per second
	capacity      int     // max tokens
	lastCleanup   time.Time
	cleanupPeriod time.Duration
}

type tokenBucket struct {
	tokens     float64
	lastRefill time.Time
}

func NewTokenBucket(rate float64,capacity int)*TokenBucket{
	return &TokenBucket{
		buckets: make(map[string]*tokenBucket),
		rate: rate,
		capacity: capacity,
		lastCleanup: time.Now(),
		cleanupPeriod: 5 *time.Minute, // Cleanup inactive buckets every 5 minutes
	}
}

func(tb *TokenBucket)AllowRequest(email string)bool{
	tb.mu.Lock()
	defer tb.mu.Unlock()

	// Periodic cleanup of inactive buckets
	if time.Since(tb.lastCleanup) > tb.cleanupPeriod{
		tb.cleanup()
		tb.lastCleanup=time.Now()
	}

	bucket,ok:=tb.buckets[email]
	if !ok{
		bucket=&tokenBucket{
			tokens: float64(tb.capacity),
			lastRefill: time.Now(),
		}
		tb.buckets[email]=bucket
	}

	now:=time.Now()
	elapsed:=now.Sub(bucket.lastRefill).Seconds()
	refill:=elapsed*tb.rate
	bucket.tokens=min(float64(tb.capacity),bucket.tokens+refill)
	bucket.lastRefill=now

	if bucket.tokens>=1{
		bucket.tokens--
		return true
	}
	return false

}

func(tb *TokenBucket) cleanup(){
	now:=time.Now()

	for email,bucket:=range tb.buckets{
		if now.Sub(bucket.lastRefill)>tb.cleanupPeriod{
			delete(tb.buckets,email)
		}
	}
}

func min(a,b float64)float64{
	if a<b{
		return a
	}
	return b
}
