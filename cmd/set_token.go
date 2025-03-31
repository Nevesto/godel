package cmd

import (
	"github.com/Nevesto/godel/auth"
	"github.com/spf13/cobra"
)

var setTokenCmd = &cobra.Command{
	Use:   "set-token [name] [token]",
	Short: "Register and save a Discord token with an identifier name",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		tokenName := args[0]
		token := args[1]
		auth.SaveToken(tokenName, token)
	},
}

func init() {
	rootCmd.AddCommand(setTokenCmd)
}
