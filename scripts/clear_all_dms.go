package scripts

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
)

func ClearAllDms(session *discordgo.Session) error {
	privateChannels := session.State.PrivateChannels

	for _, channel := range privateChannels {
		fmt.Println("Clearing DM with users: ", channel.Recipients[0].Username)
		err := ClearDM(session, channel.Recipients[0].ID)

		if err != nil {
			fmt.Printf("Failed to clear DM with user %s: %v\n", channel.Recipients[0].Username, err)
		} else {
			time.Sleep(1 * time.Second)
		}
	}

	return nil
}
