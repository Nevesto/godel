package scripts

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/fatih/color"
)

func ClearGuild(session *discordgo.Session, guildID string) error {
	channels, err := session.GuildChannels(guildID)
	if err != nil {
		color.Red("Error getting guild channels: %v", err)
		return fmt.Errorf("failed to get guild channels: %w", err)
	}

	for _, channel := range channels {
		if channel.Type != discordgo.ChannelTypeGuildText {
			continue
		}

		color.Green("Clearing channel: %s", channel.Name)

		err := ClearChannel(session, channel.ID)
		if err != nil {
			color.Red("Error clearing channel %s: %v", channel.Name, err)
			return fmt.Errorf("channel cleanup failed: %w", err)
		}

		time.Sleep(3 * time.Second) // Avoid rate limiting
	}

	return nil
}
