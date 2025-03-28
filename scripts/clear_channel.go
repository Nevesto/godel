package scripts

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/fatih/color"
)

func ClearChannel(session *discordgo.Session, channelId string) error {
	totalDeleted := 0
	var beforeID string

	for {
		messages, err := session.ChannelMessages(channelId, 35, beforeID, "", "")
		if err != nil {
			return fmt.Errorf(color.RedString("Failed to fetch messages: %w"), err)
		}

		userMessageCount := 0
		for _, message := range messages {
			if message.Author.ID == session.State.User.ID {
				userMessageCount++
			}
		}
		fmt.Println(color.YellowString("Messages to delete: %d"), userMessageCount)

		if len(messages) == 0 {
			fmt.Println(color.GreenString("No messages to delete"))
			break
		}

		currentBatchDeleted := 0
		for _, message := range messages {
			if message.Author.ID == session.State.User.ID {
				err := session.ChannelMessageDelete(channelId, message.ID)
				if err != nil {
					return fmt.Errorf(color.RedString("Can't delete message: %w"), err)
				}

				fmt.Printf(color.RedString("Deleted message: %s\n"), message.Content)
				totalDeleted++
				currentBatchDeleted++
				time.Sleep(1 * time.Second)
			}
		}

		if len(messages) > 0 {
			beforeID = messages[len(messages)-1].ID
		}

		if len(messages) < 35 {
			fmt.Println(color.GreenString("Reached final batch (%d messages)"), len(messages))
			break
		}

		fmt.Printf(color.GreenString("Deleted %d/%d user messages in this batch\n"),
			currentBatchDeleted,
			userMessageCount,
		)
		fmt.Println(color.YellowString("Waiting 15 seconds before fetching next batch of messages..."))
		time.Sleep(15 * time.Second)
	}

	fmt.Printf(color.GreenString("Total messages deleted: %d\n"), totalDeleted)
	return nil
}
