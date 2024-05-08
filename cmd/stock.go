package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// stockCmd represents the stock command
var stockCmd = &cobra.Command{
	Use:   "stock",
	Short: "Retrieve information about the stock.",
	Long:  `View the stock's current price, historical prices, profile, & statistics`,
	Example: `  equity-pulse stock --name AAPL
  equity-pulse stock --ticker aapl
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("stock called %s\n", args)
	},
}

func init() {
	rootCmd.AddCommand(stockCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// stockCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// stockCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
