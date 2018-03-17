package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/blevesearch/bleve"
	"github.com/jarsen/pb/db"
	. "github.com/spf13/cobra"
)

var dumpCmd = &Command{
	Use:   "dump",
	Short: "dump the database into the specified txt file",
	Args:  ExactArgs(1),
	Run: func(cmd *Command, args []string) {
		path := args[0]
		file, err := os.Create(path)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		w := bufio.NewWriter(file)

		var index bleve.Index
		index, err = db.Init()
		results, err := db.AllImages(index)
		if err != nil {
			log.Fatalf("Error getting images: %s\n", err)
		}
		for _, hit := range results.Hits {
			line := fmt.Sprintf("%s %s\n", hit.Fields["URL"], hit.Fields["Description"])
			_, err = w.WriteString(line)
			if err != nil {
				log.Fatal("Error writing to file")
			}
		}

		w.Flush()
	},
}

func init() {
	rootCmd.AddCommand(dumpCmd)
}
