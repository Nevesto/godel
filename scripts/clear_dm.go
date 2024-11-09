package scripts

import (
	"context"
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
	"golang.org/x/time/rate"
)

func ClearDM(session *discordgo.Session, userID string) error {
	channel, err := session.UserChannelCreate(userID)
	if err != nil {
		return fmt.Errorf("can't create dm: %w", err)
	}

	startTime := time.Now()
	totalDeleted := 0
	var myMessages []*discordgo.Message

	var beforeID string

	for {
		messages, err := session.ChannelMessages(channel.ID, 100, beforeID, "", "")
		if err != nil {
			return fmt.Errorf("error when fetching messages: %w", err)
		}

		if len(messages) == 0 {
			fmt.Println("All messages saved.")
			break
		}

		for _, message := range messages {
			if message.Author.ID == session.State.User.ID {
				myMessages = append(myMessages, message)
			}
		}

		beforeID = messages[len(messages)-1].ID
	}

	limiter := rate.NewLimiter(rate.Every(250*time.Millisecond), 1)

	ctx := context.Background()

	for _, message := range myMessages {
		err := limiter.Wait(ctx)
		if err != nil {
			return fmt.Errorf("rate limit error: %w", err)
		}

		err = session.ChannelMessageDelete(channel.ID, message.ID)
		if err != nil {
			return fmt.Errorf("can't delete message: %w", err)
		} else {
			fmt.Printf("message deleted: %s\n", message.Content)
			totalDeleted++
		}
	}

	elapsedTime := time.Since(startTime)
	fmt.Printf("All messages deleted! Total: %d, Time: %s\n", totalDeleted, elapsedTime)
	return nil
}
