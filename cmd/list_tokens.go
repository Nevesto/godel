package cmd

import (
	"fmt"

	"github.com/Nevesto/godel/scripts"
	"github.com/spf13/cobra"
)

var showUsersCmd = &cobra.Command{
	Use:   "show-users",
	Short: "Lists the aliases of registered tokens",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		aliases, err := scripts.ListTokens()
		if err != nil {
			fmt.Println("Error listing tokens:", err)
			return
		}
		fmt.Println("Registered tokens:")
		for _, alias := range aliases {
			fmt.Printf(" - %s\n", alias)
		}
	},
}

func init() {
	rootCmd.AddCommand(showUsersCmd)
}
