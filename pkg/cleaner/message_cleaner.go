package cleaner

import (
	"fmt"

	"github.com/Nevesto/godel/pkg/client"
	"github.com/Nevesto/godel/pkg/config"
	"github.com/bwmarrin/discordgo"
	"github.com/fatih/color"
)

// MessageCleaner handles message deletion operations
type MessageCleaner struct {
	client *client.EnhancedClient
	config *config.SecurityConfig
}

// NewMessageCleaner creates a new message cleaner
func NewMessageCleaner(enhancedClient *client.EnhancedClient) *MessageCleaner {
	return &MessageCleaner{
		client: enhancedClient,
		config: enhancedClient.Config,
	}
}

// ClearChannel clears all user messages in a channel
func (mc *MessageCleaner) ClearChannel(channelID string) error {
	totalDeleted := 0
	var beforeID string

	for {
		messages, err := mc.client.GetChannelMessages(
			channelID,
			mc.config.MessagesPerBatch,
			beforeID,
			"",
			"",
		)
		if err != nil {
			return fmt.Errorf("failed to fetch messages: %w", err)
		}

		if len(messages) == 0 {
			color.Yellow("No more messages to delete")
			break
		}

		// Count user messages
		userMessageCount := 0
		for _, message := range messages {
			if message.Author.ID == mc.client.Session.State.User.ID {
				userMessageCount++
			}
		}
		
		if userMessageCount > 0 {
			color.Yellow("Messages to delete in this batch: %d", userMessageCount)
		}

		// Delete user messages
		currentBatchDeleted := 0
		for _, message := range messages {
			if message.Author.ID == mc.client.Session.State.User.ID {
				err := mc.client.DeleteMessageSafe(channelID, message.ID)
				if err != nil {
					color.Red("Failed to delete message: %v", err)
					continue
				}

				content := message.Content
				if content == "" {
					content = "[empty/attachment]"
				}
				if len(content) > 40 {
					content = content[:40] + "..."
				}
				color.Green("Deleted: %s", content)

				currentBatchDeleted++
				totalDeleted++

				// Check if we've hit the max messages limit
				if mc.config.MaxMessagesTotal > 0 && totalDeleted >= mc.config.MaxMessagesTotal {
					color.Yellow("Reached maximum message limit (%d)", mc.config.MaxMessagesTotal)
					color.Green("Total messages deleted: %d", totalDeleted)
					return nil
				}
			}
		}

		// Update beforeID for next batch
		if len(messages) > 0 {
			beforeID = messages[len(messages)-1].ID
		}

		// Check if this was the last batch
		if len(messages) < mc.config.MessagesPerBatch {
			color.Green("Reached final batch (%d messages)", len(messages))
			break
		}

		if currentBatchDeleted > 0 {
			color.Green("Deleted %d messages in this batch", currentBatchDeleted)
			color.Yellow("Waiting before next batch...")
			mc.client.RateLimiter.BatchDelay()
		}
	}

	color.Green("\nTotal messages deleted: %d", totalDeleted)
	return nil
}

// ClearDM clears messages in a DM channel
func (mc *MessageCleaner) ClearDM(channelID string) error {
	return mc.ClearChannel(channelID)
}

// ClearAllDMs clears all messages in all DM channels
func (mc *MessageCleaner) ClearAllDMs() error {
	color.Cyan("Fetching all DM channels (including closed ones)...")
	
	channels, err := mc.client.GetAllDMChannels()
	if err != nil {
		return fmt.Errorf("error fetching DM channels: %w", err)
	}

	if len(channels) == 0 {
		color.Yellow("No DM channels found.")
		return nil
	}

	color.Green("Found %d DM channels:", len(channels))
	for _, ch := range channels {
		if ch.Type == discordgo.ChannelTypeDM && len(ch.Recipients) > 0 {
			color.Cyan("  - %s", ch.Recipients[0].Username)
		} else if ch.Type == discordgo.ChannelTypeGroupDM {
			color.Cyan("  - Group DM (%d recipients)", len(ch.Recipients))
		}
	}

	color.Yellow("\nStarting to clear messages...")
	
	successCount := 0
	errorCount := 0

	for _, channel := range channels {
		var channelName string
		if channel.Type == discordgo.ChannelTypeDM && len(channel.Recipients) > 0 {
			channelName = channel.Recipients[0].Username
		} else if channel.Type == discordgo.ChannelTypeGroupDM {
			channelName = fmt.Sprintf("Group DM (%d recipients)", len(channel.Recipients))
		} else {
			channelName = fmt.Sprintf("Channel %s", channel.ID)
		}

		color.Cyan("\n=== Clearing DM with: %s ===", channelName)
		
		err := mc.ClearDM(channel.ID)
		if err != nil {
			color.Red("Failed to clear DM with %s: %v", channelName, err)
			errorCount++
		} else {
			color.Green("Successfully cleared DM with %s", channelName)
			successCount++
		}

		// Extra delay between different DM channels
		mc.client.RateLimiter.BatchDelay()
	}

	color.Cyan("\n=== Summary ===")
	color.Green("Successfully cleared: %d", successCount)
	if errorCount > 0 {
		color.Red("Failed: %d", errorCount)
	}
	
	return nil
}

// ClearGuild clears all user messages in all text channels of a guild
func (mc *MessageCleaner) ClearGuild(guildID string) error {
	channels, err := mc.client.Session.GuildChannels(guildID)
	if err != nil {
		color.Red("Error getting guild channels: %v", err)
		return fmt.Errorf("failed to get guild channels: %w", err)
	}

	textChannelCount := 0
	for _, channel := range channels {
		if channel.Type == discordgo.ChannelTypeGuildText {
			textChannelCount++
		}
	}

	if textChannelCount == 0 {
		color.Yellow("No text channels found in the guild.")
		return nil
	}

	color.Green("Found %d text channels in the guild", textChannelCount)
	
	for _, channel := range channels {
		if channel.Type != discordgo.ChannelTypeGuildText {
			continue
		}

		color.Cyan("\n=== Clearing channel: %s ===", channel.Name)

		err := mc.ClearChannel(channel.ID)
		if err != nil {
			color.Red("Error clearing channel %s: %v", channel.Name, err)
			// Continue with other channels
		} else {
			color.Green("Successfully cleared channel: %s", channel.Name)
		}

		// Extra delay between channels
		mc.client.RateLimiter.BatchDelay()
	}

	return nil
}
