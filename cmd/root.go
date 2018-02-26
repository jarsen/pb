package cmd

import (
	"fmt"
	"os"

	. "github.com/spf13/cobra"
)

var rootCmd = &Command{
	Use:   "pb [SEARCH_TERMS...]",
	Short: "Copies the highest matching result to your system clipboard",
	Long:  `pb - your own personal gif/meme databas.\nBecome the envy of your friends and colleagues as you organize and search your favorite memes and animated gifs from your terminal.`,
	Args:  MinimumNArgs(1),
	Run: func(cmd *Command, args []string) {
		copyFirstSearchResult(args)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
