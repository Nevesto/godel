package cmd

import (
	"github.com/Nevesto/godel/auth"
	"github.com/Nevesto/godel/scripts"
	"github.com/spf13/cobra"
)

var clearDm = &cobra.Command{
	Use:   "clear-dm",
	Short: "Clears all messages on a specific channel",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		channelId := args[0]
		session, err := auth.Connect()
		if err != nil {
			cmd.Println("Failed to connect to discord:", err)
		}

		if session == nil {
			cmd.Println("Failed to connect to discord.")
			return
		}

		defer session.Close()

		err = scripts.ClearDM(session, channelId)
		if err != nil {
			cmd.Println("Failed to clear DM:", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(clearDm)
}
