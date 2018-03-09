package cmd

import (
	"fmt"
	"log"

	"github.com/jarsen/pb/db"
	. "github.com/spf13/cobra"
)

var listCmd = &Command{
	Use:   "list",
	Short: "list all the database entries",
	Args:  ExactArgs(0),
	Run: func(cmd *Command, args []string) {
		index, initErr := db.Init()
		if initErr != nil {
			log.Fatal(initErr)
		}

		results, err := db.AllImages(index)
		if err != nil {
			log.Fatalf("Error getting images: %s\n", err)
		}
		for _, hit := range results.Hits {
			fmt.Printf("[%s] %s %s\n", hit.Fields["ID"], hit.Fields["URL"], hit.Fields["Description"])
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
