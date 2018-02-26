package cmd

import (
	"bufio"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/blevesearch/bleve"
	"github.com/google/uuid"
	"github.com/jarsen/pb/db"
	. "github.com/spf13/cobra"
)

var loadCmd = &Command{
	Use:   "load [PATH]",
	Short: "adds all the images from a given file to the database",
	Args:  ExactArgs(1),
	Run: func(cmd *Command, args []string) {
		path := args[0]
		file, err := os.Open(path)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		var index bleve.Index
		index, err = db.Init()
		batch := index.NewBatch()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			text := scanner.Text()
			tokens := strings.Split(text, " ")
			description := strings.Join(tokens[1:], " ")
			urlString := tokens[0]
			if _, err = url.Parse(urlString); err != nil {
				fmt.Println(urlString, "is not a valid URL")
				continue
			}

			uuid := uuid.New().String()
			image := db.Image{
				ID:          uuid,
				URL:         urlString,
				Description: description,
				Date:        time.Now(),
			}
			batch.Index(uuid, image)
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
		index.Batch(batch)
	},
}

func init() {
	rootCmd.AddCommand(loadCmd)
}
