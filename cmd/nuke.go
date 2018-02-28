package cmd

import (
	"os"

	"github.com/jarsen/pb/db"
	. "github.com/spf13/cobra"
)

// nukeCmd represents the nuke command
var nukeCmd = &Command{
	Use:   "nuke",
	Short: "deletes the entire database",
	Run: func(cmd *Command, args []string) {
		dbPath := db.Path()
		os.RemoveAll(dbPath)
	},
}

func init() {
	rootCmd.AddCommand(nukeCmd)
}
