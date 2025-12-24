package cmd

import (
	"github.com/Nevesto/godel/auth"
	"github.com/Nevesto/godel/pkg/cleaner"
	"github.com/Nevesto/godel/pkg/client"
	"github.com/spf13/cobra"
)

var clearDm = &cobra.Command{
	Use:   "clear-dm",
	Short: "Clears all messages on a specific DM channel",
	Long:  `Clears all your messages from a specific DM channel. Uses safe rate limiting to avoid account bans.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		channelId := args[0]
		
		// Get security configuration based on profile
		secConfig := GetSecurityConfig()
		
		session, err := auth.Connect()
		if err != nil {
			cmd.Println("Failed to connect to discord:", err)
			return
		}

		if session == nil {
			cmd.Println("Failed to connect to discord.")
			return
		}

		defer session.Close()

		// Create enhanced client and message cleaner
		enhancedClient := client.NewEnhancedClient(session, secConfig)
		messageCleaner := cleaner.NewMessageCleaner(enhancedClient)

		// Try to get or create the DM channel
		var actualChannelID string
		channel, err := session.Channel(channelId)
		if err != nil {
			// If not found, try to create a DM channel assuming the ID is a user ID
			channel, err = session.UserChannelCreate(channelId)
			if err != nil {
				cmd.Println("Can't create DM:", err)
				return
			}
		}
		actualChannelID = channel.ID

		// Clear the DM
		err = messageCleaner.ClearDM(actualChannelID)
		if err != nil {
			cmd.Println("Failed to clear DM:", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(clearDm)
}
