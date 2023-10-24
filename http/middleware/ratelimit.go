package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

func RateLimiterMiddleware(tb RateLimiter) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !tb.Allow() {
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "Rate limit exceeded"})
			c.Abort()
		}
		c.Next()
	}
}

type RateLimiter interface {
	Allow() bool
}

type TokenBucketLimiter struct {
	capacity     int
	quantum      int
	tokens       int
	fillInterval time.Duration
	mux          sync.Mutex
}

func NewTokenBucketLimiter(capacity int, quantum int, fillInterval time.Duration) *TokenBucketLimiter {
	tb := &TokenBucketLimiter{
		capacity:     capacity,
		quantum:      quantum,
		tokens:       capacity,
		fillInterval: fillInterval,
	}
	go tb.fillToken()
	return tb
}

func (tb *TokenBucketLimiter) fillToken() {
	ticker := time.NewTicker(tb.fillInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			tb.mux.Lock()
			if tb.tokens < tb.capacity {
				if tb.tokens+tb.quantum > tb.capacity {
					tb.tokens = tb.capacity
				} else {
					tb.tokens += tb.quantum
				}
			}
			tb.mux.Unlock()
		}
	}
}

func (tb *TokenBucketLimiter) Allow() bool {
	tb.mux.Lock()
	defer tb.mux.Unlock()

	if tb.tokens <= 0 {
		return false
	}

	tb.tokens--
	return true
}
