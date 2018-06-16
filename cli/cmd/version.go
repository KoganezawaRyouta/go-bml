package cmd

import (
	"log"

	"github.com/KoganezawaRyouta/go-bml/config"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Long:  `Print the version number`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Printf("v%s-%s\n", config.Version, config.GoVersion)
		log.Printf("BuildDhash%s\n", config.BuildDhash)
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
