package cmd

import (
	"log"

	"github.com/KoganezawaRyouta/go-bml/config"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"

	"os"
)

var conf config.Config

var RootCmd = &cobra.Command{
	Use:   "root",
	Short: "Golangを思い出す",
	Long:  "Golangを思い出す",
	Run: func(cmd *cobra.Command, args []string) {
		// ...
	},
}

func init() {
	cobra.OnInitialize(func() {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
		conf.LogFile = os.Getenv("LOG_FILE")
	})
}
