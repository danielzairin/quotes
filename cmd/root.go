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
var annotationsDir = path.Join(getHomeDirOrPanic(), "annotations")
var loadedFile = path.Join(quotesAppDir, "loaded.ndjson")

var rootCmd = &cobra.Command{
	Use:   "quotes",
	Short: "Quotes from kobo",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		fmt.Println("i am root")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
