package scripts

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
)

func ClearAllDms(session *discordgo.Session) error {

	endpoint := discordgo.EndpointUser("@me/channels")

	resp, err := session.Request("GET", endpoint, nil)
	if err != nil {
		fmt.Printf("Error fetching DM channels: %v\n", err)
		return nil
	}

	privateChannels := []*discordgo.Channel{}
	err = json.Unmarshal(resp, &privateChannels)
	if err != nil {
		fmt.Printf("Error unmarshalling DM channels: %v\n", err)
		return nil
	}

	if len(privateChannels) == 0 {
		fmt.Println("No DM channels found.")
		return nil
	} else {
		for _, ch := range privateChannels {
			if ch.Type == discordgo.ChannelTypeDM && len(ch.Recipients) > 0 {
				fmt.Printf("Recipient: %s\n", ch.Recipients[0].Username)
			}
		}
	}

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
