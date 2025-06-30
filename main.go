package main

import (
	"github.com/currency-cli/cmd"
	"github.com/currency-cli/cmd/commands"
	"github.com/currency-cli/internal/logger"
)

func main() {
	logger.InitLogger()

	rootCmd := cmd.RootCmd()
	commands.Init(rootCmd)
	cmd.Execute(rootCmd)
}