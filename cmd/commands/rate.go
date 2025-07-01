package commands

import (
	"fmt"

	"github.com/currency-cli/internal/api/exchange"
	"github.com/currency-cli/internal/pkg"
	"github.com/spf13/cobra"
)

var currencyCode string

var currencyCmd = &cobra.Command{
	Use:   "rate",
	Short: "Rate allows you to get a list of exchange rates.",
	Run: func(cmd *cobra.Command, args []string) {
		if currencyCode == "" {
			fmt.Println("Specify the currency code with the --currency or -c flag.")
			return
		}

		Log.Info("Trying to get currency rate with currency code: %s\n", currencyCode)

		data, err := exchange.GetCurrencyRate(currencyCode)
		if err != nil {
			Log.Error("API error: %v", err)
			return
		}

		fmt.Printf(
			"Base code: %s\nConversionRates:\n",
			data.BaseCode,
		)

		pkg.PrintGrid(data.ConversionRates)
	},
}
