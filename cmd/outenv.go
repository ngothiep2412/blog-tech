package cmd

import (
	"github.com/spf13/cobra"
)

var outenvCmd = &cobra.Command{
	Use:   "outenv",
	Short: "Show environment variables",
	Run: func(cmd *cobra.Command, args []string) {
		NewServiceCtx().OutEnv()
	},
}
