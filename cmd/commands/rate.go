package commands

import (
	"fmt"

	"github.com/currency-cli/internal/api/rate"
	"github.com/currency-cli/internal/logger"
	"github.com/currency-cli/internal/pkg"
	"github.com/spf13/cobra"
)

const (
	ColorReset = "\033[0m"
	ColorBlue  = "\033[36m"
	ColorRed   = "\033[31m"
)

var currencyCode string

var currencyCmd = &cobra.Command{
	Use:   "rate",
	Short: "Rate allows you to get a list of exchange rates.",
	Run: func(cmd *cobra.Command, args []string) {
		if currencyCode == "" {
			fmt.Println("["+ColorBlue+"INFO"+ColorReset+"] "+"Specify the currency ID with the --currency or -c flag.")
			return
		}

		logger.Info("Trying to get currency rate with currency_id: %s", currencyCode)

		data, err := rate.GetCurrencyRate(currencyCode)
		if err != nil {
			logger.Error("API error: %v", err)
			fmt.Printf("["+ColorRed+"ERROR"+ColorReset+"] "+"Error getting currency rate.\n")
			return
		}

		fmt.Printf(
			"Base code: %s\nConversionRates:\n",
			data.BaseCode,
		)

		pkg.PrintGrid(data.ConversionRates)
	},
}
