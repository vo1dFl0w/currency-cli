package main

import (
	"os"
	
	"github.com/currency-cli/cmd"
	"github.com/currency-cli/cmd/commands"
	"github.com/currency-cli/internal/logger"
)

var Log logger.Logger

func main() {
	rootCmd := cmd.RootCmd()

	_ = rootCmd.ParseFlags(os.Args[1:])
	Log = logger.InitLogger(cmd.VerboseFlag)
	defer Log.Close()

	commands.Init(rootCmd, Log)
	cmd.Execute(rootCmd)
}