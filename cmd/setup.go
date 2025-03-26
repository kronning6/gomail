package cmd

import (
	"github.com/kronning6/gomail/gmail"
	"github.com/spf13/cobra"
)

var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Setup connection to Gmail",
	Run: func(cmd *cobra.Command, args []string) {
		gmail.Setup()
	},
}
