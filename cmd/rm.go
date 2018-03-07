package cmd

import (
	"log"

	"github.com/jarsen/pb/db"
	. "github.com/spf13/cobra"
)

var rmCmd = &Command{
	Use:   "rm ID",
	Short: "removes an entry from the database",
	Args:  ExactArgs(1),
	Run: func(cmd *Command, args []string) {
		id := args[0]
		index, err := db.Init()
		if err != nil {
			log.Fatal("Error loading database")
		}
		batch := index.NewBatch()
		batch.Delete(id)
		err = index.Batch(batch)
		if err != nil {
			log.Fatal("Error deleting from database")
		}
	},
}

func init() {
	rootCmd.AddCommand(rmCmd)
}
