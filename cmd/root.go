package cmd

import (
	"fmt"
	"os"
	"path"

	"github.com/spf13/cobra"
)

func getHomeDirOrPanic() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	return homeDir
}

var quotesAppDir = path.Join(getHomeDirOrPanic(), ".quotes")
var loadedFile = path.Join(quotesAppDir, "loaded.ndjson")

var rootCmd = &cobra.Command{
	Use: "quotes",
	CompletionOptions: cobra.CompletionOptions{
		HiddenDefaultCmd: true,
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
