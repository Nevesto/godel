package cmd

import (
	"github.com/Nevesto/godel/auth"
	"github.com/Nevesto/godel/scripts"
	"github.com/spf13/cobra"
)

var clearAllDms = &cobra.Command{
	Use:   "clear-all-dms",
	Short: "Clears all messages on all DM channels",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		session, err := auth.Connect()
		if err != nil {
			cmd.Println("Failed to connect to discord: ", err)
		}

		if session == nil {
			cmd.Println("Failed to connect to discord.")
			return
		}

		defer session.Close()

		err = scripts.ClearAllDms(session)
		if err != nil {
			cmd.Println("Failed to clear all DMs: ", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(clearAllDms)
}
