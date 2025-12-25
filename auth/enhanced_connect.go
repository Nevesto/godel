package auth

import (
	"github.com/Nevesto/godel/pkg/client"
	"github.com/Nevesto/godel/pkg/config"
)

// ConnectEnhanced creates an enhanced Discord client with security features
func ConnectEnhanced(securityConfig *config.SecurityConfig) (*client.EnhancedClient, error) {
	session, err := Connect()
	if err != nil {
		return nil, err
	}

	enhancedClient := client.NewEnhancedClient(session, securityConfig)
	return enhancedClient, nil
}

// ConnectEnhancedDefault creates an enhanced Discord client with default security settings
func ConnectEnhancedDefault() (*client.EnhancedClient, error) {
	return ConnectEnhanced(config.DefaultSecurityConfig())
}
