package cmd

import (
	"fmt"
	"log"
	"quotes/quotes"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(infoCmd)
}

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Gets info of loaded quotes",
	Run: func(cmd *cobra.Command, args []string) {
		q, err := quotes.GetLoadedQuotesFromFile(loadedFile)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("There are %d quotes\n", len(q))
	},
}
