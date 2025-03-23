package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var starredCmd = &cobra.Command{
	Use:   "starred",
	Short: "Show starred emails",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("This subcommand will show starred emails")
	},
}
