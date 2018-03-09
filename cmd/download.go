package cmd

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"path"
	"path/filepath"

	"github.com/cheggaaa/pb"
	"github.com/jarsen/pb/db"
	. "github.com/spf13/cobra"
)

var downloadCmd = &Command{
	Use:   "download TARGET_DIR",
	Short: "downloads all the image in the library to disk",
	Args:  ExactArgs(1),
	Run: func(cmd *Command, args []string) {
		targetDir := args[0]

		index, err := db.Init()
		if err != nil {
			log.Fatal(err)
		}

		results, err := db.AllImages(index)
		if err != nil {
			log.Fatal(err)
		}

		errc := make(chan error)
		for i, hit := range results.Hits {
			url := hit.Fields["URL"].(string)
			filePath := filePathForURL(targetDir, url, i)
			go func() {
				_, err := downloadUrlToFile(url, filePath)
				if err != nil {
					errc <- fmt.Errorf("Error downloading %s: %s", url, err)
				} else {
					errc <- nil
				}
			}()
		}

		total := int(results.Total)
		bar := pb.New(total)
		bar.SetMaxWidth(80)
		bar.Start()
		for i := 0; i < total; i++ {
			err = <-errc
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
			bar.Increment()
		}
		bar.Finish()
	},
}

func filePathForURL(dir string, urlStr string, unique int) string {
	url, _ := url.Parse(urlStr)
	urlFileName := path.Base(url.EscapedPath())
	fileName := fmt.Sprintf("%d%s", unique, urlFileName)
	return filepath.Join(dir, fileName)
}

func init() {
	rootCmd.AddCommand(downloadCmd)
}
