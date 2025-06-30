package commands

import "github.com/spf13/cobra"

func Init(rootCmd *cobra.Command) {
	currencyCmd.Flags().StringVarP(&currencyCode, "currency", "c", "", "Currency ID")

	convertCmd.Flags().StringVarP(&fromCurrency, "from", "f", "", "Initial currency")
	convertCmd.Flags().StringVarP(&toCurrency, "to", "t", "", "Final currency")
	convertCmd.Flags().StringVar(&value, "value", "", "Value to convert")

	convertCmd.MarkFlagRequired(fromCurrency)
	convertCmd.MarkFlagRequired(toCurrency)

	rootCmd.AddCommand(currencyCmd, convertCmd)
}