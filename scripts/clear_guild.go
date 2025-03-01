package scripts

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
)

func ClearGuild(session *discordgo.Session, guildID string) error {
	channels, err := session.GuildChannels(guildID)
	if err != nil {
		return fmt.Errorf("failed to get guild channels: %w", err)
	}

	for _, channel := range channels {
		if channel.Type != discordgo.ChannelTypeGuildText {
			continue
		}

		fmt.Println("Clearing channel: ", channel.Name)

		err := ClearChannel(session, channel.ID)
		if err != nil {
			return fmt.Errorf("channel cleanup failed: %w", err)
		}

		time.Sleep(3 * time.Second) // Avoid rate limiting
	}

	return nil
}
