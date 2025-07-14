package cmd

import (
	"os"
	"github.com/spf13/cobra"
)

var outenvCmd = &cobra.Command{
	Use:   "outenv",
	Short: "Show environment variables",
	Run: func(cmd *cobra.Command, args []string) {
		for _, env := range os.Environ() {
			println(env)
		}
	},
}
