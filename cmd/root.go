package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	version     = "v2.0.0"
	showVersion bool
)

func RootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "currency",
		Short: "Currency CLI shows the current exchange rate of the currency.",
		Long: `Currency CLI is a simple tool to get currency rate or convert currency using currency identifier.
		Example:
			currency rate --currency
			currency convert --from=EUR --to=USD
		`,
		Run: func(cmd *cobra.Command, args []string) {
			if showVersion {
				fmt.Printf("Currency CPI %s\n", version)
			}
		},
	}

	rootCmd.PersistentFlags().BoolVarP(&showVersion, "version", "v", false, "Shows")

	return rootCmd
}

func Execute(rootCmd *cobra.Command) {
	cobra.CheckErr(rootCmd.Execute())
}
