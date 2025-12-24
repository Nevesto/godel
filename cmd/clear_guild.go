package cmd

import (
	"github.com/Nevesto/godel/auth"
	"github.com/Nevesto/godel/pkg/cleaner"
	"github.com/Nevesto/godel/pkg/client"
	"github.com/spf13/cobra"
)

var clear_guild = &cobra.Command{
	Use:   "clear-guild",
	Short: "Clears all messages on a specific guild",
	Long:  `Clears all your messages from all text channels in a specific guild. Uses safe rate limiting to avoid account bans.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		guildId := args[0]
		
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

		// Clear the guild
		err = messageCleaner.ClearGuild(guildId)
		if err != nil {
			cmd.Println("Failed to clear guild:", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(clear_guild)
}
