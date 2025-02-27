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
		messages, err := session.ChannelMessages(channel.ID, 30, beforeID, "", "")
		if err != nil {
			return fmt.Errorf("error when fetching messages: %w", err)
		}

		if len(messages) == 0 {
			fmt.Println("No more messages to delete.")
			break
		}

		for _, message := range messages {
			if message.Author.ID == session.State.User.ID {
				if err := session.ChannelMessageDelete(channel.ID, message.ID); err != nil {
					return fmt.Errorf("can't delete message: %w", err)
				}
				fmt.Printf("Deleted message: %s\n", message.Content)
				totalDeleted++

				time.Sleep(1 * time.Second) // Avoid rate limiting
			}
		}

		time.Sleep(15 * time.Second)

		beforeID = messages[len(messages)-1].ID
	}

	fmt.Printf("Deleted %d messages\n", totalDeleted)
	return nil
}
