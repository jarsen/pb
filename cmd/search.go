package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/blevesearch/bleve"
	_ "github.com/blevesearch/bleve/config"
	"github.com/blevesearch/bleve/search/highlight/format/ansi"
	"github.com/jarsen/pb/db"
	. "github.com/spf13/cobra"
)

// searchCmd represents the search command
var searchCmd = &Command{
	Use:   "search [keywords*]",
	Short: "search images from the database",
	Args:  MinimumNArgs(1),
	Run: func(cmd *Command, args []string) {
		queryString := strings.Join(args, " ")
		query := bleve.NewMatchQuery(queryString)
		query.SetField("Description")
		searchRequest := bleve.NewSearchRequest(query)
		searchRequest.Highlight = bleve.NewHighlightWithStyle(ansi.Name)
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
			rv := ""
			for _, fragments := range hit.Fragments {
				for _, fragment := range fragments {
					rv += fmt.Sprintf("%s", fragment)
				}
			}
			fmt.Println(hit.Fields["URL"], rv)
		}
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
}
