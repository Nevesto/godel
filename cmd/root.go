package cmd

import (
	"os"

	"github.com/Nevesto/godel/auth"
	"github.com/spf13/cobra"
)

var (
	securityProfile string
)

var rootCmd = &cobra.Command{
	Use:   "godel",
	Short: "A cli to clear dms on discord",
	Long:  `Godel is a command-line tool build with cobra to clear dms on discord.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		auth.LoadConfig()
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&securityProfile, "security", "s", "default", "Security profile: default, conservative, or aggressive")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
