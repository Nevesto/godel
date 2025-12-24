package cmd

import (
	"github.com/Nevesto/godel/pkg/config"
)

// GetSecurityConfig returns the security configuration based on the selected profile
func GetSecurityConfig() *config.SecurityConfig {
	switch securityProfile {
	case "conservative":
		return config.ConservativeConfig()
	case "aggressive":
		return config.AggressiveConfig()
	default:
		return config.DefaultSecurityConfig()
	}
}
