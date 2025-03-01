// read this: follow my friend <3: https://www.instagram.com/_jarxdd/

package scripts

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
)

func ClearDM(session *discordgo.Session, userID string) error {
	channel, err := session.UserChannelCreate(userID)
	if err != nil {
		return fmt.Errorf("can't create dm: %w", err)
	}

	totalDeleted := 0
	var beforeID string

	for {
		messages, err := session.ChannelMessages(channel.ID, 35, beforeID, "", "")
		if err != nil {
			return fmt.Errorf("error fetching messages: %w", err)
		}

		// Pre-count user messages
		var userMessages []*discordgo.Message
		for _, msg := range messages {
			if msg.Author.ID == session.State.User.ID {
				userMessages = append(userMessages, msg)
			}
		}
		fmt.Printf("Potential messages to delete: %d (from %d fetched)\n", len(userMessages), len(messages)) // lots of messages to delete, value is not accurate

		if len(messages) == 0 {
			fmt.Println("No messages found in this batch")
			break
		}

		currentBatchDeleted := 0
		for _, message := range userMessages {
			if err := session.ChannelMessageDelete(channel.ID, message.ID); err != nil {
				return fmt.Errorf("can't delete message: %w", err)
			}

			// Handle empty message content
			content := message.Content
			if content == "" {
				content = "[empty content]"
			}
			fmt.Printf("Deleted message: %.40s\n", content)

			currentBatchDeleted++
			totalDeleted++
			time.Sleep(1 * time.Second)
		}

		// Update beforeID after processing entire batch
		if len(messages) > 0 {
			beforeID = messages[len(messages)-1].ID
		}

		if len(messages) < 35 {
			fmt.Printf("Reached final batch (%d messages)\n", len(messages))
			break
		}

		fmt.Printf("Deleted %d messages in this batch\n", currentBatchDeleted)
		fmt.Println("Waiting 15 seconds before next batch...")
		time.Sleep(15 * time.Second)
	}

	fmt.Printf("\nTotal messages deleted: %d\n", totalDeleted)
	return nil
}
