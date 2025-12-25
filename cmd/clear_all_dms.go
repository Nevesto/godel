package cmd

import (
	"github.com/Nevesto/godel/auth"
	"github.com/Nevesto/godel/pkg/cleaner"
	"github.com/Nevesto/godel/pkg/client"
	"github.com/spf13/cobra"
)

var clearAllDms = &cobra.Command{
	Use:   "clear-all-dms",
	Short: "Clears all messages on all DM channels (including closed DMs)",
	Long:  `Clears all your messages from all DM channels including those that are closed. Uses safe rate limiting to avoid account bans.`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		// Get security configuration based on profile
		secConfig := GetSecurityConfig()
		
		// Connect with enhanced security features
		session, err := auth.Connect()
		if err != nil {
			cmd.Println("Failed to connect to discord: ", err)
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

		// Clear all DMs
		err = messageCleaner.ClearAllDMs()
		if err != nil {
			cmd.Println("Failed to clear all DMs: ", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(clearAllDms)
}
