package config

import (
	"time"
)

// SecurityConfig holds security parameters to avoid account bans
type SecurityConfig struct {
	// Rate limiting
	RequestsPerSecond float64
	MinDelay          time.Duration
	MaxDelay          time.Duration
	BatchDelay        time.Duration
	
	// Retry configuration
	MaxRetries        int
	RetryBackoff      time.Duration
	
	// Request settings
	UserAgent         string
	UseRandomUserAgent bool
	
	// Message deletion settings
	MessagesPerBatch  int
	MaxMessagesTotal  int // 0 means unlimited
}

// DefaultSecurityConfig returns a safe default configuration
func DefaultSecurityConfig() *SecurityConfig {
	return &SecurityConfig{
		RequestsPerSecond:  0.5, // 1 request every 2 seconds
		MinDelay:          time.Duration(1500) * time.Millisecond,
		MaxDelay:          time.Duration(3000) * time.Millisecond,
		BatchDelay:        time.Duration(15) * time.Second,
		MaxRetries:        3,
		RetryBackoff:      time.Duration(5) * time.Second,
		UserAgent:         "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
		UseRandomUserAgent: true,
		MessagesPerBatch:  35,
		MaxMessagesTotal:  0,
	}
}

// AggressiveConfig returns a more aggressive configuration (higher risk)
func AggressiveConfig() *SecurityConfig {
	return &SecurityConfig{
		RequestsPerSecond:  1.0,
		MinDelay:          time.Duration(800) * time.Millisecond,
		MaxDelay:          time.Duration(1500) * time.Millisecond,
		BatchDelay:        time.Duration(5) * time.Second,
		MaxRetries:        3,
		RetryBackoff:      time.Duration(3) * time.Second,
		UserAgent:         "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
		UseRandomUserAgent: true,
		MessagesPerBatch:  50,
		MaxMessagesTotal:  0,
	}
}

// ConservativeConfig returns a very safe configuration (lower risk)
func ConservativeConfig() *SecurityConfig {
	return &SecurityConfig{
		RequestsPerSecond:  0.25, // 1 request every 4 seconds
		MinDelay:          time.Duration(2500) * time.Millisecond,
		MaxDelay:          time.Duration(5000) * time.Millisecond,
		BatchDelay:        time.Duration(30) * time.Second,
		MaxRetries:        5,
		RetryBackoff:      time.Duration(10) * time.Second,
		UserAgent:         "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
		UseRandomUserAgent: true,
		MessagesPerBatch:  25,
		MaxMessagesTotal:  0,
	}
}

// GetUserAgents returns a list of common user agents for rotation
func GetUserAgents() []string {
	return []string{
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:121.0) Gecko/20100101 Firefox/121.0",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.1 Safari/605.1.15",
		"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36",
		"Mozilla/5.0 (X11; Linux x86_64; rv:121.0) Gecko/20100101 Firefox/121.0",
	}
}
