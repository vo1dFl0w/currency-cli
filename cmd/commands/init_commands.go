package commands

import "github.com/spf13/cobra"

func Init(rootCmd *cobra.Command) {
	currencyCmd.Flags().StringVarP(&currencyID, "currency", "c", "", "Currency ID")

	rootCmd.AddCommand(currencyCmd)

}