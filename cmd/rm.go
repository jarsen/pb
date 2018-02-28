package cmd

import (
	"fmt"

	. "github.com/spf13/cobra"
)

// rmCmd represents the rm command
var rmCmd = &Command{
	Use:   "rm ID",
	Short: "removes an entry from the database",
	Run: func(cmd *Command, args []string) {
		fmt.Println("not yet implemented")
	},
}

func init() {
	rootCmd.AddCommand(rmCmd)
}
