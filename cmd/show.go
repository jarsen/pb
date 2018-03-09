package cmd

import (
	"fmt"
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
		filePath := tempFilePath(url)
		_, err = downloadUrlToFile(url, filePath)
		if err != nil {
			log.Fatalf("Error downloading file: %s", err)
		}
		err = syscall.Exec(binary, []string{"imgcat", filePath}, os.Environ())
		if err != nil {
			log.Fatalf("Error showing image: %s", err)
		}
	},
}

func tempFilePath(url string) string {
	return filepath.Join(os.TempDir(), path.Base(url))
}

func downloadUrlToFile(url string, filePath string) (int64, error) {
	out, err := os.Create(filePath)
	if err != nil {
		return 0, err
	}
	defer out.Close()

	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("bad status :%s", resp.Status)
	}

	return io.Copy(out, resp.Body)
}

func init() {
	rootCmd.AddCommand(showCmd)
}
