package cmd

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
	"quotes/quotes"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(loadCmd)
}

var loadCmd = &cobra.Command{
	Use:   "load",
	Short: "Loads a kobo annotations directory",
	Run: func(cmd *cobra.Command, args []string) {
		quotes, err := quotes.LoadQuotes(annotationsDir)
		if err != nil {
			log.Fatal(err)
		}

		file, err := os.Create(loadedFile)
		if os.IsNotExist(err) {
			err = os.Mkdir(quotesAppDir, 0755)
			if err != nil {
				log.Fatal(err)
			}
			file, err = os.Create(loadedFile)
		}
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		w := bufio.NewWriter(file)

		for _, quote := range quotes {
			b, err := json.Marshal(quote)
			if err != nil {
				log.Fatal(err)
			}
			_, err = w.WriteString(string(b) + "\n")
			if err != nil {
				log.Fatal(err)
			}
		}

		err = w.Flush()
		if err != nil {
			log.Fatal(err)
		}
	},
}
