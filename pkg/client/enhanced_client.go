package client

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"

	"github.com/Nevesto/godel/pkg/config"
	"github.com/Nevesto/godel/pkg/ratelimit"
	"github.com/bwmarrin/discordgo"
)

// EnhancedClient wraps the Discord session with additional security features
type EnhancedClient struct {
	Session      *discordgo.Session
	Config       *config.SecurityConfig
	RateLimiter  *ratelimit.Limiter
	userAgents   []string
}

// NewEnhancedClient creates a new enhanced Discord client
func NewEnhancedClient(session *discordgo.Session, securityConfig *config.SecurityConfig) *EnhancedClient {
	rateLimiter := ratelimit.NewLimiter(
		securityConfig.RequestsPerSecond,
		securityConfig.MinDelay,
		securityConfig.MaxDelay,
		securityConfig.BatchDelay,
	)

	return &EnhancedClient{
		Session:     session,
		Config:      securityConfig,
		RateLimiter: rateLimiter,
		userAgents:  config.GetUserAgents(),
	}
}

// GetRandomUserAgent returns a random user agent from the list
func (c *EnhancedClient) GetRandomUserAgent() string {
	if !c.Config.UseRandomUserAgent || len(c.userAgents) == 0 {
		return c.Config.UserAgent
	}
	return c.userAgents[rand.Intn(len(c.userAgents))]
}

// RequestWithRetry performs an HTTP request with retry logic
func (c *EnhancedClient) RequestWithRetry(method, endpoint string, data interface{}) ([]byte, error) {
	var lastErr error
	
	for attempt := 0; attempt < c.Config.MaxRetries; attempt++ {
		if attempt > 0 {
			c.RateLimiter.WaitWithBackoff(attempt)
		} else {
			c.RateLimiter.Wait()
		}

		// Set custom user agent
		if c.Session.Client == nil {
			c.Session.Client = &http.Client{}
		}
		
		resp, err := c.Session.Request(method, endpoint, data)
		if err != nil {
			lastErr = err
			// Check if it's a rate limit error
			if restErr, ok := err.(*discordgo.RESTError); ok {
				if restErr.Response != nil && restErr.Response.StatusCode == 429 {
					// Rate limited, continue to retry with backoff
					continue
				}
			}
			// For other errors, return immediately
			return nil, err
		}
		
		return resp, nil
	}
	
	return nil, fmt.Errorf("max retries exceeded: %w", lastErr)
}

// GetAllDMChannels retrieves all DM channels including closed ones
func (c *EnhancedClient) GetAllDMChannels() ([]*discordgo.Channel, error) {
	endpoint := discordgo.EndpointUser("@me/channels")
	
	resp, err := c.RequestWithRetry("GET", endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("error fetching DM channels: %w", err)
	}

	var channels []*discordgo.Channel
	err = json.Unmarshal(resp, &channels)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling DM channels: %w", err)
	}

	return channels, nil
}

// DeleteMessageSafe deletes a message with rate limiting and retry logic
func (c *EnhancedClient) DeleteMessageSafe(channelID, messageID string) error {
	var lastErr error
	
	for attempt := 0; attempt < c.Config.MaxRetries; attempt++ {
		if attempt > 0 {
			c.RateLimiter.WaitWithBackoff(attempt)
		} else {
			c.RateLimiter.Wait()
		}

		err := c.Session.ChannelMessageDelete(channelID, messageID)
		if err != nil {
			lastErr = err
			// Check if it's a rate limit error
			if restErr, ok := err.(*discordgo.RESTError); ok {
				if restErr.Response != nil && restErr.Response.StatusCode == 429 {
					// Rate limited, continue to retry with backoff
					continue
				}
			}
			// For other errors, return immediately
			return err
		}
		
		return nil
	}
	
	return fmt.Errorf("max retries exceeded: %w", lastErr)
}

// GetChannelMessages retrieves messages from a channel with rate limiting
func (c *EnhancedClient) GetChannelMessages(channelID string, limit int, beforeID, afterID, aroundID string) ([]*discordgo.Message, error) {
	c.RateLimiter.Wait()
	
	messages, err := c.Session.ChannelMessages(channelID, limit, beforeID, afterID, aroundID)
	if err != nil {
		return nil, err
	}
	
	return messages, nil
}
