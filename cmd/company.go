/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// companyCmd represents the company command
var companyCmd = &cobra.Command{
	Use:   "company",
	Short: "Retrieve information about the company.",
	Long:  `View the company's current stock price, historical stock prices, profile, & statistics`,
	Example: `  equity-pulse company --name AAPL
  equity-pulse company --ticker aapl
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("company called %s\n", args)
	},
}

func init() {
	rootCmd.AddCommand(companyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// companyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// companyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
