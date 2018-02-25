package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/blevesearch/bleve"
	"github.com/jarsen/pb/db"
	. "github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &Command{
	Use:   "list [keywords*]",
	Short: "list images from the database",
	Args:  MinimumNArgs(1),
	Run: func(cmd *Command, args []string) {
		queryString := strings.Join(args, " ")
		// query := bleve.NewQueryStringQuery(queryString)
		query := bleve.NewMatchQuery(queryString)
		query.SetField("Description")
		searchRequest := bleve.NewSearchRequest(query)
		index, initErr := db.Init()
		if initErr != nil {
			log.Fatal(initErr)
		}
		results, err := index.Search(searchRequest)
		if err != nil {
			log.Fatal(err)
		}
		for _, hit := range results.Hits {
			uuid := hit.ID
			var buf []byte
			buf, err = index.GetInternal([]byte(uuid))
			if err != nil {
				log.Fatal(err)
			}
			image := db.Image{}
			if err := json.Unmarshal(buf, &image); err != nil {
				log.Fatal(err)
			}
			fmt.Println(image.Url)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
