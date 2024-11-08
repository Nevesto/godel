package cmd

import (
	"github.com/Nevesto/godel/auth"
	"github.com/spf13/cobra"
)

var setTokenCmd = &cobra.Command{
	Use:   "set-token [token]",
	Short: "Sets and saves your Discord account token",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		token := args[0]
		auth.SaveToken(token)
	},
}

func init() {
	rootCmd.AddCommand(setTokenCmd)
}
