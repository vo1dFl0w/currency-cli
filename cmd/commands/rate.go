package commands

import (
	"fmt"

	"github.com/currency-cli/internal/api/rate"
	"github.com/currency-cli/internal/pkg"
	"github.com/spf13/cobra"
)

var currencyID string

var currencyCmd = &cobra.Command{
	Use:   "rate",
	Short: "",
	Run: func(cmd *cobra.Command, args []string) {
		if currencyID == "" {
			fmt.Println("Specify ...")
		}

		fmt.Println("Trying to get currency rate with currency_id: ", currencyID)

		data, err := rate.GetCurrencyRate(currencyID)
		if err != nil {
			fmt.Println(fmt.Errorf("Error getting currency rate: %v", err))
		}

		fmt.Printf(
			"Base code: %s\nConversionRates:\n",
			data.BaseCode,
		)

		pkg.PrintGrid(data.ConversionRates)
	},
}
