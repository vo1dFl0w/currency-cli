package commands

import (
	"github.com/currency-cli/internal/logger"
	"github.com/spf13/cobra"
)

const (
	ColorReset  = "\033[0m"
	ColorBlue   = "\033[36m"
	ColorRed    = "\033[31m"
	ColorYellow = "\033[33m"
)

var Log logger.Logger

func Init(rootCmd *cobra.Command, l logger.Logger) {
	Log = l
	currencyCmd.Flags().StringVarP(&currencyCode, "currency", "c", "", "Currency ID")

	convertCmd.Flags().StringVarP(&fromCurrency, "from", "f", "", "Initial currency")
	convertCmd.Flags().StringVarP(&toCurrency, "to", "t", "", "Final currency")
	convertCmd.Flags().StringVar(&value, "value", "", "Value to convert")

	convertCmd.MarkFlagRequired(fromCurrency)
	convertCmd.MarkFlagRequired(toCurrency)

	rootCmd.AddCommand(currencyCmd, convertCmd)
}