package cmd

import (
	"fmt"
	"os"

	"github.com/kronning6/gomail/gmail"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gomail",
	Short: "gomail is an email screening power tool for Gmail",
	Run: func(cmd *cobra.Command, args []string) {
		gmail.Screener()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
