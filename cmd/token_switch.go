package cmd

import (
	"fmt"

	"github.com/Nevesto/godel/scripts"
	"github.com/spf13/cobra"
)

var tokenSwitchCmd = &cobra.Command{
	Use:   "token-switch [name]",
	Short: "Changes the active token to the specified alias",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		tokenName := args[0]

		if err := scripts.SwitchToken(tokenName); err != nil {
			fmt.Println("Error switching token:", err)
			return
		}
		fmt.Printf("Active token is now: %s\n", tokenName)
	},
}

func init() {
	rootCmd.AddCommand(tokenSwitchCmd)
}
