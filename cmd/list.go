package cmd

import (
	"fmt"
	"log"

	"github.com/blevesearch/bleve"
	"github.com/jarsen/pb/db"
	. "github.com/spf13/cobra"
)

var listCmd = &Command{
	Use:   "list",
	Short: "list all the database entries",
	Args:  ExactArgs(0),
	Run: func(cmd *Command, args []string) {
		query := bleve.NewMatchAllQuery()
		searchRequest := bleve.NewSearchRequest(query)
		searchRequest.Fields = []string{"*"}
		index, initErr := db.Init()
		if initErr != nil {
			log.Fatal(initErr)
		}
		results, err := index.Search(searchRequest)
		if err != nil {
			log.Fatal(err)
		}
		for _, hit := range results.Hits {
			fmt.Printf("[%s] %s %s\n", hit.Fields["ID"], hit.Fields["URL"], hit.Fields["Description"])
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
