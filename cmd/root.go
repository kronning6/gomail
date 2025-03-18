package cmd

import (
	"fmt"
	"os"

	"github.com/kronning6/gmail-screener/gmail"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gmail",
	Short: "Gmail is an email screening power tool",
	Run: func(cmd *cobra.Command, args []string) {
		gmail.Mail()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
