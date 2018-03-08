package cmd

import (
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"syscall"

	. "github.com/spf13/cobra"
)

var showCmd = &Command{
	Use:   "show",
	Short: "shows the image for the top result and copies url to clipboard",
	Long:  `Downloads the image to a tmp directory, then shows it`,
	Run: func(cmd *Command, args []string) {
		binary, err := exec.LookPath("imgcat")
		if err != nil {
			log.Fatal("Must have imgcat installed")
		}
		url := copyFirstSearchResult(args)
		tempFile := downloadUrlToTempFile(url)
		err = syscall.Exec(binary, []string{"imgcat", tempFile}, os.Environ())
		if err != nil {
			log.Fatalf("Error showing image: %s", err)
		}
	},
}

func downloadUrlToTempFile(url string) string {
	tmpPath := filepath.Join(os.TempDir(), path.Base(url))
	out, err := os.Create(tmpPath)
	if err != nil {
		log.Fatal("error creating temp file")
	}
	defer out.Close()

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("error downloading file")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("bad status :%s", resp.Status)
	}

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		log.Fatalf("error copying file")
	}

	return tmpPath
}

func init() {
	rootCmd.AddCommand(showCmd)
}
