package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/atotto/clipboard"
	"github.com/blevesearch/bleve"
	"github.com/blevesearch/bleve/search/highlight/format/ansi"
	"github.com/jarsen/pb/db"
	"github.com/spf13/cobra"
)

var copyCmd = &cobra.Command{
	Use:   "copy [SEARCH_TERMS..]",
	Short: "Copies the highest matching result to your system clipboard",
	Run: func(cmd *cobra.Command, args []string) {
		copyFirstSearchResult(args)
	},
}

func copyFirstSearchResult(args []string) {
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
	topMatch := results.Hits[0]
	rv := ""
	for _, fragments := range topMatch.Fragments {
		for _, fragment := range fragments {
			rv += fmt.Sprintf("%s", fragment)
		}
	}
	url := topMatch.Fields["URL"].(string)
	fmt.Println(url, rv)
	clipboard.WriteAll(url)
}

func init() {
	rootCmd.AddCommand(copyCmd)
}
