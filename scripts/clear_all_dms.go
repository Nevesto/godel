package scripts

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/fatih/color"
)

func ClearAllDms(session *discordgo.Session) error {
	endpoint := discordgo.EndpointUser("@me/channels")

	resp, err := session.Request("GET", endpoint, nil)
	if err != nil {
		fmt.Println(color.RedString(fmt.Sprintf("Error fetching DM channels: %v", err)))
		return nil
	}

	privateChannels := []*discordgo.Channel{}
	err = json.Unmarshal(resp, &privateChannels)
	if err != nil {
		fmt.Println(color.RedString(fmt.Sprintf("Error unmarshalling DM channels: %v", err)))
		return nil
	}

	if len(privateChannels) == 0 {
		fmt.Println(color.YellowString("No DM channels found."))
		return nil
	} else {
		for _, ch := range privateChannels {
			if ch.Type == discordgo.ChannelTypeDM && len(ch.Recipients) > 0 {
				fmt.Println(color.GreenString(fmt.Sprintf("Recipient: %s", ch.Recipients[0].Username)))
			}
		}
	}

	for _, channel := range privateChannels {
		fmt.Println(color.GreenString(fmt.Sprintf("Clearing DM with users: %s", channel.Recipients[0].Username)))
		err := ClearDM(session, channel.Recipients[0].ID)

		if err != nil {
			fmt.Println(color.RedString(fmt.Sprintf("Failed to clear DM with user %s: %v", channel.Recipients[0].Username, err)))
		} else {
			fmt.Println(color.GreenString(fmt.Sprintf("Successfully cleared DM with %s", channel.Recipients[0].Username)))
			time.Sleep(1 * time.Second)
		}
	}

	return nil
}
