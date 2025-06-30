package commands

import (
	"fmt"
	"strconv"

	"github.com/currency-cli/internal/api/convert"
	"github.com/currency-cli/internal/logger"
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
			fmt.Println("[" + ColorBlue + "INFO" + ColorReset + "] " + "Specify the base and target codes using the --from -to flags.")
			return
		}

		switch {
		case value != "":
			logger.Info("Trying to convert %f %s to %s", fromCurrency, toCurrency, value)
		default:
			logger.Info("Trying to convert %s to %s", fromCurrency, toCurrency)
		}

		data, err := convert.Convert(fromCurrency, toCurrency)
		if err != nil {
			logger.Error("API error: %v", err)
			fmt.Printf("[" + ColorRed + "ERROR" + ColorReset + "] " + "Error getting convertion result.\n")
			return
		}

		switch {
		case value != "":
			v, err := strconv.ParseFloat(value, 64)
			if err != nil {
				logger.Error("Cannot convert value '%s' to float64: %s", value, err)
				fmt.Printf("Wrong value expression: %s\n", value)
				return
			}
			fmt.Printf("Base code: %s\nTarget code: %s\nConversion rate:%.4f\n", data.BaseCode, data.TargetCode, data.ConversionRate*v)
		default:
			fmt.Printf("Base code: %s\nTarget code: %s\nConversion rate:%.4f\n", data.BaseCode, data.TargetCode, data.ConversionRate)
		}
	},
}
