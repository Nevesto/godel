package scripts

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
)

func ClearDM(session *discordgo.Session, userID string) error {
	channel, err := session.UserChannelCreate(userID)
	if err != nil {
		return fmt.Errorf("error creating DM channel: %w", err)
	}

	for {
		messages, err := session.ChannelMessages(channel.ID, 100, "", "", "")
		if err != nil {
			return fmt.Errorf("error fetching messages: %w", err)
		}

		if len(messages) == 0 {
			break
		}

		for _, message := range messages {
			if message.Author.ID == session.State.User.ID {
				err := session.ChannelMessageDelete(channel.ID, message.ID)
				if err != nil {
					return fmt.Errorf("error deleting message: %w", err)
				} else {
					fmt.Printf("Deleted message: %s\n", message.Content)
				}

				time.Sleep(500 * time.Millisecond)
			}
		}
	}

	fmt.Println("All messages deleted successfully!")
	return nil
}
