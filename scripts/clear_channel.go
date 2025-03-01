package scripts

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
)

func ClearChannel(session *discordgo.Session, channelId string) error {
	totalDeleted := 0
	var beforeID string

	for {
		messages, err := session.ChannelMessages(channelId, 35, beforeID, "", "")
		if err != nil {
			return err
		}

		userMessageCount := 0
		for _, message := range messages {
			if message.Author.ID == session.State.User.ID {
				userMessageCount++
			}
		}
		fmt.Println("Messages to delete: ", userMessageCount)

		if len(messages) == 0 {
			fmt.Println("No messages to delete")
			break
		}

		currentBatchDeleted := 0
		for _, message := range messages {
			if message.Author.ID == session.State.User.ID {
				err := session.ChannelMessageDelete(channelId, message.ID)
				if err != nil {
					return fmt.Errorf("\ncan't delete message: %w", err)
				}

				fmt.Printf("Deleted message: %s\n", message.Content)
				totalDeleted++
				currentBatchDeleted++
				time.Sleep(1 * time.Second)
			}
		}

		if len(messages) > 0 {
			beforeID = messages[len(messages)-1].ID
		}

		if len(messages) < 35 {
			fmt.Printf("Reached final batch (%d messages)\n", len(messages))
			break
		}

		fmt.Printf("Deleted %d/%d user messages in this batch\n",
			currentBatchDeleted,
			userMessageCount,
		)
		fmt.Println("Waiting 15 seconds before fetching next batch of messages...")
		time.Sleep(15 * time.Second)
	}

	fmt.Printf("\nTotal messages deleted: %d\n", totalDeleted)
	return nil
}
