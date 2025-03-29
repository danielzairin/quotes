package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"quotes/quotes"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(randomCmd)
}

var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "Get a random quote",
	Run: func(cmd *cobra.Command, args []string) {
		var q []quotes.Quote
		file, err := os.Open(loadedFile)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			var quote quotes.Quote
			err := json.Unmarshal(scanner.Bytes(), &quote)
			if err != nil {
				log.Fatal(err)
			}
			q = append(q, quote)
		}

		randomIndex := rand.Intn(len(q))
		randomQuote := q[randomIndex]

		fmt.Println(randomQuote)
	},
}
