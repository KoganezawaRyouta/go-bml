package main

import (
	"fmt"
	"os"

	"github.com/KoganezawaRyouta/go-bml/cli/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
