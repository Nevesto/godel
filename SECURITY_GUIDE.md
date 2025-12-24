# Security Guide

## Overview

Godel includes several security features to help reduce the risk of account bans when clearing Discord messages. This guide explains how these features work and how to use them effectively.

## Security Profiles

Godel offers three pre-configured security profiles that balance speed and safety:

### Conservative Profile (Recommended for New Users)
**Use this if**: You prioritize account safety over speed, or it's your first time using the tool.

**Settings**:
- Rate: 1 request every 4 seconds (0.25 req/s)
- Delay between messages: 2.5-5 seconds (randomized)
- Delay between batches: 30 seconds
- Messages per batch: 25
- Max retries: 5

**Usage**:
```bash
godel clear-all-dms --security conservative
```

**Pros**:
- Lowest risk of account ban
- Best for accounts you want to keep safe
- Recommended for high-value accounts

**Cons**:
- Slowest option
- Takes longer to clear many messages

### Default Profile (Balanced)
**Use this if**: You want a balance between speed and safety.

**Settings**:
- Rate: 1 request every 2 seconds (0.5 req/s)
- Delay between messages: 1.5-3 seconds (randomized)
- Delay between batches: 15 seconds
- Messages per batch: 35
- Max retries: 3

**Usage**:
```bash
godel clear-all-dms
# or explicitly:
godel clear-all-dms --security default
```

**Pros**:
- Good balance of speed and safety
- Reasonable protection against bans
- Default choice for most users

**Cons**:
- Moderate risk compared to conservative
- Still relatively slow for large amounts of messages

### Aggressive Profile (Advanced Users)
**Use this if**: You're willing to accept higher risk for faster deletion, or using a disposable account.

**Settings**:
- Rate: 1 request per second (1.0 req/s)
- Delay between messages: 0.8-1.5 seconds (randomized)
- Delay between batches: 5 seconds
- Messages per batch: 50
- Max retries: 3

**Usage**:
```bash
godel clear-all-dms --security aggressive
```

**Pros**:
- Fastest option
- Good for disposable accounts
- Useful when time is critical

**Cons**:
- Higher risk of account ban
- May trigger Discord's rate limiting
- Not recommended for main accounts

## Built-in Security Features

All profiles include these safety mechanisms:

### 1. Rate Limiting
Prevents sending too many requests too quickly to Discord's API. Each profile has different limits optimized for different use cases.

### 2. Random Delays
Instead of using fixed delays (which look bot-like), Godel adds randomized delays between operations to appear more human-like. For example, with default profile, delays range from 1.5 to 3 seconds randomly.

### 3. Exponential Backoff
When Discord returns a rate limit error (HTTP 429), Godel automatically:
- Waits longer before retrying
- Doubles the wait time with each failed attempt
- Adds random jitter to avoid synchronized retries

### 4. Batch Processing
Messages are processed in batches with longer delays between batches. This prevents sustained high-rate activity that might trigger detection.

### 5. User Agent Rotation
Godel can rotate between different browser user agents to appear like regular browser traffic (currently configured but not actively rotating in requests).

## Best Practices

1. **Start Conservative**: If you're unsure, always start with the conservative profile.

2. **Test with Disposable Account**: If possible, test with an account you don't care about first.

3. **Monitor for Rate Limits**: If you see rate limit errors even with conservative settings, consider adding manual delays between runs.

4. **Don't Run 24/7**: Avoid running the tool continuously. Take breaks between clearing sessions.

5. **Use During Peak Hours**: Running during times when Discord is busiest may help blend in with normal traffic.

6. **Keep Token Safe**: Never share your Discord token. Treat it like a password.

7. **Understand the Risk**: No tool can guarantee you won't be banned. Discord explicitly prohibits account automation.

## Customization

If you need custom settings beyond the three profiles, you can modify the security configurations in `pkg/config/security.go`.

## Rate Limit Handling

If you encounter rate limits:

1. **Let the Tool Handle It**: Godel automatically backs off and retries
2. **Switch to Conservative**: Use a slower profile
3. **Take a Break**: Wait several hours before trying again
4. **Check Discord Status**: Ensure Discord's API is operating normally

## Warning Signs

Stop using the tool if you notice:
- Repeated rate limit errors (HTTP 429)
- Account verification requests
- Unusual security checks from Discord
- Temporary account restrictions

## Legal Disclaimer

Using any form of automation with Discord violates their Terms of Service. This tool is for educational purposes only. The developers are not responsible for any account bans or restrictions. Use at your own risk.
