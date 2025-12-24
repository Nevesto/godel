package ratelimit

import (
	"math/rand"
	"time"

	"golang.org/x/time/rate"
)

// Limiter provides rate limiting and backoff functionality
type Limiter struct {
	limiter         *rate.Limiter
	minDelay        time.Duration
	maxDelay        time.Duration
	batchDelay      time.Duration
	useRandomDelay  bool
	backoffMultiplier float64
	maxBackoff      time.Duration
}

// NewLimiter creates a new rate limiter with configurable parameters
func NewLimiter(requestsPerSecond float64, minDelay, maxDelay, batchDelay time.Duration) *Limiter {
	return &Limiter{
		limiter:         rate.NewLimiter(rate.Limit(requestsPerSecond), 1),
		minDelay:        minDelay,
		maxDelay:        maxDelay,
		batchDelay:      batchDelay,
		useRandomDelay:  true,
		backoffMultiplier: 2.0,
		maxBackoff:      5 * time.Minute,
	}
}

// Wait blocks until the rate limiter allows the next operation
func (l *Limiter) Wait() {
	l.limiter.Wait(nil)
	if l.useRandomDelay {
		delay := l.RandomDelay()
		time.Sleep(delay)
	} else {
		time.Sleep(l.minDelay)
	}
}

// RandomDelay returns a random delay between min and max delay
func (l *Limiter) RandomDelay() time.Duration {
	if l.minDelay >= l.maxDelay {
		return l.minDelay
	}
	deltaMs := l.maxDelay.Milliseconds() - l.minDelay.Milliseconds()
	randomMs := rand.Int63n(deltaMs)
	return l.minDelay + time.Duration(randomMs)*time.Millisecond
}

// BatchDelay waits for the configured batch delay
func (l *Limiter) BatchDelay() {
	time.Sleep(l.batchDelay)
}

// ExponentialBackoff performs exponential backoff with the given attempt number
func (l *Limiter) ExponentialBackoff(attempt int) time.Duration {
	backoff := l.minDelay * time.Duration(1<<uint(attempt))
	if backoff > l.maxBackoff {
		backoff = l.maxBackoff
	}
	// Add jitter
	jitter := time.Duration(rand.Int63n(int64(backoff / 4)))
	return backoff + jitter
}

// WaitWithBackoff waits with exponential backoff for the given attempt
func (l *Limiter) WaitWithBackoff(attempt int) {
	backoff := l.ExponentialBackoff(attempt)
	time.Sleep(backoff)
}
