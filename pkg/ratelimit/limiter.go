package ratelimit

import (
	"context"
	"math/rand"
	"time"

	"golang.org/x/time/rate"
)

func init() {
	// Seed the random number generator for unpredictable delays
	rand.Seed(time.Now().UnixNano())
}

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
	l.limiter.Wait(context.Background())
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
	if deltaMs <= 0 {
		return l.minDelay
	}
	randomMs := rand.Int63n(deltaMs)
	return l.minDelay + time.Duration(randomMs)*time.Millisecond
}

// BatchDelay waits for the configured batch delay
func (l *Limiter) BatchDelay() {
	time.Sleep(l.batchDelay)
}

// ExponentialBackoff performs exponential backoff with the given attempt number
func (l *Limiter) ExponentialBackoff(attempt int) time.Duration {
	// Cap attempt number to prevent overflow
	if attempt > 20 {
		attempt = 20
	}
	
	backoff := l.minDelay * time.Duration(1<<uint(attempt))
	if backoff > l.maxBackoff {
		backoff = l.maxBackoff
	}
	
	// Add jitter (up to 25% of backoff or at least 1ms)
	jitterMax := backoff / 4
	if jitterMax < time.Millisecond {
		jitterMax = time.Millisecond
	}
	jitter := time.Duration(rand.Int63n(int64(jitterMax)))
	
	return backoff + jitter
}

// WaitWithBackoff waits with exponential backoff for the given attempt
func (l *Limiter) WaitWithBackoff(attempt int) {
	backoff := l.ExponentialBackoff(attempt)
	time.Sleep(backoff)
}
