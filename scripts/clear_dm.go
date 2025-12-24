package scripts

import (
	"fmt"

	"github.com/Nevesto/godel/pkg/cleaner"
	"github.com/Nevesto/godel/pkg/client"
	"github.com/Nevesto/godel/pkg/config"
	"github.com/bwmarrin/discordgo"
	"github.com/fatih/color"
)

func ClearDM(session *discordgo.Session, id string, isChannel bool) error {
	var channelID string

	if isChannel {
		// Assume the ID is already a valid channel ID
		channelID = id
	} else {
		// Attempt to fetch the channel. If it exists, use it.
		channel, err := session.Channel(id)
		if err != nil {
			// If not found, try to create a DM channel assuming the ID is a user ID.
			channel, err = session.UserChannelCreate(id)
			if err != nil {
				return fmt.Errorf("%s", color.RedString(fmt.Sprintf("can't create dm: %v", err)))
			}
		}
		channelID = channel.ID
	}

	// Create enhanced client with default security settings
	enhancedClient := client.NewEnhancedClient(session, config.DefaultSecurityConfig())
	
	// Create message cleaner
	messageCleaner := cleaner.NewMessageCleaner(enhancedClient)
	
	// Clear the DM using the new modular approach
	return messageCleaner.ClearDM(channelID)
}
