package cmd

import (
	"github.com/Nevesto/godel/auth"
	"github.com/Nevesto/godel/scripts"
	"github.com/spf13/cobra"
)

var clear_guild = &cobra.Command{
	Use:   "clear-guild",
	Short: "Clears all messages on a specific guild",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		guildId := args[0]
		session, err := auth.Connect()
		if err != nil {
			cmd.Println("Failed to connect to discord:", err)
		}

		if session == nil {
			cmd.Println("Failed to connect to discord.")
			return
		}

		defer session.Close()

		err = scripts.ClearGuild(session, guildId)
		if err != nil {
			cmd.Println("Failed to clear DM:", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(clear_guild)
}
