package cmd

import (
	"log"
	"net/url"
	"strings"
	"time"

	"github.com/blevesearch/bleve"
	"github.com/google/uuid"
	"github.com/jarsen/pb/db"
	. "github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &Command{
	Use:   "add URL [DESCRIPTION...]",
	Short: "add a link and some description text",
	Args:  MinimumNArgs(2),
	Run: func(cmd *Command, args []string) {
		addImage(args[0], args[1:])
	},
}

func addImage(urlString string, descriptionTokens []string) {
	description := strings.Join(descriptionTokens, " ")

	var err error
	if _, err = url.Parse(urlString); err != nil {
		log.Fatal("First argument must be a valid URL")
	}

	image := db.Image{
		ID:          uuid.New().String(),
		URL:         urlString,
		Description: description,
		Date:        time.Now(),
	}

	var index bleve.Index
	index, err = db.Init()
	if err != nil {
		log.Fatal(err)
	}
	image.AddTo(index)
}

func init() {
	rootCmd.AddCommand(addCmd)
}
