package main

import (
	"log"

	"github.com/currency-cli/cmd"
	"github.com/currency-cli/cmd/commands"
)

func main() {
	rootCmd := cmd.RootCmd()
	commands.Init(rootCmd)

	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("initial settup failed: %v\n", err)
	}
}