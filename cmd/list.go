package cmd

import (
	. "github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &Command{
	Use:   "list [keywords*]",
	Short: "list images from the database",
	Args:  MinimumNArgs(1),
	Run: func(cmd *Command, args []string) {
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
