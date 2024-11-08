package auth

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/spf13/viper"
)

func Connect() (*discordgo.Session, error) {
	token := viper.GetString("discord_token")
	if token == "" {
		return nil, fmt.Errorf("discord token not set please run 'godel set-token [token]' to set it")
	}

	dg, err := discordgo.New(token)
	if err != nil {
		return nil, fmt.Errorf("error creating discord session: %v", err)
	}

	dg.Token = token
	dg.Identify.Token = token

	dg.StateEnabled = false
	dg.Identify.Compress = false
	dg.Identify.LargeThreshold = 0

	// This makes the account to apear offline, uncomment to enable (optional)
	// dg.Identify.Presence = discordgo.GatewayStatusUpdate{
	// 	Status: "invisible",
	// }

	err = dg.Open()
	if err != nil {
		return nil, fmt.Errorf("error opening connection to discord: %v", err)
	}

	fmt.Println("Connected to: " + dg.State.User.Username)
	return dg, nil
}
