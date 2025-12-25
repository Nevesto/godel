package scripts

import (
	"github.com/Nevesto/godel/pkg/cleaner"
	"github.com/Nevesto/godel/pkg/client"
	"github.com/Nevesto/godel/pkg/config"
	"github.com/bwmarrin/discordgo"
)

func ClearChannel(session *discordgo.Session, channelId string) error {
	// Create enhanced client with default security settings
	enhancedClient := client.NewEnhancedClient(session, config.DefaultSecurityConfig())
	
	// Create message cleaner
	messageCleaner := cleaner.NewMessageCleaner(enhancedClient)
	
	// Clear the channel using the new modular approach
	return messageCleaner.ClearChannel(channelId)
}
