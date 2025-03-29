package scripts

import (
	"fmt"
	"time"

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

	totalDeleted := 0
	var beforeID string

	for {
		messages, err := session.ChannelMessages(channelID, 35, beforeID, "", "")
		if err != nil {
			return fmt.Errorf("error fetching messages: %w", err)
		}

		var userMessages []*discordgo.Message
		for _, msg := range messages {
			if msg.Author.ID == session.State.User.ID {
				userMessages = append(userMessages, msg)
			}
		}
		color.Green("Potential messages to delete: %d (from %d fetched)", len(userMessages), len(messages))

		if len(messages) == 0 {
			color.Yellow("No messages found in this batch")
			break
		}

		currentBatchDeleted := 0
		for _, message := range userMessages {
			if err := session.ChannelMessageDelete(channelID, message.ID); err != nil {
				return fmt.Errorf("%s", color.RedString(fmt.Sprintf("can't delete message: %v", err)))
			}

			content := message.Content
			if content == "" {
				content = "[empty content]"
			}
			color.Green("Deleted message: %.40s", content)

			currentBatchDeleted++
			totalDeleted++
			time.Sleep(1 * time.Second)
		}

		if len(messages) > 0 {
			beforeID = messages[len(messages)-1].ID
		}

		if len(messages) < 35 {
			color.Green("Reached final batch (%d messages)", len(messages))
			break
		}

		color.Green("Deleted %d messages in this batch", currentBatchDeleted)
		color.Yellow("Waiting 15 seconds before next batch...")
		time.Sleep(15 * time.Second)
	}

	color.Green("\nTotal messages deleted: %d", totalDeleted)
	return nil
}
