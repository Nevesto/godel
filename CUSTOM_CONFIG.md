# Example Custom Configuration

This document shows how to create custom security configurations if the built-in profiles don't meet your needs.

## Configuration Structure

Security configurations are defined in `pkg/config/security.go`. Each configuration has these parameters:

### Rate Limiting Parameters

```go
RequestsPerSecond float64  // How many requests per second (0.25 = 1 request every 4 seconds)
```

### Delay Parameters

```go
MinDelay      time.Duration  // Minimum delay between operations
MaxDelay      time.Duration  // Maximum delay between operations (randomized between min and max)
BatchDelay    time.Duration  // Delay between batches of messages
```

### Retry Parameters

```go
MaxRetries   int            // Maximum number of retry attempts on failure
RetryBackoff time.Duration  // Initial backoff time for retries (increases exponentially)
```

### Request Parameters

```go
UserAgent         string  // User agent string to use in requests
UseRandomUserAgent bool    // Whether to rotate user agents
```

### Message Processing Parameters

```go
MessagesPerBatch int  // Number of messages to process per batch
MaxMessagesTotal int  // Maximum total messages to delete (0 = unlimited)
```

## Creating a Custom Profile

To create a custom profile, edit `pkg/config/security.go` and add a new function:

```go
// UltraConservativeConfig returns an extremely safe configuration
func UltraConservativeConfig() *SecurityConfig {
	return &SecurityConfig{
		RequestsPerSecond:  0.1,  // 1 request every 10 seconds
		MinDelay:          time.Duration(5000) * time.Millisecond,
		MaxDelay:          time.Duration(10000) * time.Millisecond,
		BatchDelay:        time.Duration(60) * time.Second,
		MaxRetries:        5,
		RetryBackoff:      time.Duration(15) * time.Second,
		UserAgent:         "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
		UseRandomUserAgent: true,
		MessagesPerBatch:  20,
		MaxMessagesTotal:  0,
	}
}
```

Then update `cmd/security_helper.go` to include your profile:

```go
func GetSecurityConfig() *config.SecurityConfig {
	switch securityProfile {
	case "conservative":
		return config.ConservativeConfig()
	case "aggressive":
		return config.AggressiveConfig()
	case "ultra-conservative":
		return config.UltraConservativeConfig()
	default:
		return config.DefaultSecurityConfig()
	}
}
```

## Example Custom Configurations

### Ultra-Safe (Paranoid Mode)
For maximum safety with extremely low risk:

```go
RequestsPerSecond:  0.1   // 1 request every 10 seconds
MinDelay:          5s
MaxDelay:          10s
BatchDelay:        60s
MaxRetries:        10
RetryBackoff:      30s
MessagesPerBatch:  10
```

**Estimated time**: ~30 seconds per message

### Speed Demon (High Risk)
For disposable accounts or when speed is critical:

```go
RequestsPerSecond:  2.0   // 2 requests per second
MinDelay:          300ms
MaxDelay:          800ms
BatchDelay:        2s
MaxRetries:        2
RetryBackoff:      2s
MessagesPerBatch:  100
```

**Estimated time**: ~1 second per message
**Warning**: Very high ban risk!

### Night Owl (Slow and Steady)
For running overnight without triggering alarms:

```go
RequestsPerSecond:  0.05  // 1 request every 20 seconds
MinDelay:          10s
MaxDelay:          20s
BatchDelay:        120s  // 2 minutes between batches
MaxRetries:        10
RetryBackoff:      30s
MessagesPerBatch:  15
```

**Estimated time**: ~45 seconds per message

## Rate Limit Math

To calculate how long deletion will take:

```
Time per message = (MinDelay + MaxDelay) / 2 + (1 / RequestsPerSecond)
Time per batch = Time per message * MessagesPerBatch + BatchDelay
```

Example for Default profile:
```
Time per message = (1.5s + 3s) / 2 + (1 / 0.5) = 2.25s + 2s = 4.25s
Time per batch = 4.25s * 35 + 15s = 163.75s ≈ 2.7 minutes
```

For 1000 messages:
```
Batches needed = 1000 / 35 = 29 batches
Total time = 29 * 2.7min = 78.3 minutes ≈ 1.3 hours
```

## Important Notes

1. **Lower is Safer**: Lower `RequestsPerSecond` and higher delays are safer
2. **Randomization Helps**: Always keep `MinDelay < MaxDelay` for randomization
3. **Batch Delays Matter**: Longer `BatchDelay` helps avoid sustained high-rate activity
4. **Test First**: Always test custom configurations with disposable accounts
5. **Monitor Results**: Watch for rate limit errors and adjust accordingly

## Discord Rate Limits

Discord's actual rate limits are not publicly documented, but general observations:

- **Global Rate Limit**: ~50 requests per second (shared across all endpoints)
- **Per-Route Rate Limit**: Varies by endpoint (typically 5-10 per second)
- **Message Delete**: Conservative estimates suggest 1-2 deletes per second is safe

These are conservative estimates. Actual limits may vary.

## Rebuilding After Changes

After modifying configuration files:

```bash
go build -o godel
```

## Safety Reminder

Remember: Any automation violates Discord's Terms of Service. Even with the safest configuration, account bans are possible. Use responsibly and at your own risk.
