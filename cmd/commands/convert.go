package commands

import (
	"fmt"
	"strconv"

	"github.com/currency-cli/internal/api/exchange"
	"github.com/spf13/cobra"
)

var (
	fromCurrency, toCurrency, value string
)

var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Convert allows you to get the exchange rate of a certain currency in relation to another or get the result of currency conversion.",
	Run: func(cmd *cobra.Command, args []string) {
		if fromCurrency == "" || toCurrency == "" {
			fmt.Println("Specify the base and target codes using the --from -to flags.")
			return
		}

		switch {
		case value != "":
			Log.Info("Trying to convert %s %s to %s", value, fromCurrency, toCurrency)
		default:
			Log.Info("Trying to convert %s to %s", fromCurrency, toCurrency)
		}

		data, err := exchange.Convert(fromCurrency, toCurrency)
		if err != nil {
			Log.Error("API error: %v", err)
			return
		}

		switch {
		case value != "":
			v, err := strconv.ParseFloat(value, 64)
			if err != nil {
				Log.Error("Cannot convert value '%s' to float64: %s", value, err)
				return
			}
			fmt.Printf("Base code: %s\nTarget code: %s\nConversion rate:%.4f\n", data.BaseCode, data.TargetCode, data.ConversionRate*v)
		default:
			fmt.Printf("Base code: %s\nTarget code: %s\nConversion rate:%.4f\n", data.BaseCode, data.TargetCode, data.ConversionRate)
		}
	},
}
